package upgrade

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Do things after upgrade", func() {
	It("binds the unbound instance", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			Expect(cf.Cf("bind-service", "upgrade-service-broker", "bind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("bind-me"))
			Expect(services.Out.Contents()).To(MatchRegexp("bind-me\\s+fake-service\\s+fake-plan\\s+upgrade-service-broker"))

			service := cf.Cf("service", "bind-me").Wait(DEFAULT_TIMEOUT)
			Expect(service.Out.Contents()).To(ContainSubstring("Status: create succeeded"))
		})
	})

	Context("updating the instance", func() {
		It("updates the plan", func() {
			cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
				Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

				Expect(cf.Cf("update-service", "-p", "fake-plan-2", "update-my-plan").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

				services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
				Expect(services.Out.Contents()).To(ContainSubstring("update-my-plan"))
				Expect(services.Out.Contents()).To(ContainSubstring("fake-plan-2"))

				service := cf.Cf("service", "update-my-plan").Wait(DEFAULT_TIMEOUT)
				Expect(service.Out.Contents()).To(ContainSubstring("Status: update succeeded"))
			})
		})

		It("updates the name", func() {
			cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
				Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

				Expect(cf.Cf("rename-service", "update-my-name", "my-name-is-updated").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

				services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
				Expect(services.Out.Contents()).To(ContainSubstring("my-name-is-updated"))

				service := cf.Cf("service", "my-name-is-updated").Wait(DEFAULT_TIMEOUT)
				Expect(service.Out.Contents()).To(ContainSubstring("Status: update succeeded"))
			})
		})
	})

	It("deletes an instance", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			Expect(cf.Cf("delete-service", "delete-me", "-f").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).NotTo(ContainSubstring("delete-me"))
		})
	})

	It("unbinds an instance from an app", func() {
		cf.AsUser(context.AdminUserContext(), DEFAULT_TIMEOUT, func() {
			Expect(cf.Cf("target", "-o", "upgrade-org", "-s", "upgrade-space").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			Expect(cf.Cf("unbind-service", "upgrade-service-broker", "unbind-me").Wait(DEFAULT_TIMEOUT)).To(Exit(0))

			services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
			Expect(services.Out.Contents()).To(ContainSubstring("unbind-me"))
			Expect(services.Out.Contents()).NotTo(MatchRegexp("unbind-me\\s+fake-service\\s+fake-plan\\s+upgrade-service-broker"))

			service := cf.Cf("service", "unbind-me").Wait(DEFAULT_TIMEOUT)
			Expect(service.Out.Contents()).To(ContainSubstring("Status: create succeeded"))
		})
	})
})
