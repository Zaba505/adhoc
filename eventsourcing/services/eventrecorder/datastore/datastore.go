package datastore

import (
	"context"

	pb "github.com/Zaba505/adhoc/eventsourcing/services/eventrecorder/proto"

	"go.uber.org/zap"
)

var _ Interface = &RedisClient{}

// AppendRequest
type AppendRequest struct {
	Record *pb.EventRecord
}

// AppendOnly
type AppendOnly interface {
	Append(context.Context, *AppendRequest) error
}

// ReadRequest
type ReadRequest struct {
	Filters []*pb.Filter
	Order   *pb.Order
}

// ReadOnly
type ReadOnly interface {
	Read(context.Context, *ReadRequest) (chan *pb.EventRecord, error)
}

// Interface
type Interface interface {
	AppendOnly
	ReadOnly
}

// RedisClient
type RedisClient struct {
	log *zap.Logger
}

// NewRedisClient
func NewRedisClient(log *zap.Logger) *RedisClient {
	return &RedisClient{
		log: log,
	}
}

// Append implements the AppendOnly interface by appending an EventRecord to a Redis stream.
func (ds *RedisClient) Append(ctx context.Context, req *AppendRequest) error {
	return nil
}

// Read implements the ReadOnly interface.
func (ds *RedisClient) Read(ctx context.Context, req *ReadRequest) (chan *pb.EventRecord, error) {
	return nil, nil
}
