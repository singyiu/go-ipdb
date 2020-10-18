package model

import (
	"github.com/alecthomas/jsonschema"
	"github.com/singyiu/go-ipdb/pkg/common"
	"github.com/singyiu/go-threads/util"
)

type SchemaRecord struct {
	ID                   string `json:"_id"`
	SId                  []byte               `protobuf:"bytes,1,opt,name=sId,proto3" json:"sId,omitempty"`
	Type                 string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	MetaData             []byte               `protobuf:"bytes,4,opt,name=metaData,proto3" json:"metaData,omitempty"`
	CreatedBy            []byte               `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedAt            int64 `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"` //unixnano
}

func (sr SchemaRecord) GetJsonSchema() (*jsonschema.Schema, error) {
	if sr.Type != "json" {
		return nil, common.Errorf(nil, "schema type conversion not supported")
	}
	return util.SchemaFromSchemaString(string(sr.Data)), nil
}