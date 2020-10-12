package api

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mr-tron/base58"
	pb "github.com/singyiu/go-ipdb/api/pb"
	"github.com/singyiu/go-ipdb/pkg/model"
	"github.com/singyiu/go-ipdb/pkg/multischema"
	"github.com/singyiu/go-ipdb/pkg/threadshelper"
	"log"
	"net/http"
	"time"
)

//go:generate protoc pb/ipdb.proto -I. --go_out=plugins=grpc:.

type Service struct {
}

func Setup() {
	err := threadshelper.MyClientStruct.Setup()
	if err != nil {
		log.Fatal(err)
	}
}

func SchemaDetailToSchemaRecord(detail *pb.SchemaDetail) (model.SchemaRecord, error) {
	//check if schema is valid
	//generate schema hash
	sId, err := multischema.EncodeToSchemaHash(detail.Type, detail.Data)
	if err != nil {
		return model.SchemaRecord{}, err
	}

	return model.SchemaRecord{
		ID:        base58.Encode(sId),
		SId:       sId,
		Type:      detail.Type,
		Data:      detail.Data,
		MetaData:  detail.MetaData,
		CreatedBy: detail.CreatedBy,
		UpdatedAt: time.Now().UnixNano(),
	}, nil
}

func SchemaRecordToSchemaDetail(record model.SchemaRecord) pb.SchemaDetail {
	t := time.Unix(0, record.UpdatedAt)
	return pb.SchemaDetail{
		SId:                  record.SId,
		Type:                 record.Type,
		Data:                 record.Data,
		MetaData:             record.MetaData,
		CreatedBy:            record.CreatedBy,
		UpdatedAt:            &timestamp.Timestamp{
			Seconds: int64(t.Second()),
			Nanos: int32(t.Nanosecond()),
		},
	}
}

func (s *Service) RegisterSchema(ctx context.Context, req *pb.RegisterSchemaRequest) (*pb.RegisterSchemaReply, error) {
	record, err := SchemaDetailToSchemaRecord(req.SchemaDetail)
	if err != nil {
		return nil, err
	}

	err = threadshelper.MyClientStruct.RegisterSchema(record)
	if err != nil {
		return nil, err
	}

	detail := SchemaRecordToSchemaDetail(record)
	return &pb.RegisterSchemaReply{
		Result:               &pb.Result{
			Code:                 http.StatusOK,
			Str:                  "ok",
		},
		SchemaDetail:         &detail,
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