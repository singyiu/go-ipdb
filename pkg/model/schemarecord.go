package model

type SchemaRecord struct {
	ID        string `json:"_id"` //base58.encode(SId)
	SId                  []byte               `protobuf:"bytes,1,opt,name=sId,proto3" json:"sId,omitempty"`
	Type                 string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	MetaData             []byte               `protobuf:"bytes,4,opt,name=metaData,proto3" json:"metaData,omitempty"`
	CreatedBy            []byte               `protobuf:"bytes,5,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	UpdatedAt            int64 `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"` //unixnano
}
