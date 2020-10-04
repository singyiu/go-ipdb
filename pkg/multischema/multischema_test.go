package multischema_test

import (
	"encoding/hex"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/singyiu/go-ipdb/pkg/multischema"
)

var _ = Describe("Multischema", func() {
	testPayload := []byte(`{"type": "string"}`)
	expectedMHash, _ := hex.DecodeString("1440A86E4455994365B7B4216834B5ECF30E0AC1872E3FB768454CD30566EDD54CF5DD983736E9343B2CEF88FB41F9D97DB7AC907C40E24D4AED968AA2F0D29EEC12")
	expectedSchemaHash, _ := hex.DecodeString("111440A86E4455994365B7B4216834B5ECF30E0AC1872E3FB768454CD30566EDD54CF5DD983736E9343B2CEF88FB41F9D97DB7AC907C40E24D4AED968AA2F0D29EEC12")

	Describe("EncodePayloadToMultiHash", func() {
		mHash, err := EncodePayloadToMultiHash(testPayload)
		Expect(err).Should(BeNil())
		Expect(mHash).Should(Equal(expectedMHash))
	})
	Describe("GetHashFromJsonSchema", func() {
		schemaHash, err := EncodeToSchemaHash("json", testPayload)
		Expect(err).Should(BeNil())
		Expect(schemaHash).Should(Equal(expectedSchemaHash))
	})
})
