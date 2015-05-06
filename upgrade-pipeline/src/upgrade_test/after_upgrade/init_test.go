package upgrade

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
)

var (
	DEFAULT_TIMEOUT       = 30 * time.Second
	CF_PUSH_TIMEOUT       = 2 * time.Minute
	BROKER_START_TIMEOUT  = 5 * time.Minute
	CF_BOOT_TIMEOUT       = 5 * time.Minute
	CF_BOOT_POLL_INTERVAL = 1 * time.Second
)

var context helpers.SuiteContext

var _ = BeforeSuite(func() {
	apiUrl := context.AdminUserContext().ApiUrl
	Eventually(func() *Session {
		return cf.Cf("api", apiUrl, "--skip-ssl-validation").Wait(DEFAULT_TIMEOUT)
	}, CF_BOOT_TIMEOUT, CF_BOOT_POLL_INTERVAL).Should(Exit(0))
})

func TestApplications(t *testing.T) {
	RegisterFailHandler(Fail)
	config := helpers.LoadConfig()
	context = helpers.NewContext(config)
	RunSpecs(t, "UpgradeTest Suite")
}
