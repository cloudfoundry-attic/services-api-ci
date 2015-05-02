package upgrade

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
)

var (
	DEFAULT_TIMEOUT      = 30 * time.Second
	CF_PUSH_TIMEOUT      = 2 * time.Minute
	BROKER_START_TIMEOUT = 5 * time.Minute
)

var context helpers.SuiteContext

func TestApplications(t *testing.T) {
	RegisterFailHandler(Fail)
	config := helpers.LoadConfig()
	context = helpers.NewContext(config)
	RunSpecs(t, "UpgradeTest Suite")
}
