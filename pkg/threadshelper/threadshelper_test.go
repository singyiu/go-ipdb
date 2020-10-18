package threadshelper_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/singyiu/go-ipdb/api"
	"github.com/singyiu/go-ipdb/pkg/multischema/exampleschema"
	. "github.com/singyiu/go-ipdb/pkg/threadshelper"
	"log"
)

var _ = Describe("Threadshelper", func() {
	testBaseThreadIdStr := api.DefaultBaseThreadIdStr
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
	XDescribe("CreateBaseDb", func() {
		It("should create the db correctly", func() {
			dbInfo, err := testClientStruct.CreateBaseDb()
			Expect(err).Should(BeNil())
			Expect(dbInfo).ShouldNot(BeNil())
			log.Printf("dbInfo: %+v", dbInfo)
		})
	})
	XDescribe("CreateBaseCollection", func() {
		It("should create the collection correctly", func() {
			err := testClientStruct.CreateBaseCollection()
			Expect(err).Should(BeNil())
		})
	})
	Describe("PublishPayload", func() {
		It("should publish the payload correctly", func() {
			dataStruct := exampleschema.GetExampleSensorDataStruct()
			payload, err := dataStruct.Bytes()
			Expect(err).Should(BeNil())
			_, err = testClientStruct.PublishPayload(context.Background(), exampleschema.GetSensorDataSId(), payload)
			Expect(err).Should(BeNil())
		})
	})
})
