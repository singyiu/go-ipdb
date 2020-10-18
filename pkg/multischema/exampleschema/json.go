package exampleschema

import (
	"encoding/json"
	"github.com/singyiu/go-ipdb/pkg/common"
	"github.com/singyiu/go-ipdb/pkg/multischema"
	"log"
	"math/rand"
)

const (
	SensorDataSIdHexStr = "111440576651A039F55A6CA0FEF2C80EFC5298FB634E4A2205DF3A2BFA50582F9F2381FC9DCA827EAC0B7F4B712BF1EE20AF1E1B580889377115DCCBC89484A5114AA0"
	SensorDataSchema = `{
    "$schema": "http://json-schema.org/draft-07/schema",
    "$id": "http://example.com/example.json",
    "type": "object",
    "title": "The root schema",
    "description": "The root schema comprises the entire JSON document.",
    "default": {},
    "examples": [
        {
            "_id": "uuid",
            "deviceType": "indoor co2 sensor",
            "deviceId": "serial number",
            "deviceManufacturer": "Elsys",
            "deviceModel": "ERSCO2",
            "locationLat": 38.8951,
            "locationLon": -77.0364,
            "dataValue": 400.0,
            "dataUnit": "ppm",
            "timestampMilli": 1602948996254,
            "metadata": "HW v1.0, SW v1.0",
            "publishedBy": "signature"
        }
    ],
    "required": [
        "_id",
        "deviceType",
        "deviceId",
        "deviceManufacturer",
        "deviceModel",
        "locationLat",
        "locationLon",
        "dataValue",
        "dataUnit",
        "timestampMilli",
        "metadata",
        "publishedBy"
    ],
    "properties": {
        "_id": {
            "$id": "#/properties/_id",
            "type": "string",
            "title": "The _id schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "uuid"
            ]
        },
        "deviceType": {
            "$id": "#/properties/deviceType",
            "type": "string",
            "title": "The deviceType schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "indoor co2 sensor"
            ]
        },
        "deviceId": {
            "$id": "#/properties/deviceId",
            "type": "string",
            "title": "The deviceId schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "serial number"
            ]
        },
        "deviceManufacturer": {
            "$id": "#/properties/deviceManufacturer",
            "type": "string",
            "title": "The deviceManufacturer schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "Elsys"
            ]
        },
        "deviceModel": {
            "$id": "#/properties/deviceModel",
            "type": "string",
            "title": "The deviceModel schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "ERSCO2"
            ]
        },
        "locationLat": {
            "$id": "#/properties/locationLat",
            "type": "number",
            "title": "The locationLat schema",
            "description": "An explanation about the purpose of this instance.",
            "default": 0.0,
            "examples": [
                38.8951
            ]
        },
        "locationLon": {
            "$id": "#/properties/locationLon",
            "type": "number",
            "title": "The locationLon schema",
            "description": "An explanation about the purpose of this instance.",
            "default": 0.0,
            "examples": [
                -77.0364
            ]
        },
        "dataValue": {
            "$id": "#/properties/dataValue",
            "type": "number",
            "title": "The dataValue schema",
            "description": "An explanation about the purpose of this instance.",
            "default": 0.0,
            "examples": [
                400.0
            ]
        },
        "dataUnit": {
            "$id": "#/properties/dataUnit",
            "type": "string",
            "title": "The dataUnit schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "ppm"
            ]
        },
        "timestampMilli": {
            "$id": "#/properties/timestampMilli",
            "type": "integer",
            "title": "The timestampMilli schema",
            "description": "An explanation about the purpose of this instance.",
            "default": 0,
            "examples": [
                1602948996254
            ]
        },
        "metadata": {
            "$id": "#/properties/metadata",
            "type": "string",
            "title": "The metadata schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "HW v1.0, SW v1.0"
            ]
        },
        "publishedBy": {
            "$id": "#/properties/publishedBy",
            "type": "string",
            "title": "The publishedBy schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "signature"
            ]
        }
    },
    "additionalProperties": true
}`
)

func GetSensorDataSId() multischema.SchemaId {
	sId, err := multischema.SchemaIdFromHexString(SensorDataSIdHexStr)
	if err != nil {
		log.Fatalf("multischema.SchemaIdFromHexString failed, SensorDataSIdHexStr: %+v", SensorDataSIdHexStr)
	}
	return sId
}

type SensorDataStruct struct {
	Id string `json:"_id"`
	DeviceType string `json:"deviceType"`
	DeviceId string `json:"deviceId"`
	DeviceManufacturer string `json:"deviceManufacturer"`
	DeviceModel string `json:"deviceModel"`
	LocationLat float64 `json:"locationLat"`
	LocationLon float64 `json:"locationLon"`
	DataValue float64 `json:"dataValue"`
	DataUnit string `json:"dataUnit"`
	TimestampMilli int64 `json:"timestampMilli"`
	Metadata string `json:"metadata"`
	PublishedBy string `json:"publishedBy"`
}

func GetExampleSensorDataStruct() SensorDataStruct {
	return SensorDataStruct{
		Id:                 "",
		DeviceType:         "indoor co2 sensor",
		DeviceId:           "1234",
		DeviceManufacturer: "Elsys",
		DeviceModel:        "ERSCO2",
		LocationLat:        38.8951,
		LocationLon:        -77.0364,
		DataValue:          400.0 + rand.Float64()*100,
		DataUnit:           "ppm",
		TimestampMilli:     common.GetTimeStampMillisecond(),
		Metadata:           "HW v1.0, SW v1.0",
		PublishedBy:        "12D3KooWBrYBi2PCjNUH4T9pyocAApCcne6hfWZRg6LJJwzFDeq7",
	}
}

func (ds SensorDataStruct) Bytes() ([]byte, error) {
	return json.Marshal(ds)
}