// Copyright (c) 2022 Zaba505
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type logLevel zapcore.Level

func (l logLevel) String() string {
	return (zapcore.Level)(l).String()
}

func (l *logLevel) Set(s string) error {
	return (*zapcore.Level)(l).Set(s)
}

func (l logLevel) Type() string {
	return "Level"
}

func buildTicTacToeCmd(v *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "tictactoe",
		Short:             "",
		SilenceErrors:     true,
		PersistentPreRunE: withPersistentPreRun()(v),
	}

	lvl := logLevel(zapcore.InfoLevel)
	cmd.PersistentFlags().Var(&lvl, "log-level", "Specify log level")
	cmd.PersistentFlags().String("log-file", "stderr", "Specify log file")

	return cmd
}
