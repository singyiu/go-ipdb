package multischema

import (
	"errors"
	"github.com/multiformats/go-multihash"
	"golang.org/x/crypto/sha3"
)

// errors
var (
	ErrUnknownType      = errors.New("multischema type not supported yet")
)

// constants
const (
	IDENTITY = 0x00
	JSON = 0x11
)

// TypeToCodeMap maps the type of a hash to the code
var TypeToCodeMap = map[string]byte{
	"identity":                  IDENTITY,
	"json":                      JSON,
}

// CodeToTypeMap maps a hash code to it's type
var CodeToTypeMap = map[byte]string{
	IDENTITY:                  "identity",
	JSON:                      "json",
}

//EncodePayloadToMultiHash returns payload encoded with sha3 in MultiHash format
func EncodePayloadToMultiHash(payload []byte) ([]byte, error) {
	h := sha3.New512()
	h.Write(payload)
	rawHash := h.Sum(nil)
	return multihash.EncodeName(rawHash, "sha3")
}

//EncodeToSchemaHash returns SchemaHash in the format of <SchemaCode 1 byte><Multihash>
func EncodeToSchemaHash(schemaType string, payload []byte) ([]byte, error) {
	switch schemaType {
	case "json":
	default:
		return nil, ErrUnknownType
	}

	schemaCode, ok := TypeToCodeMap[schemaType]
	if !ok {
		return nil, ErrUnknownType
	}

	mHash, err := EncodePayloadToMultiHash(payload)
	if err != nil {
		return nil, err
	}

	var output []byte
	output = append(output, schemaCode)
	output = append(output, mHash...)
	return output, nil
}
