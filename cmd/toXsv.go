/*
copyright © 2026 Dmitry Mozzherin <dmozzherin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/to/txsv"
	"github.com/spf13/cobra"
)

// toXsvCmd represents the toXsv command
var toXsvCmd = &cobra.Command{
	Use:   "xsv",
	Short: "Converts SFGA file to CSV format",
	Long: `Converts a Species File Group Archive (SFGA) file into a
Comma-Separated Value (CSV) file.

The command requires two arguments: the path to the input SFGA file (local
path or URL) and the desired path for the output CSV file.

Use the --zip-output flag to compress the output file.

Examples:
    sf to xsv input.sfga.sql output.csv
    sf to xsv input.sfga.sqlite.zip output.csv.zip --zip-output
    sf to xsv https://example.com/data.sfga.sql output.csv
`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := []flagFunc{
			zipFlag,
		}
		for _, v := range flags {
			v(cmd)
		}

		cfg := config.New(opts...)

		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		sfgaPath := args[0]
		xsvPath := args[1]

		slog.Info("Extracting SFGA data", "path", sfgaPath)

		xsv := txsv.New(cfg)
		err := xsv.Export(sfgaPath, xsvPath)
		if err != nil {
			slog.Error("Cannot export CSV", "error", err)
			os.Exit(1)
		}

	},
}

func init() {
	toCmd.AddCommand(toXsvCmd)

	toXsvCmd.Flags().BoolP("zip-output", "z", false, "compress output with zip")
}
