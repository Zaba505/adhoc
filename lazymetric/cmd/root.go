package cmd

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/Zaba505/adhoc/lazymetric/metric"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lazymetric",
	Short: "Construct a distance matrix from a stream of points.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stream := getPointStream(cmd.InOrStdin(), args)
		m := metric.Compute(stream, metric.Euclidean)

		err := displayMatrix(cmd.OutOrStdout(), m)
		if err != nil {
			zap.L().Fatal("failed to display matrix", zap.Error(err))
		}
	},
}

func getPointStream(r io.Reader, args []string) <-chan float64 {
	stream := make(chan float64)
	src := func() {
		for _, arg := range args {
			x, err := strconv.ParseFloat(strings.TrimSpace(arg), 64)
			if err != nil {
				continue
			}
			stream <- x
		}
	}
	if args[0] == "-" {
		br := bufio.NewReader(r)
		src = func() {
			for {
				s, err := br.ReadString(' ')
				if err == io.EOF {
					return
				}
				if err != nil {
					return
				}
				if len(s) == 0 {
					return
				}

				x, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
				if err != nil {
					return
				}
				stream <- x
			}
		}
	}

	go func() {
		defer close(stream)
		src()
	}()

	return stream
}

func displayMatrix(w io.Writer, m *metric.DistanceMatrix) error {
	var buf bytes.Buffer
	tbw := tabwriter.NewWriter(w, 0, 8, 0, '\t', 0)

	// distance matricies are nxn symmetric
	n, _ := m.Dims()

	// Write column headers
	for j := 0; j < n; j++ {
		buf.WriteRune('\t')
		buf.WriteString(strconv.Itoa(j))
	}
	buf.WriteRune('\n')
	buf.WriteTo(tbw)
	buf.Reset()

	// Write out rows
	for i := 0; i < n; i++ {
		buf.WriteString(strconv.Itoa(i))
		buf.WriteRune('\t')
		for j := 0; j < n; j++ {
			x := m.At(i, j)
			buf.WriteString(strconv.FormatFloat(x, 'f', 4, 64))
			if j != n-1 {
				buf.WriteRune('\t')
			}
		}
		buf.WriteRune('\n')
		buf.WriteTo(tbw)
		buf.Reset()
	}

	return tbw.Flush()
}
