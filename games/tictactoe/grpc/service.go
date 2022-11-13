// Copyright (c) 2022 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package grpc

import (
	"context"
	"net"

	"github.com/Zaba505/adhoc/games/tictactoe/tictactoepb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ErrServerStopped = grpc.ErrServerStopped

var _ tictactoepb.TicTacToeServer = &TicTacToeService{}

// TicTacToeService
type TicTacToeService struct {
	tictactoepb.UnimplementedTicTacToeServer

	log *zap.Logger
}

// ServiceOption
type ServiceOption func(*TicTacToeService)

// WithLogger
func WithLogger(logger *zap.Logger) ServiceOption {
	return func(s *TicTacToeService) {
		s.log = logger
	}
}

// NewTicTacToeService
func NewTicTacToeService(opts ...ServiceOption) *TicTacToeService {
	s := &TicTacToeService{
		log: zap.NewNop(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Serve
func (s *TicTacToeService) Serve(ctx context.Context, ls net.Listener) error {
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	tictactoepb.RegisterTicTacToeServer(grpcServer, s)

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
