package upgrade

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/services"

	 . "github.com/onsi/ginkgo"
	 . "github.com/onsi/gomega"
	 . "github.com/onsi/gomega/gexec"
)


var _ = Describe("Setup before upgrade", func() {
	var broker services.ServiceBroker
	var serviceBrokerPath = "../github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker"

	BeforeSuite(func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("create-org", "upgrade-org", "-q", "paid").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("create-space", "upgrade-space", "-o", "upgrade-org").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			broker = services.NewServiceBroker("upgrade-service-broker", serviceBrokerPath, context)
			broker.Push()
			broker.Create()

			Expect(cf.Cf("enable-service-access", "fake-service").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			
			Expect(cf.Cf("create-service", "fake-service", "fake-plan", "bind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("create-service", "fake-service", "fake-plan", "update-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("create-service", "fake-service", "fake-plan", "delete-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
			Expect(cf.Cf("create-service", "fake-service", "fake-plan", "unbind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			Expect(cf.Cf("bind-service", "upgrade-service-broker", "unbind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		})
	})

	It("creates a service broker", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			brokers := cf.Cf("service-brokers").Wait(DEFAULT_TIMEOUT)
			Expect(brokers.Out.Contents()).To(ContainSubstring("upgrade-service-broker"))

			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("marketplace").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("fake-service"))
		})
	})

	It("creates an unbound instance that will later be bound", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("bind-me"))
		})
	})
	It("creates an unbound instance that will later be updated", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("update-me"))
		})
	})
	It("creates an unbound instance that will later be deleted", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("delete-me"))
		})
	})
	It("creates an bound instance that will later be unbound", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("unbind-me"))
			Expect(services.Out.Contents()).To(ContainSubstring("upgrade-service-broker"))
		})
	})
})
