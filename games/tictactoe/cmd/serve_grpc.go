// Copyright (c) 2022 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			},
		}

		return cmd
	}
}
