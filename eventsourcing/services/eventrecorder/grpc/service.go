package grpc

import (
	"context"
	"net"

	"github.com/Zaba505/adhoc/eventsourcing/services/eventrecorder/datastore"
	pb "github.com/Zaba505/adhoc/eventsourcing/services/eventrecorder/proto"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ErrServerStopped = grpc.ErrServerStopped

// EventRecorder
type EventRecorder struct {
	pb.UnimplementedEventRecorderServer

	ds datastore.Interface

	log *zap.Logger
}

// NewEventRecorder
func NewEventRecorder(ds datastore.Interface, log *zap.Logger) *EventRecorder {
	return &EventRecorder{
		log: log,
	}
}

// Serve instantiates the gRPC server and registers the EventRecorder service with it.
func Serve(ctx context.Context, ls net.Listener, s *EventRecorder) error {
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterEventRecorderServer(grpcServer, s)

	errCh := make(chan error, 1)
	go func() {
		defer close(errCh)
		err := grpcServer.Serve(ls)
		errCh <- err
	}()

	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	select {
	case <-cctx.Done():
		grpcServer.GracefulStop()
		<-errCh
		return ErrServerStopped
	case err := <-errCh:
		cancel()
		return err
	}
}

// RecordEvent
func (s *EventRecorder) RecordEvent(ctx context.Context, req *pb.RecordEventRequest) (*pb.RecordEventResponse, error) {
	s.log.Info("recording new event", withNewRecordInfo(req.Record)...)

	// TODO: put record into datastore

	// TODO: If it fails return
	// TODO: If it succeeds publish notification

	status := &pb.RecordEventResponse_Success{
		Success: &pb.RecordEventSuccess{
			RecordId: "1234",
		},
	}

	resp := &pb.RecordEventResponse{
		Status: status,
	}

	return resp, nil
}

// GetRecord
func (s *EventRecorder) GetRecord(ctx context.Context, req *pb.GetRecordRequest) (*pb.GetRecordResponse, error) {
	s.log.Info("retrieving record", zap.String("record_id", req.RecordId))

	// TODO: get record from datastore

	record := &pb.EventRecord{
		TraceId:   "123-456-789",
		RecordId:  "1234",
		EventType: "test",
	}
	s.log.Info("retrieved record", withRecordInfo(record)...)

	resp := &pb.GetRecordResponse{
		Record: record,
	}
	return resp, nil
}

// SliceRecords
func (s *EventRecorder) SliceRecords(req *pb.SliceRecordsRequest, stream pb.EventRecorder_SliceRecordsServer) error {
	s.log.Info("retrieveing records")

	// TODO: read records from datastore

	return nil
}

func withNewRecordInfo(record *pb.NewEventRecord) []zapcore.Field {
	return []zapcore.Field{
		zap.String("trace_id", record.TraceId),
		zap.String("event_type", record.EventType),
	}
}

func withRecordInfo(record *pb.EventRecord) []zapcore.Field {
	return []zapcore.Field{
		zap.String("trace_id", record.TraceId),
		zap.String("record_id", record.RecordId),
		zap.String("event_type", record.EventType),
	}
}
