package quest13_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test13(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "13 Suite")
}