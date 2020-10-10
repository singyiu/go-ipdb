package api

import (
	"context"
	pb "github.com/singyiu/go-ipdb/api/pb"
	"github.com/singyiu/go-ipdb/pkg/multischema"
	"net/http"
)

//go:generate protoc pb/ipdb.proto -I. --go_out=plugins=grpc:.

type Service struct {
}

func (s *Service) RegisterSchema(ctx context.Context, req *pb.RegisterSchemaRequest) (*pb.RegisterSchemaReply, error) {
	//check if schema is valid
	//generate schema hash
	schemaHash, err := multischema.EncodeToSchemaHash(req.SchemaDetail.SchemaType, req.SchemaDetail.SchemaData)
	if err != nil {
		return nil, err
	}

	//check if schemaHash already exist in db


	//register schemaHash in db

	//return schema hash
	return &pb.RegisterSchemaReply{
		Result:               &pb.Result{
			Code:                 http.StatusOK,
			Str:                  "ok",
		},
		SchemaHash:           schemaHash,
	}, nil
}

/*
func (s *Service) GetSchemaDetail(context.Context, *pb.GetSchemaDetailRequest) (*pb.GetSchemaDetailReply, error) {

}

func (s *Service) Publish(context.Context, *pb.PublishRequest) (*pb.PublishReply, error) {

}

func (s *Service) Subscribe(*pb.SubscribeRequest, pb.API_SubscribeServer) error {

}

func (s *Service) Query(context.Context, *pb.QueryRequest) (*pb.QueryReply, error) {

}
*/