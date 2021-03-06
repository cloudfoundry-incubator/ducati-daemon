package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/appc/cni/pkg/types"
	"github.com/cloudfoundry-incubator/ducati-daemon/client"
	"github.com/cloudfoundry-incubator/ducati-daemon/config"
	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"
	"github.com/cloudfoundry-incubator/ducati-daemon/models"
	"github.com/cloudfoundry-incubator/ducati-daemon/network"
	"github.com/cloudfoundry-incubator/ducati-daemon/ossupport"
	"github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-golang/lager/lagertest"
)

var _ = Describe("Sandbox Consistency", func() {
	var (
		session        *gexec.Session
		ducatiCmd      *exec.Cmd
		address        string
		containerID    string
		vni            int
		sandboxName    string
		sandboxRepoDir string
		spaceID        string
		appID          string
		networkID      string

		logger *lagertest.TestLogger
		// sandboxRepo        namespace.Repository
		containerRepo      namespace.Repository
		containerNamespace namespace.Namespace

		configFilePath string
	)

	BeforeEach(func() {
		var err error
		address = fmt.Sprintf("127.0.0.1:%d", 4001+GinkgoParallelNode())
		sandboxRepoDir, err = ioutil.TempDir("", "sandbox")
		Expect(err).NotTo(HaveOccurred())

		logger = lagertest.NewTestLogger("test")
		threadLocker := &ossupport.OSLocker{}

		// sandboxRepo, err = namespace.NewRepository(logger, sandboxRepoDir, threadLocker)
		// Expect(err).NotTo(HaveOccurred())

		containerRepoDir, err := ioutil.TempDir("", "containers")
		Expect(err).NotTo(HaveOccurred())

		containerRepo, err = namespace.NewRepository(logger, containerRepoDir, threadLocker)
		Expect(err).NotTo(HaveOccurred())

		guid, err := uuid.NewV4()
		Expect(err).NotTo(HaveOccurred())

		containerNamespace, err = containerRepo.Create(guid.String())
		Expect(err).NotTo(HaveOccurred())

		configFilePath = writeConfigFile(config.Daemon{
			ListenHost:        "127.0.0.1",
			ListenPort:        4001 + GinkgoParallelNode(),
			LocalSubnet:       "192.168.1.0/16",
			OverlayNetwork:    "192.168.0.0/16",
			SandboxDir:        sandboxRepoDir,
			Database:          testDatabase.DBConfig(),
			HostAddress:       "10.11.12.13",
			OverlayDNSAddress: "192.168.255.254",
			ExternalDNSServer: "8.8.8.8",
			Suffix:            "potato",
		})

		ducatiCmd = exec.Command(ducatidPath, "-configFile", configFilePath)
		session, err = gexec.Start(ducatiCmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		// GinkgoParallelNode() necessary to avoid test pollution in parallel
		spaceID = fmt.Sprintf("some-space-id-%x", GinkgoParallelNode())
		networkID = spaceID
		containerID = fmt.Sprintf("some-container-id-%x", rand.Int())
		appID = fmt.Sprintf("some-app-id-%x", rand.Int())

		networkMapper := &network.FixedNetworkMapper{DefaultNetworkID: "default"}
		vni, err = networkMapper.GetVNI(spaceID)
		Expect(err).NotTo(HaveOccurred())

		sandboxName = fmt.Sprintf("vni-%d", vni)
	})

	AfterEach(func() {
		session.Interrupt()
		Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))
		Expect(containerRepo.Destroy(containerNamespace)).To(Succeed())
	})

	var serverIsAvailable = func() error {
		return VerifyTCPConnection(address)
	}

	Context("ducatid server restarts", func() {
		var (
			upSpec       models.CNIAddPayload
			downSpec     models.CNIDelPayload
			daemonClient *client.DaemonClient
			ipamResult   types.Result
		)

		BeforeEach(func() {
			Eventually(serverIsAvailable).Should(Succeed())

			daemonClient = client.New("http://"+address, http.DefaultClient)

			By("generating config and creating the request")
			upSpec = models.CNIAddPayload{
				Args:               "FOO=BAR;ABC=123",
				ContainerNamespace: containerNamespace.Name(),
				InterfaceName:      "vx-eth0",
				Network: models.NetworkPayload{
					models.Properties{
						AppID:   appID,
						SpaceID: spaceID,
					},
				},
				ContainerID: containerID,
			}

			downSpec = models.CNIDelPayload{
				InterfaceName:      "vx-eth0",
				ContainerNamespace: containerNamespace.Name(),
				ContainerID:        containerID,
			}

			By("adding the container to a network")
			var err error
			ipamResult, err = daemonClient.ContainerUp(upSpec)
			Expect(err).NotTo(HaveOccurred())
		})

		It("can delete a container that was created before ducatid was restarted", func() {
			var err error
			session.Interrupt()
			Eventually(session, DEFAULT_TIMEOUT).Should(gexec.Exit(0))
			Eventually(serverIsAvailable, DEFAULT_TIMEOUT).ShouldNot(Succeed())

			ducatiCmd = exec.Command(ducatidPath, "-configFile", configFilePath)
			session, err = gexec.Start(ducatiCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(serverIsAvailable).Should(Succeed())

			Expect(filepath.Join(sandboxRepoDir, sandboxName)).To(BeAnExistingFile())
			Expect(daemonClient.ContainerDown(downSpec)).To(Succeed())
			Expect(filepath.Join(sandboxRepoDir, sandboxName)).NotTo(BeAnExistingFile())
		})
	})
})
