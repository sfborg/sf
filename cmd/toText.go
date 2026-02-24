/*
Copyright © 2026 Dmitry Mozzherin <dmozzherin@gmail.com>

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
	"github.com/spf13/cobra"
	"log/slog"
	"os"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/to/ttext"
)

// toTextCmd represents the toText command
var toTextCmd = &cobra.Command{
	Use:   "text",
	Short: "Converts SFGA file to list of names.",
	Long: `Converts a Species File Group Archive (SFGA) file into a plain text
file containing one scientific name per line.

The command requires two arguments: the path to the input SFGA file (local
path or URL) and the desired path for the output text file.

Examples:
    sf to text input.sfga.sql names.txt
    sf to text input.sfga.sqlite.zip names.txt
    sf to text https://example.com/data.sfga.sql names.txt
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
			_ = cmd.Help()
			os.Exit(0)
		}
		sfgaPath := args[0]
		xsvPath := args[1]

		slog.Info("Extracting SFGA data", "path", sfgaPath)

		txt := ttext.New(cfg)
		err := txt.Export(sfgaPath, xsvPath)
		if err != nil {
			slog.Error("Cannot export to text", "error", err)
			os.Exit(1)
		}

	},
}

func init() {
	toCmd.AddCommand(toTextCmd)
}
