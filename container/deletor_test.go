package container_test

import (
	"errors"

	"github.com/cloudfoundry-incubator/ducati-daemon/container"
	"github.com/cloudfoundry-incubator/ducati-daemon/executor/commands"
	cmd_fakes "github.com/cloudfoundry-incubator/ducati-daemon/executor/commands/fakes"
	exec_fakes "github.com/cloudfoundry-incubator/ducati-daemon/executor/fakes"
	"github.com/cloudfoundry-incubator/ducati-daemon/fakes"
	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	var (
		deletor           container.Deletor
		executor          *exec_fakes.Executor
		sandboxRepoLocker *cmd_fakes.Locker
		watcher           *fakes.MissWatcher
	)

	BeforeEach(func() {
		executor = &exec_fakes.Executor{}
		sandboxRepoLocker = &cmd_fakes.Locker{}
		watcher = &fakes.MissWatcher{}
		deletor = container.Deletor{
			Executor: executor,
			Locker:   sandboxRepoLocker,
			Watcher:  watcher,
		}
	})

	It("should construct the correct command sequence", func() {
		deletorConfig := container.DeletorConfig{
			InterfaceName:   "some-interface-name",
			ContainerNSPath: "/some/container/namespace/path",
			SandboxNSPath:   "/some/sandbox/namespace/path",
			VxlanDeviceName: "some-vxlan",
		}

		err := deletor.Delete(deletorConfig)
		Expect(err).NotTo(HaveOccurred())

		Expect(executor.ExecuteCallCount()).To(Equal(1))

		Expect(executor.ExecuteArgsForCall(0)).To(Equal(
			commands.All(
				commands.InNamespace{
					Namespace: namespace.NewNamespace("/some/container/namespace/path"),
					Command: commands.DeleteLink{
						LinkName: "some-interface-name",
					},
				},

				commands.CleanupSandbox{
					Namespace:       namespace.NewNamespace("/some/sandbox/namespace/path"),
					Locker:          sandboxRepoLocker,
					Watcher:         watcher,
					VxlanDeviceName: "some-vxlan",
				},
			),
		))
	})

	Context("when executing fails", func() {
		It("should return the error", func() {
			executor.ExecuteReturns(errors.New("boom"))

			deletorConfig := container.DeletorConfig{
				InterfaceName:   "some-interface-name",
				ContainerNSPath: "/some/container/namespace/path",
				SandboxNSPath:   "/some/sandbox/namespace/path",
			}

			err := deletor.Delete(deletorConfig)
			Expect(err).To(MatchError("boom"))
		})
	})

})
