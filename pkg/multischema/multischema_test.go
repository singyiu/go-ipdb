package multischema_test

import (
	"encoding/hex"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/singyiu/go-ipdb/pkg/multischema"
	"github.com/singyiu/go-ipdb/pkg/multischema/exampleschema"
)

var _ = Describe("Multischema", func() {
	testPayload := []byte(`{"type": "string"}`)
	testSIdHexStr := "111440A86E4455994365B7B4216834B5ECF30E0AC1872E3FB768454CD30566EDD54CF5DD983736E9343B2CEF88FB41F9D97DB7AC907C40E24D4AED968AA2F0D29EEC12"
	expectedMHash, _ := hex.DecodeString("1440A86E4455994365B7B4216834B5ECF30E0AC1872E3FB768454CD30566EDD54CF5DD983736E9343B2CEF88FB41F9D97DB7AC907C40E24D4AED968AA2F0D29EEC12")

	Describe("EncodePayloadToMultiHash", func() {
		It("should encode the payload correctly", func() {
			mHash, err := EncodePayloadToMultiHash(testPayload)
			Expect(err).Should(BeNil())
			Expect(mHash).Should(Equal(expectedMHash))
		})
	})
	Describe("GetHashFromJsonSchema", func() {
		It("should be encoded to schema hash correctly", func() {
			expectedSId, err := SchemaIdFromHexString(testSIdHexStr)
			Expect(err).Should(BeNil())
			sId, err := EncodeToSchemaId("json", testPayload)
			Expect(err).Should(BeNil())
			Expect(sId).Should(Equal(expectedSId))
			sId, err = EncodeToSchemaId("json", []byte(exampleschema.SensorDataSchema))
			Expect(err).Should(BeNil())
			//log.Printf("sId: %s\n", sId.String())
			Expect(err).Should(BeNil())
			Expect(sId).Should(Equal(exampleschema.GetSensorDataSId()))
		})
	})
})
