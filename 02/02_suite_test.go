package quest02_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test02(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "02 Suite")
}