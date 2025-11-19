package quest12_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test12(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "12 Suite")
}