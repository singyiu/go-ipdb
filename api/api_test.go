package api_test

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/singyiu/go-ipdb/api"
	pb "github.com/singyiu/go-ipdb/api/pb"
	"github.com/singyiu/go-ipdb/pkg/multischema"
	"github.com/singyiu/go-ipdb/pkg/multischema/exampleschema"
	"net/http"
)

const (
	baseThreadIdStr = api.DefaultBaseThreadIdStr
	testSignature = api.DefaultSignature
)

var _ = Describe("Api", func() {
	Describe("service", func() {
		testService := api.Service{}
		testService.Setup(context.Background(), baseThreadIdStr)
		XDescribe("RegisterSchema", func() {
			testRequest := pb.RegisterSchemaRequest{
				Requester:            &pb.Identity{
					Signature:            []byte(testSignature),
				},
				SchemaDetail:         &pb.SchemaDetail{
					SId:                  nil,
					Type:                 "json",
					Data:                 []byte(exampleschema.SensorDataSchema),
					MetaData:             nil,
					CreatedBy:            []byte(testSignature),
					UpdatedAt:            nil,
				},
			}
			It("Should register the schema correctly", func() {
				reply, err := testService.RegisterSchema(context.Background(), &testRequest)
				Expect(err).Should(BeNil())
				Expect(reply).ShouldNot(BeNil())
				//log.Printf("reply: %+v\nsId: %X", reply, reply.SchemaDetail.SId)
				Expect(multischema.SchemaId(reply.SchemaDetail.SId)).Should(Equal(exampleschema.GetSensorDataSId()))
			})
		})
		Describe("Publish", func() {
			It("Should publish the data correctly", func() {
				examplePayload, err := exampleschema.GetExampleSensorDataStruct().Bytes()
				Expect(err).Should(BeNil())
				req := pb.PublishRequest{
					Requester:            &pb.Identity{
						Signature:            []byte(testSignature),
					},
					SId:                  exampleschema.GetSensorDataSId(),
					Payload:              examplePayload,
				}
				reply, err := testService.Publish(context.Background(), &req)
				Expect(err).Should(BeNil())
				Expect(reply.Result.Code).Should(Equal(int32(http.StatusOK)))
			})
		})
	})
})
