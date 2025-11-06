package quest03_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test03(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "03 Suite")
}