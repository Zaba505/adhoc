// Copyright (c) 2022 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"errors"
	"net"

	"github.com/Zaba505/adhoc/games/tictactoe/grpc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func withServeGrpcCmd() func(*viper.Viper) *cobra.Command {
	return func(v *viper.Viper) *cobra.Command {
		cmd := &cobra.Command{
			Use:   "grpc",
			Short: "Serve requests over gRPC",
			PersistentPreRunE: withPersistentPreRun(
				loadConfigFile(v),
			)(v),
			Run: func(cmd *cobra.Command, args []string) {
				addr := v.GetString("addr")
				ls, err := net.Listen("tcp", addr)
				if err != nil {
					zap.L().Fatal(
						"unexpected error when trying to listen on address",
						zap.String("addr", addr),
						zap.Error(err),
					)
					return
				}
				zap.L().Info("listening for grpc requests", zap.String("addr", addr))

				tictactoe := grpc.NewTicTacToeService(
					grpc.WithLogger(zap.L()),
				)
				err = tictactoe.Serve(cmd.Context(), ls)
				if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
					zap.L().Fatal(
						"unexpected error when serving grpc traffic",
						zap.String("addr", addr),
						zap.Error(err),
					)
					return
				}
			},
		}

		return cmd
	}
}
