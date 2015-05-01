package upgrade

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Setup before upgrade", func() {

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
