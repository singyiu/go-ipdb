package validator

type SchemaValidator interface {
	//AreFieldNamesUpperCase() bool
	IsSchemaValid() bool
}
