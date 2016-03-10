package commands_test

import (
	"errors"

	"github.com/cloudfoundry-incubator/ducati-daemon/executor/commands"
	"github.com/cloudfoundry-incubator/ducati-daemon/executor/commands/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetLinkMaster", func() {
	var (
		context       *fakes.Context
		masterSetter  *fakes.MasterSetter
		setLinkMaster commands.SetLinkMaster
	)

	BeforeEach(func() {
		context = &fakes.Context{}
		masterSetter = &fakes.MasterSetter{}
		context.MasterSetterReturns(masterSetter)

		setLinkMaster = commands.SetLinkMaster{
			Master: "master-dev",
			Slave:  "slave-dev",
		}
	})

	It("assigns a master to the slave", func() {
		err := setLinkMaster.Execute(context)
		Expect(err).NotTo(HaveOccurred())

		Expect(masterSetter.SetMasterCallCount()).To(Equal(1))
		slave, master := masterSetter.SetMasterArgsForCall(0)
		Expect(slave).To(Equal("slave-dev"))
		Expect(master).To(Equal("master-dev"))
	})

	Context("when the master setter fails", func() {
		BeforeEach(func() {
			masterSetter.SetMasterReturns(errors.New("you're not a slave"))
		})

		It("wraps and propogates the error", func() {
			err := setLinkMaster.Execute(context)
			Expect(err).To(MatchError("set link master: you're not a slave"))
		})
	})

	Describe("String", func() {
		It("describes itself", func() {
			Expect(setLinkMaster.String()).To(Equal("ip link set slave-dev master master-dev"))
		})
	})
})