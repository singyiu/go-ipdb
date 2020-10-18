package main

import (
	"context"
	"encoding/json"
	"github.com/namsral/flag"
	"github.com/reactivex/rxgo/v2"
	"github.com/singyiu/go-ipdb/api"
	pb "github.com/singyiu/go-ipdb/api/pb"
	"github.com/singyiu/go-ipdb/pkg/common"
	"github.com/singyiu/go-ipdb/pkg/multischema"
	"github.com/singyiu/go-ipdb/pkg/multischema/exampleschema"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var MyService api.Service

var publishPeriod = time.Duration(5) * time.Second

type LogSubscriber struct {}

func (ls LogSubscriber) Send(reply *pb.DataRecordReply) error {
	if reply.Result.Str != "ok" {
		log.Fatalf("Got error reply: %+v", reply.Result)
	}
	sId := multischema.SchemaId(reply.DataRecord.SId)
	if sId.String() == exampleschema.SensorDataSIdHexStr {
		dataStruct := exampleschema.SensorDataStruct{}
		err := json.Unmarshal(reply.DataRecord.Payload, &dataStruct)
		if err != nil {
			return common.Errorf(err, "json.Unmarshal failed")
		}
		log.Printf("sId: %s\npayload: %+v\n", sId.String(), dataStruct)
	} else {
		log.Printf("sId: %s\npayload: %+v\n", sId.String(), reply.DataRecord.Payload)
	}
	return nil
}

func GetMapFuncIntervalToPublish(sId multischema.SchemaId, payload []byte) func(ctx context.Context, _ interface{})(interface{}, error) {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		var outputPayload []byte
		if sId.String() == exampleschema.SensorDataSIdHexStr {
			dataStruct := exampleschema.GetExampleSensorDataStruct()
			outputPayload, _ = dataStruct.Bytes()
		} else {
			outputPayload = payload
		}

		req := pb.PublishRequest{
			Requester:            &pb.Identity{
				Signature:            []byte(api.DefaultSignature),
			},
			SId:                  sId.Bytes(),
			Payload:              outputPayload,
		}
		log.Printf("Requester: %s\nsId: %s\nPayload size: %d\n", req.Requester, sId.String(), len(outputPayload))

		reply, err := MyService.Publish(ctx, &req)
		if err != nil {
			log.Fatalf("MyService.Subscribe failed: %+v", err)
		}
		log.Printf("reply: %+v", reply.Result.Str)
		return reply, nil
	}
}

func ProcessPublish(ctx context.Context, sIdStr string, payload []byte) {
	if len(sIdStr) == 0 {
		log.Fatalf("empty sId")
	}
	sId, err := multischema.SchemaIdFromHexString(sIdStr)
	if err != nil {
		log.Fatalf("multischema.SchemaIdFromHexString failed: %+v\n sIdStr: %+v", err, sIdStr)
	}
	log.Printf("Start publishing to sId: %s", sId.String())
	intervalProducer := rxgo.Interval(rxgo.WithDuration(publishPeriod))
	publishCh := intervalProducer.
		Map(GetMapFuncIntervalToPublish(sId, payload)).
		Observe(rxgo.WithErrorStrategy(rxgo.ContinueOnError))

	for {
		select {
		case _, ok := <-publishCh:
			if !ok {
				//N.B. Fatal if the channel is closed, otherwise the select statement will keep on reading the zero value from the closed channel repeatedly
				log.Fatal("publishCh closed")
			}
		case <-ctx.Done():
			return
		}
	}
}

func ProcessSubscribe(_ context.Context, sIdStr string) {
	if len(sIdStr) == 0 {
		log.Fatalf("empty sId")
	}
	sId, err := multischema.SchemaIdFromHexString(sIdStr)
	if err != nil {
		log.Fatalf("multischema.SchemaIdFromHexString failed: %+v\n sIdStr: %+v", err, sIdStr)
	}
	log.Printf("Start subscribing to sId: %s", sId.String())
	req := pb.SubscribeRequest{
		Requester:            &pb.Identity{
			Signature:            []byte(api.DefaultSignature),
		},
		SId:                  sId.Bytes(),
	}
	log.Printf("SubscribeRequest: %+v", req)
	err = MyService.Subscribe(&req, LogSubscriber{})
	if err != nil {
		log.Fatalf("MyService.Subscribe failed: %+v", err)
	}
}

func Start(ctx context.Context) {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "IPDB", 0)
	cmdStr := fs.String("cmd", "publish", "target command")
	sIdStr := fs.String("sId", exampleschema.SensorDataSIdHexStr, "target sId")
	payloadStr := fs.String("payload", "{}", "payload")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	MyService.Setup(ctx, api.DefaultBaseThreadIdStr)

	switch *cmdStr {
	case "publish":
		ProcessPublish(ctx, *sIdStr, []byte(*payloadStr))
	case "subscribe":
		ProcessSubscribe(ctx, *sIdStr)
	default:
		log.Fatalf("unsupported cmd: %+v", cmdStr)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Start(ctx)
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait
	cancel()
}
