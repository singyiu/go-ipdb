package multischema

import (
	"encoding/hex"
	"fmt"
)

type SchemaId []byte

func (sId SchemaId) Bytes() []byte {
	return []byte(sId)
}

func (sId SchemaId) String() string {
	return fmt.Sprintf("%X", sId.Bytes())
}

func SchemaIdFromHexString(str string) (SchemaId, error) {
	b, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return SchemaId(b), nil
}
