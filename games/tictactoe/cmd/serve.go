// Copyright (c) 2022 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func withServeCmd(subcommandBuilders ...func(*viper.Viper) *cobra.Command) func(*viper.Viper) *cobra.Command {
	return func(v *viper.Viper) *cobra.Command {
		cmd := &cobra.Command{
			Use:   "serve",
			Short: "Serve a specific Tic-Tac-Toe game service API",
			PersistentPreRunE: withPersistentPreRun(
				loadConfigFile(v),
			)(v),
		}

		cmd.PersistentFlags().String("addr", "0.0.0.0:8080", "Address to listen for connections.")
		cmd.PersistentFlags().String("config-file", "", "Specify config file")

		for _, b := range subcommandBuilders {
			cmd.AddCommand(b(v))
		}

		return cmd
	}
}
