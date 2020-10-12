package threadshelper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestThreadshelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Threadshelper Suite")
}
