package dummy

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Redis Service", func() {
	SynchronizedBeforeSuite(func() []byte {
		// No action for node1
		return nil
	}, func(data []byte) {
		// No action for all nodes
	})

	BeforeEach(func() {
		// Nothing before each
	})

	AfterEach(func() {
		// Nothing after each
	})

	SynchronizedAfterSuite(func() {
		// Nothing for all nodes
	}, func() {
		// Nothing for node 1
	})
})
