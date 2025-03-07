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

	"github.com/sfborg/sf/internal/io/text"
	"github.com/sfborg/sf/pkg/config"
	"github.com/spf13/cobra"
)

// fromTextCmd represents the fromText command
var fromTextCmd = &cobra.Command{
	Use:   "text text-file output-file",
	Short: "Converts a list of scientific names to SFGA format",
	Long: `This command imports a list of scientific names from a plain text
file and converts it into the Species File Group Archive (SFGA) format. Each
line in the input text file should contain a single scientific name. The
output of this command will be a new SFGA database containing the
supplied names and parsed information.

Names with authorship would perform better for further comparison with 
other checklists.
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
		txt := text.New(cfg)

		err := txt.Import(src, out)
		if err != nil {
			file := filepath.Base(src)
			slog.Error("Cannot import text file", "file", file, "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	fromCmd.AddCommand(fromTextCmd)
}
