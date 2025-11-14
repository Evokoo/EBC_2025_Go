package quest09_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test09(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "09 Suite")
}