package json

type SchemaValidator []byte

func (v SchemaValidator) AreFieldNamesUpperCase() bool {
	return true
}

func (v SchemaValidator) IsSchemaValid() bool {
	return v.AreFieldNamesUpperCase()
}
