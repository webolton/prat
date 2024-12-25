package bootstrap_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("environment.go", func() {

	Describe("getEnvironment", func() {
		Context("when no environment is set", func() {
			It("should be a novel", func() {
				Expect("true").To(Equal(true))
			})
		})
	})
})
