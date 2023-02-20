// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func withPlayCvCCmd() func(*viper.Viper) *cobra.Command {
	return func(v *viper.Viper) *cobra.Command {
		cmd := &cobra.Command{
			Use:               "computer-vs-computer",
			Aliases:           []string{"cvc"},
			Short:             "Pit two AI models against each other in a round of TicTacToe",
			PersistentPreRunE: withPersistentPreRun()(v),
			RunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		return cmd
	}
}
