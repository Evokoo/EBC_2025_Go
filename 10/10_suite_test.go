package quest10_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test10(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "10 Suite")
}