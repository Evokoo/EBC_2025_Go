package quest05_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test05(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "05 Suite")
}