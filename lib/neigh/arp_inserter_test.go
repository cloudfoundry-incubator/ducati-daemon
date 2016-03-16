package neigh_test

import (
	"errors"
	"net"
	"os"

	"github.com/cloudfoundry-incubator/ducati-daemon/fakes"
	"github.com/cloudfoundry-incubator/ducati-daemon/lib/neigh"
	nl_fakes "github.com/cloudfoundry-incubator/ducati-daemon/lib/nl/fakes"
	"github.com/cloudfoundry-incubator/ducati-daemon/watcher"
	"github.com/pivotal-golang/lager/lagertest"
	"github.com/vishvananda/netlink"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("ARPInserter", func() {
	var (
		inserter     *neigh.ARPInserter
		ns           *fakes.Namespace
		netlinker    *nl_fakes.Netlinker
		resolved     chan watcher.Neighbor
		logger       *lagertest.TestLogger
		threadLocker *fakes.OSThreadLocker
	)

	BeforeEach(func() {
		ns = &fakes.Namespace{}
		netlinker = &nl_fakes.Netlinker{}
		resolved = make(chan watcher.Neighbor, 3)
		logger = lagertest.NewTestLogger("test")
		threadLocker = &fakes.OSThreadLocker{}

		inserter = &neigh.ARPInserter{
			Logger:         logger,
			Netlinker:      netlinker,
			OSThreadLocker: threadLocker,
		}
		ns.ExecuteStub = func(callback func(ns *os.File) error) error {
			callback(nil)
			return nil
		}
	})

	Describe("HandleResolvedNeighbors", func() {
		var neighbor watcher.Neigh

		BeforeEach(func() {
			mac, err := net.ParseMAC("01:02:03:04:05:06")
			Expect(err).NotTo(HaveOccurred())

			neighbor = watcher.Neigh{
				LinkIndex:    99,
				Family:       22,
				State:        11,
				Type:         31,
				Flags:        17,
				IP:           net.ParseIP("1.2.3.4"),
				HardwareAddr: mac,
			}

			resolved <- watcher.Neighbor{
				SandboxName: "some-sandbox",
				Neigh:       neighbor,
			}

			close(resolved)
		})

		It("adds a neighbor entry via the netlinker from within the namespace", func() {
			ns.ExecuteStub = func(callback func(ns *os.File) error) error {
				Expect(netlinker.AddNeighCallCount()).To(Equal(0))
				callback(nil)
				Expect(netlinker.AddNeighCallCount()).To(Equal(1))

				Expect(netlinker.AddNeighArgsForCall(0)).To(Equal(&netlink.Neigh{
					LinkIndex:    neighbor.LinkIndex,
					Family:       neighbor.Family,
					State:        neighbor.State,
					Type:         neighbor.Type,
					Flags:        neighbor.Flags,
					IP:           neighbor.IP,
					HardwareAddr: neighbor.HardwareAddr,
				}))
				return nil
			}

			err := inserter.HandleResolvedNeighbors(ns, resolved)
			Expect(err).NotTo(HaveOccurred())

			Eventually(ns.ExecuteCallCount()).Should(Equal(1))
		})

		It("runs everything in a goroutine and does not block", func() {
			quitChan := make(chan struct{})
			ns.ExecuteStub = func(callback func(_ *os.File) error) error {
				<-quitChan
				return callback(nil)
			}

			exited := make(chan struct{})
			go func() {
				inserter.HandleResolvedNeighbors(ns, resolved)
				close(exited)
			}()

			close(quitChan)
			Eventually(exited).Should(BeClosed())
		})

		It("locks and unlocks the OS thread", func() {
			netlinker.AddNeighStub = func(_ *netlink.Neigh) error {
				Expect(threadLocker.UnlockOSThreadCallCount()).To(Equal(0))
				return nil
			}

			err := inserter.HandleResolvedNeighbors(ns, resolved)
			Expect(err).NotTo(HaveOccurred())

			Expect(threadLocker.LockOSThreadCallCount()).To(Equal(1))
			Expect(threadLocker.UnlockOSThreadCallCount()).To(Equal(1))
		})

		Context("when executing in namespace fails", func() {
			It("returns the error", func() {
				ns.ExecuteReturns(errors.New("peppers"))

				err := inserter.HandleResolvedNeighbors(ns, resolved)
				Expect(err).To(MatchError("namespace execute failed: peppers"))
			})
		})

		Context("when adding a neighbor entry fails", func() {
			It("logs the error", func() {
				netlinker.AddNeighReturns(errors.New("durian"))

				err := inserter.HandleResolvedNeighbors(ns, resolved)
				Expect(err).NotTo(HaveOccurred())

				Eventually(logger).Should(gbytes.Say("add-neighbor-failed.*durian"))
			})
		})
	})
})
