package cmd

import (
	"errors"
	"net"

	"github.com/Zaba505/adhoc/eventsourcing/services/eventrecorder/datastore"
	"github.com/Zaba505/adhoc/eventsourcing/services/eventrecorder/grpc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

var rootCmd = &cobra.Command{
	Use:   "eventrecorder",
	Short: "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var lvl zapcore.Level
		lvlStr := cmd.Flags().Lookup("log-level").Value.String()
		err := lvl.UnmarshalText([]byte(lvlStr))
		if err != nil {
			panic(err)
		}

		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		cfg.OutputPaths = []string{viper.GetString("log-file")}
		l, err := cfg.Build(zap.IncreaseLevel(lvl))
		if err != nil {
			panic(err)
		}

		zap.ReplaceGlobals(l)
	},
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("addr")
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

		ds := datastore.NewRedisClient(zap.L())

		s := grpc.NewEventRecorder(ds, zap.L())
		err = grpc.Serve(cmd.Context(), ls, s)
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

func init() {
	// Persistent flags
	lvl := logLevel(zapcore.InfoLevel)
	rootCmd.PersistentFlags().Var(&lvl, "log-level", "Specify log level")
	rootCmd.PersistentFlags().String("log-file", "stderr", "Specify log file")

	rootCmd.Flags().String("addr", "0.0.0.0:8080", "Address to listen for connections.")

	viper.BindPFlag("log-file", rootCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("addr", rootCmd.Flags().Lookup("addr"))
}
