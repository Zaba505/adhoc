// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func withPlayPvCCmd() func(*viper.Viper) *cobra.Command {
	return func(v *viper.Viper) *cobra.Command {
		cmd := &cobra.Command{
			Use:               "player-vs-computer",
			Aliases:           []string{"pvc"},
			Short:             "Play a round of TicTacToe against an AI model",
			PersistentPreRunE: withPersistentPreRun()(v),
			RunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		return cmd
	}
}
