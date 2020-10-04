package multischema_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultischema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multischema Suite")
}
