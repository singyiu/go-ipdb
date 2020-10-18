package api

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/singyiu/go-ipdb/api/pb"
	"github.com/singyiu/go-ipdb/pkg/common"
	"github.com/singyiu/go-ipdb/pkg/model"
	"github.com/singyiu/go-ipdb/pkg/multischema"
	"github.com/singyiu/go-ipdb/pkg/threadshelper"
	"github.com/singyiu/go-threads/api/client"
	"log"
	"net/http"
	"time"
)

//go:generate protoc pb/ipdb.proto -I. --go_out=plugins=grpc:.

const (
	DefaultBaseThreadIdStr = "bafk5ibp7tq5iel4cw7wtnrv27h6dj3zn543fgatnj5cb5qjmz3jtr7y"
	DefaultSignature = `12D3KooWBrYBi2PCjNUH4T9pyocAApCcne6hfWZRg6LJJwzFDeq7`
)

type Service struct {
	ThreadClientStruct *threadshelper.ClientStruct
	ctx context.Context
}

func (s *Service) Setup(ctx context.Context, baseThreadIdStr string) {
	var err error
	s.ctx = ctx
	s.ThreadClientStruct, err = threadshelper.NewClientStruct(baseThreadIdStr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		<-ctx.Done()
		defer s.ThreadClientStruct.Close()
	}()
}

func SchemaDetailToSchemaRecord(detail *pb.SchemaDetail) (model.SchemaRecord, error) {
	//check if schema is valid
	//generate schema hash
	sId, err := multischema.EncodeToSchemaId(detail.Type, detail.Data)
	if err != nil {
		return model.SchemaRecord{}, err
	}

	return model.SchemaRecord{
		ID:        sId.String(),
		SId:       sId.Bytes(),
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

	err = s.ThreadClientStruct.RegisterSchema(ctx, record)
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

func (s *Service) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishReply, error) {
	_, err := s.ThreadClientStruct.PublishPayload(ctx, req.SId, req.Payload)
	if err != nil {
		return nil, common.Errorf(err, "s.ThreadClientStruct.PublishPayload failed")
	}
	return &pb.PublishReply{
		Result:               &pb.Result{
			Code:                 http.StatusOK,
			Str:                  "ok",
		},
	}, nil
}

type Local_SubscribeServer interface {
	Send(*pb.DataRecordReply) error
}

//func (s *Service) Subscribe(req *pb.SubscribeRequest, subscriber pb.API_SubscribeServer) error {
func (s *Service) Subscribe(req *pb.SubscribeRequest, subscriber Local_SubscribeServer) error {
	sId := multischema.SchemaId(req.SId)

	go func() {
		events, err := s.ThreadClientStruct.Client.Listen(s.ctx, s.ThreadClientStruct.BaseThreadId, []client.ListenOption{{
			Type: client.ListenAll,
			Collection: sId.String(),  // Omit to receive events from all collections
		}})
		if err != nil {
			log.Printf("s.ThreadClientStruct.Client.Listen failed: %+v", err)
			return
		}

		for event := range events {
			eventSId, err := multischema.SchemaIdFromHexString(event.Action.Collection)
			if err != nil {
				log.Printf("multischema.SchemaIdFromHexString failed: %+v", err)
				break
			}
			err = subscriber.Send(&pb.DataRecordReply{
				Result:               &pb.Result{
					Code:                 http.StatusOK,
					Str:                  "ok",
				},
				DataRecord:           &pb.DataRecord{
					SId:                  eventSId.Bytes(),
					Payload:              event.Action.Instance,
				},
			})
			if err != nil {
				log.Printf("subscriber.Send failed: %+v", err)
				break
			}
			if s.ctx.Err() != nil {
				break
			}
		}
	}()

	return nil
}

/*
func (s *Service) GetSchemaDetail(context.Context, *pb.GetSchemaDetailRequest) (*pb.GetSchemaDetailReply, error) {

}



func (s *Service) Query(context.Context, *pb.QueryRequest) (*pb.QueryReply, error) {

}
*/