/*
Copyright © 2025 Dmitry Mozzherin <dmozzherin@gmail.com>

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
	"path/filepath"

	"github.com/sfborg/sf/internal/io/xsv"
	"github.com/sfborg/sf/pkg/config"
	"github.com/spf13/cobra"
)

// fromXsvCmd represents the fromXsv command
var fromXsvCmd = &cobra.Command{
	Use:   "xsv src-file-or-url output",
	Short: "Converts CSV/TSV/PSV fiels to SFGA format",
	Long: `Takes a Comma/Tab/Pipe Separated Value (XSV) file and converts it to
a Species File Group Archive file. The file must have headers where at least
some DarwinCore or CoLDP terms are provided. The "ScientificName" field is
required. The resuting SFGA file can later be used to compare it with other
SFGA files, or be converted to other common formats.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		src := args[0]
		out := args[1]

		cfg := config.New()
		xsv := xsv.New(cfg)

		err := xsv.Import(src, out)
		if err != nil {
			file := filepath.Base(src)
			slog.Error("Cannot import XSV file", "file", file, "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	fromCmd.AddCommand(fromXsvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fromXsvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	fromXsvCmd.Flags().StringP("delimiter", "d", "",
		`delimiter used in the file:
  auto, csv, tsv, psv are supported
  auto is default.`)
}
