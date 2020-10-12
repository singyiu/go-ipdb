package threadshelper

import (
	"context"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"github.com/alecthomas/jsonschema"
	"github.com/libp2p/go-libp2p-core/crypto"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/singyiu/go-ipdb/pkg/common"
	"github.com/singyiu/go-ipdb/pkg/model"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/go-threads/util"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
)

const (
	CollectionNameSchemaRecord = "SchemaRecord"
)

type ClientStruct struct {
	BaseThreadId thread.ID
	Client *client.Client
	Identity thread.Identity
	Token thread.Token
}

//test bafk5ibp7tq5iel4cw7wtnrv27h6dj3zn543fgatnj5cb5qjmz3jtr7y
//bafkqin6miovgsfvgibzxzoyu6fxxgixmmgj2hbh746nirh4lah3cmxi
func NewClientStruct(baseThreadIdStr string) (*ClientStruct, error) {
	var err error
	cs := ClientStruct{}

	//create a random threadId if baseThreadIdStr is empty
	if len(baseThreadIdStr) == 0 {
		baseThreadIdStr = thread.NewIDV1(thread.Raw, 32).String()
	}

	cs.BaseThreadId, err = thread.Decode(baseThreadIdStr)
	if err != nil {
		return nil, common.Errorf(err, "thread.Decode failed")
	}
	log.Printf("BaseThreadId: %+v", cs.BaseThreadId)

	addr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/6006")
	if err != nil {
		return nil, common.Errorf(err, "ma.NewMultiaddr failed")
	}
	log.Printf("addr: %+v", addr)
	target, err := util.TCPAddrFromMultiAddr(addr)
	if err != nil {
		return nil, common.Errorf(err, "util.TCPAddrFromMultiAddr failed")
	}
	log.Printf("target: %+v", target)
	cs.Client, err = client.NewClient(target, grpc.WithInsecure())
	if err != nil {
		return nil, common.Errorf(err, "client.NewClient failed")
	}

	dbs, err := cs.Client.ListDBs(context.Background())
	if err != nil {
		return nil, common.Errorf(err, "Client.ListDBs failed")
	}
	log.Printf("dbs: %+v", dbs)

	key, err := LoadPrivateKeyFromPemFile("privatekey.pem")
	if err != nil {
		return nil, err
	}
	cs.Identity = thread.NewLibp2pIdentity(key)
	cs.Token, err = cs.Client.GetToken(context.Background(), cs.Identity)

	return &cs, nil
}

func SavePrivateKeyToPemFile(fileName string) (crypto.PrivKey, error) {
	privateKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return privateKey, err
	}
	privateKeyBytes, err := crypto.MarshalPrivateKey(privateKey)
	if err != nil {
		return privateKey, err
	}
	keyFile, err := os.Create(fileName)
	if err != nil {
		return privateKey, err
	}
	if err := pem.Encode(keyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes}); err != nil {
		return privateKey, err
	}
	if err := keyFile.Close(); err != nil {
		return privateKey, err
	}
	return privateKey, nil
}

func LoadPrivateKeyFromPemFile(fileName string) (crypto.PrivKey, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(data)
	if pemBlock.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("invalid input file")
	}
	return crypto.UnmarshalPrivateKey(pemBlock.Bytes)
}

func (cs *ClientStruct) CreateDb() (db.Info, error) {
	log.Printf("Creating DB with threadID: %+v", cs.BaseThreadId)
	err := cs.Client.NewDB(context.Background(), cs.BaseThreadId)
	if err != nil {
		return db.Info{}, common.Errorf(err, "Client.NewDB failed")
	}
	return cs.Client.GetDBInfo(context.Background(), cs.BaseThreadId)
}

func (cs *ClientStruct) CreateCollection() error {
	log.Printf("Creating SchemaRecord collection")
	reflector := jsonschema.Reflector{}
	schemaRecordSchema := reflector.Reflect(&model.SchemaRecord{}) // Generate a JSON Schema from a struct
	//log.Printf("schemaDetailSchema: %+v %+v", schemaRecordSchema.Type, schemaRecordSchema.Definitions)

	err := cs.Client.NewCollection(context.Background(), cs.BaseThreadId, db.CollectionConfig{
	//err := cs.Client.UpdateCollection(context.Background(), cs.BaseThreadId, db.CollectionConfig{
		Name:    CollectionNameSchemaRecord,
		Schema:  schemaRecordSchema,
		Indexes: []db.Index{{
			Path:   "sId", // Value matches json tags
			Unique: true, // Create a unique index on "name"
		}},
	})

	return err
}

func (cs *ClientStruct) RegisterSchema(record model.SchemaRecord) error {
	_, err := cs.Client.Create(context.Background(), cs.BaseThreadId, CollectionNameSchemaRecord, client.Instances{&record})
	return err
}