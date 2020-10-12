package threadshelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/singyiu/go-ipdb/pkg/threadshelper"
	"log"
)

var _ = Describe("Threadshelper", func() {
	testBaseThreadIdStr := "bafk5ibp7tq5iel4cw7wtnrv27h6dj3zn543fgatnj5cb5qjmz3jtr7y"
	testKeyFileName := "test.pem"

	var testClientStruct *ClientStruct
	Describe("SavePrivateKeyToPemFile", func() {
		It("should save the private key to a local pem file correctly", func() {
			key, err := SavePrivateKeyToPemFile(testKeyFileName)
			Expect(err).Should(BeNil())
			Expect(key).ShouldNot(BeNil())
		})
	})
	Describe("LoadPrivateKeyFromPemFile", func() {
		It("should load the private key correctly", func() {
			key, err := LoadPrivateKeyFromPemFile(testKeyFileName)
			Expect(err).Should(BeNil())
			Expect(key).ShouldNot(BeNil())
		})
	})
	Describe("Setup", func() {
		It("should return the thread client correctly", func() {
			var err error
			testClientStruct, err = NewClientStruct(testBaseThreadIdStr)
			Expect(err).Should(BeNil())
			Expect(testClientStruct.Client).ShouldNot(BeNil())
		})
	})
	XDescribe("CreateDb", func() {
		It("should create the db correctly", func() {
			dbInfo, err := testClientStruct.CreateDb()
			Expect(err).Should(BeNil())
			Expect(dbInfo).ShouldNot(BeNil())
			log.Printf("dbInfo: %+v", dbInfo)
		})
	})
	XDescribe("CreateCollection", func() {
		It("should create the collection correctly", func() {
			err := testClientStruct.CreateCollection()
			Expect(err).Should(BeNil())
		})
	})
})
