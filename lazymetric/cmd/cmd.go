package cmd

import (
	"context"
	"os"
	"os/signal"

	"go.uber.org/zap"
)

func Execute(args ...string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if len(args) == 0 {
		args = os.Args
	}

	rootCmd.SetArgs(args[1:])
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		zap.L().Fatal("unexpected error", zap.Error(err))
	}
}
