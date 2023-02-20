// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func withPlayCmd(subcommandBuilders ...func(*viper.Viper) *cobra.Command) func(*viper.Viper) *cobra.Command {
	return func(v *viper.Viper) *cobra.Command {
		cmd := &cobra.Command{
			Use:               "play",
			Short:             "Play a round of tictactoe",
			PersistentPreRunE: withPersistentPreRun()(v),
		}

		for _, b := range subcommandBuilders {
			cmd.AddCommand(b(v))
		}

		dir, _ := os.UserCacheDir()
		if dir != "" {
			dir = filepath.Join(dir, "tictactoe", "data")
		}

		cmd.PersistentFlags().StringP("game-data-dir", "d", dir, "Specify a directory for saving game data.")
		cmd.PersistentFlags().String("player-one", "X", "Specify the piece used to represent player one.")
		cmd.PersistentFlags().String("player-two", "O", "Specify the piece used to represent player one.")

		viper.BindPFlag("game-data-dir", cmd.PersistentFlags().Lookup("game-data-dir"))
		viper.BindPFlag("player-one", cmd.PersistentFlags().Lookup("player-one"))
		viper.BindPFlag("player-two", cmd.PersistentFlags().Lookup("player-two"))

		return cmd
	}
}
