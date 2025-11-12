package quest07_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test07(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "07 Suite")
}