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

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/from/fdwca"
	"github.com/spf13/cobra"
)

// fromDwcaCmd represents the fromDwca command
var fromDwcaCmd = &cobra.Command{
	Use:   "dwca <input-dwca.zip> [output-sfga] [options]",
	Short: "Converts DwCA format to SFGA format",
	Long: `This command imports data from a Darwin Core Archive (DwCA) file and
converts it into the Species File Group Archive (SFGA) format. DwCA is a
common format for sharing biodiversity data, and this command allows you
to integrate such data into the SFGA ecosystem.

The command requires the path to the input DwCA file (which should be a
ZIP archive) and the desired path for the output SFGA file.

Example:

  sf from dwca my_dwca_data.zip /path/to/output-sfga

Input can be either a local file or a URL to a remote DwCA file.
If an output path is not provided, it will be generated in the the current
directory.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		src := args[0]
		out := args[1]

		flags := []flagFunc{jobsFlag, zipFlag}
		// append opts using flags input
		for _, v := range flags {
			v(cmd)
		}
		cfg := config.New(opts...)
		dwca := fdwca.New(cfg)

		err := dwca.Import(src, out)
		if err != nil {
			file := filepath.Base(src)
			slog.Error("Cannot import DwCA file", "file", file, "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	fromCmd.AddCommand(fromDwcaCmd)
}
