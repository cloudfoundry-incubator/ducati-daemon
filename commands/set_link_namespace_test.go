package commands_test

import (
	"errors"

	"github.com/cloudfoundry-incubator/ducati-daemon/commands"
	"github.com/cloudfoundry-incubator/ducati-daemon/commands/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetLinkNamespace", func() {
	var (
		context          *fakes.Context
		setNamespacer    *fakes.SetNamespacer
		setLinkNamespace commands.SetLinkNamespace
	)

	BeforeEach(func() {
		context = &fakes.Context{}
		setNamespacer = &fakes.SetNamespacer{}
		context.SetNamespacerReturns(setNamespacer)

		setLinkNamespace = commands.SetLinkNamespace{
			Name:      "link-name",
			Namespace: "some-namespace-path",
		}
	})

	It("moves the link to the target namespace", func() {
		err := setLinkNamespace.Execute(context)
		Expect(err).NotTo(HaveOccurred())

		Expect(setNamespacer.SetNamespaceCallCount()).To(Equal(1))
		name, ns := setNamespacer.SetNamespaceArgsForCall(0)
		Expect(name).To(Equal("link-name"))
		Expect(ns).To(Equal("some-namespace-path"))
	})

	Context("when moving the link fails", func() {
		It("propagates the error", func() {
			setNamespacer.SetNamespaceReturns(errors.New("welp"))

			err := setLinkNamespace.Execute(context)
			Expect(err).To(MatchError("welp"))
		})
	})
})
