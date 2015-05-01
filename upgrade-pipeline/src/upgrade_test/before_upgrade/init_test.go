package upgrade

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
	"github.com/cloudfoundry/cf-acceptance-tests/services"
)

var (
	DEFAULT_TIMEOUT       = 30 * time.Second
	CF_PUSH_TIMEOUT       = 2 * time.Minute
	BROKER_START_TIMEOUT  = 5 * time.Minute
	CF_BOOT_TIMEOUT       = 5 * time.Minute
	CF_BOOT_POLL_INTERVAL = 1 * time.Second
	context               helpers.SuiteContext
	broker                services.ServiceBroker
	serviceBrokerPath     = "../../github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker"
)

var _ = BeforeSuite(func() {

	apiUrl := context.AdminUserContext().ApiUrl
	Eventually(func() *Session {
		return cf.Cf("api", apiUrl, "--skip-ssl-validation").Wait(DEFAULT_TIMEOUT)
	}, CF_BOOT_TIMEOUT, CF_BOOT_POLL_INTERVAL).Should(Exit(0))

	cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {

		Expect(cf.Cf("create-org", "upgrade-org").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		Expect(cf.Cf("create-space", "upgrade-space", "-o", "upgrade-org").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

		broker = services.NewServiceBroker("upgrade-service-broker", serviceBrokerPath, context)
		broker.Push()
		broker.Create()

		Expect(cf.Cf("enable-service-access", "fake-service").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

		Expect(cf.Cf("create-service", "fake-service", "fake-plan", "bind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		Expect(cf.Cf("create-service", "fake-service", "fake-plan", "update-my-plan").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		Expect(cf.Cf("create-service", "fake-service", "fake-plan", "update-my-name").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		Expect(cf.Cf("create-service", "fake-service", "fake-plan", "delete-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		Expect(cf.Cf("create-service", "fake-service", "fake-plan", "unbind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

		Expect(cf.Cf("bind-service", "upgrade-service-broker", "unbind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
	})
})

func TestApplications(t *testing.T) {
	RegisterFailHandler(Fail)
	config := helpers.LoadConfig()
	context = helpers.NewContext(config)
	RunSpecs(t, "UpgradeTest Suite")
}
