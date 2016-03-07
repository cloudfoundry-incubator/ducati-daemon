package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"

	"github.com/cloudfoundry-incubator/ducati-daemon/container"
	exec_fakes "github.com/cloudfoundry-incubator/ducati-daemon/executor/fakes"
	"github.com/cloudfoundry-incubator/ducati-daemon/fakes"
	"github.com/cloudfoundry-incubator/ducati-daemon/handlers"
	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"
	"github.com/cloudfoundry-incubator/ducati-daemon/models"
	"github.com/pivotal-golang/lager/lagertest"
	"github.com/tedsuo/rata"
)

var _ = Describe("NetworksDeleteContainer", func() {
	var (
		logger      *lagertest.TestLogger
		datastore   *fakes.Store
		executor    *exec_fakes.Executor
		deletor     *fakes.Deletor
		handler     http.Handler
		request     *http.Request
		osLocker    *fakes.OSThreadLocker
		unmarshaler *fakes.Unmarshaler

		sandboxRepo *fakes.Repository

		payload models.NetworksDeleteContainerPayload
	)

	var setPayload = func() {
		payloadBytes, err := json.Marshal(payload)
		Expect(err).NotTo(HaveOccurred())
		request.Body = ioutil.NopCloser(bytes.NewBuffer(payloadBytes))
	}

	BeforeEach(func() {
		osLocker = &fakes.OSThreadLocker{}

		unmarshaler = &fakes.Unmarshaler{}
		unmarshaler.UnmarshalStub = json.Unmarshal

		logger = lagertest.NewTestLogger("test")
		datastore = &fakes.Store{}
		executor = &exec_fakes.Executor{}
		deletor = &fakes.Deletor{}

		sandboxRepo = &fakes.Repository{}

		deleteHandler := &handlers.NetworksDeleteContainer{
			Unmarshaler:    unmarshaler,
			Logger:         logger,
			Datastore:      datastore,
			Deletor:        deletor,
			OSThreadLocker: osLocker,
			SandboxRepo:    sandboxRepo,
		}

		sandboxRepo.GetReturns(namespace.NewNamespace("/some/sandbox/repo/path"), nil)

		handler, request = rataWrap(deleteHandler, "DELETE", "/networks/:network_id/:container_id", rata.Params{
			"network_id":   "some-network-id",
			"container_id": "some-container-id",
		})
		payload = models.NetworksDeleteContainerPayload{
			InterfaceName:      "some-interface-name",
			ContainerNamespace: "/some/container/namespace/path",
			VNI:                42,
		}
		setPayload()
	})

	It("computes the sandbox name from the VNI", func() {
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, request)

		Expect(sandboxRepo.GetCallCount()).To(Equal(1))
		Expect(sandboxRepo.GetArgsForCall(0)).To(Equal("vni-42"))
	})

	It("deletes the container from the network", func() {
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, request)

		Expect(deletor.DeleteCallCount()).To(Equal(1))
		Expect(deletor.DeleteArgsForCall(0)).To(Equal(container.DeletorConfig{
			InterfaceName:   "some-interface-name",
			ContainerNSPath: "/some/container/namespace/path",
			SandboxNSPath:   "/some/sandbox/repo/path",
			VxlanDeviceName: "vxlan42",
		}))
	})

	It("deletes the container from the datastore", func() {
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, request)

		Expect(datastore.DeleteCallCount()).To(Equal(1))
		containerID := datastore.DeleteArgsForCall(0)
		Expect(containerID).To(Equal("some-container-id"))
	})

	It("responds with status no content", func() {
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, request)

		Expect(resp.Code).To(Equal(http.StatusNoContent))
	})

	It("locks and unlocks the os thread", func() {
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, request)

		Expect(osLocker.LockOSThreadCallCount()).To(Equal(1))
		Expect(osLocker.UnlockOSThreadCallCount()).To(Equal(1))
	})

	Context("when the request body cannot be read", func() {
		BeforeEach(func() {
			request.Body = ioutil.NopCloser(&badReader{})
		})

		It("should log and respond with status 400", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(resp.Code).To(Equal(http.StatusBadRequest))
			Expect(logger).To(gbytes.Say("networks-delete-containers.*body-read-failed"))
		})
	})

	Context("when the request body is not valid JSON", func() {
		BeforeEach(func() {
			request.Body = ioutil.NopCloser(strings.NewReader(`{{{`))
		})

		It("should log and respond with status 400", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(resp.Code).To(Equal(http.StatusBadRequest))
			Expect(logger).To(gbytes.Say("networks-delete-containers.*unmarshal-failed"))
		})
	})

	DescribeTable("missing payload fields",
		func(paramToRemove, jsonName string) {
			field := reflect.ValueOf(&payload).Elem().FieldByName(paramToRemove)
			if !field.IsValid() {
				Fail("invalid test: payload does not have a field named " + paramToRemove)
			}
			field.Set(reflect.Zero(field.Type()))
			setPayload()

			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(resp.Code).To(Equal(http.StatusBadRequest))
			Expect(logger).To(gbytes.Say(fmt.Sprintf(
				"networks-delete-containers.bad-request.*missing-%s", jsonName)))
		},
		Entry("interface", "InterfaceName", "interface_name"),
		Entry("container_namespace_path", "ContainerNamespace", "container_namespace"),
		Entry("vni", "VNI", "vni"),
	)

	Context("when the sandbox repo fails", func() {
		BeforeEach(func() {
			sandboxRepo.GetReturns(nil, errors.New("some-repo-error"))
		})

		It("should log and respond with status 500", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			Expect(logger).To(gbytes.Say("networks-delete-containers.sandbox-repo.*some-repo-error"))
		})
	})

	Context("when deleting the container from the network fails", func() {
		BeforeEach(func() {
			deletor.DeleteReturns(errors.New("some-deletor-error"))
		})

		It("should log and respond with status 500", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			Expect(logger).To(gbytes.Say("networks-delete-containers.deletor.delete-failed.*some-deletor-error"))
		})

		It("should not remove the container from the datastore", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(datastore.DeleteCallCount()).To(Equal(0))
		})

		It("locks and unlocks the os thread", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(osLocker.LockOSThreadCallCount()).To(Equal(1))
			Expect(osLocker.UnlockOSThreadCallCount()).To(Equal(1))
		})
	})

	Context("when deleting from the datastore fails", func() {
		BeforeEach(func() {
			datastore.DeleteReturns(errors.New("some-datastore-error"))
		})

		It("should log and respond with status 500", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			Expect(logger).To(gbytes.Say("networks-delete-containers.datastore.delete-failed.*some-datastore-error"))
		})

		It("locks and unlocks the os thread", func() {
			resp := httptest.NewRecorder()
			handler.ServeHTTP(resp, request)

			Expect(osLocker.LockOSThreadCallCount()).To(Equal(1))
			Expect(osLocker.UnlockOSThreadCallCount()).To(Equal(1))
		})
	})
})
