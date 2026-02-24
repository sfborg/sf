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

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/to/tcoldp"
	"github.com/spf13/cobra"
)

// toColdpCmd represents the toColdp command
var toColdpCmd = &cobra.Command{
	Use:   "coldp <input-coldp.zip> [output-coldp.zip] [options]",
	Short: "Converts SFGA file to CoLDP format",
	Long: `Converts a Species File Group Archive (SFGA) file into the Catalogue
of Life Data Package (CoLDP) format.

The command requires two arguments: the path to the input SFGA file (local
path or URL) and the desired path for the output CoLDP zip file.

By default, the output contains separate Name, Taxon, and Synonym files.
Use the --name-usage flag to combine them into a single NameUsage file.

Examples:
    sf to coldp input.sfga.sql output.zip
    sf to coldp input.sfga.sqlite.zip output.zip --name-usage
    sf to coldp https://example.com/data.sfga.sql output.zip
`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := []flagFunc{
			coldpNameUsageFlag,
		}
		for _, v := range flags {
			v(cmd)
		}

		cfg := config.New(opts...)

		if len(args) != 2 {
			_ = cmd.Help()
			os.Exit(0)
		}

		sfgaPath := args[0]
		coldpPath := args[1]

		slog.Info("Extracting SFGA data", "path", sfgaPath)

		coldp := tcoldp.New(cfg)
		err := coldp.Export(sfgaPath, coldpPath)
		if err != nil {
			slog.Error("Cannot export CoLDP", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	toCmd.AddCommand(toColdpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toColdpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	toColdpCmd.Flags().BoolP(
		"name-usage", "u", false,
		"combine name, taxon, synonym to name_usage",
	)
}
