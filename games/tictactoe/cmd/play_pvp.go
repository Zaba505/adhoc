// Copyright (c) 2023 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/Zaba505/adhoc/games/tictactoe/tictactoe"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func withPlayPvPCmd() func(*viper.Viper) *cobra.Command {
	return func(v *viper.Viper) *cobra.Command {
		cmd := &cobra.Command{
			Use:               "player-vs-player",
			Aliases:           []string{"pvp"},
			Short:             "Play a round of TicTacToe against another human player",
			PersistentPreRunE: withPersistentPreRun()(v),
			RunE: func(cmd *cobra.Command, args []string) error {
				dir := viper.GetString("game-data-dir")
				if dir == "" {
					dir = filepath.Join(os.TempDir(), "tictactoe", "data")
				}

				eng, err := tictactoe.InitEngine(dir)
				if err != nil {
					return Error{
						Cmd:   cmd,
						Cause: err,
					}
				}
				defer eng.Shutdown()

				p1Str := strings.TrimSpace(viper.GetString("player-one"))
				p2Str := strings.TrimSpace(viper.GetString("player-two"))
				if p1Str == "" || p2Str == "" {
					return Error{
						Cmd:   cmd,
						Cause: errors.New("tictactoe: player pieces can not be blank"),
					}
				}
				if utf8.RuneCountInString(p1Str) > 1 || utf8.RuneCountInString(p2Str) > 1 {
					return Error{
						Cmd:   cmd,
						Cause: errors.New("tictactoe: player pieces can only be 1 utf8 character long"),
					}
				}

				p1, _ := utf8.DecodeRuneInString(p1Str)
				p2, _ := utf8.DecodeRuneInString(p2Str)

				g := eng.NewGame(tictactoe.HumanPlayer(p1), tictactoe.HumanPlayer(p2))
				defer g.Stop()

				err = g.Run()
				if err != nil {
					return Error{
						Cmd:   cmd,
						Cause: err,
					}
				}
				return nil
			},
		}

		return cmd
	}
}
