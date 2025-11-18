package quest11_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test11(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "11 Suite")
}