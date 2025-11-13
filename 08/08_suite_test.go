package quest08_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test08(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "08 Suite")
}