package quest06_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test06(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "06 Suite")
}