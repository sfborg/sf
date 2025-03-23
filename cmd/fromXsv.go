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

	"github.com/sfborg/sf/internal/io/fxsv"
	"github.com/sfborg/sf/pkg/config"
	"github.com/spf13/cobra"
)

// fromXsvCmd represents the fromXsv command
var fromXsvCmd = &cobra.Command{
	Use:   "xsv src-file-or-url output-file",
	Short: "Converts CSV/TSV/PSV fiels to SFGA format",
	Long: `Imports data from Comma-Separated Value (CSV), Tab-Separated
Value (TSV), or Pipe-Separated Value (PSV) files and converts it into
the Species File Group Archive (SFGA) format.

This command is designed for transforming tabular data, commonly used in
biological and taxonomic datasets, into a structured SFGA database.

INPUT REQUIREMENTS:

Header Row: The XSV file *must* include a header row as the first line.
  This row defines the column names for the data.

Column Names: Column names should correspond to either Darwin Core or
  CoLDP (Catalogue of Life Data Package) terms. This ensures the data is
  correctly mapped into the SFGA structure.

ScientificName Field: A column named "ScientificName" is *mandatory*.
  This field contains the scientific names of taxa. Other columns, while
  optional, allow for a more comprehensive dataset.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		src := args[0]
		out := args[1]

		flags := []flagFunc{
			zipFlag, detailsFlag,
		}
		// append opts using flags input
		for _, v := range flags {
			v(cmd)
		}
		cfg := config.New(opts...)
		xsv := fxsv.New(cfg)

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
