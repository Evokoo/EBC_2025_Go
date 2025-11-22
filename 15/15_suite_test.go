package quest15_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test15(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "15 Suite")
}