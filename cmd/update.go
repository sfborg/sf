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
	"github.com/sfborg/sf/pkg/from/fsfga"
	"github.com/spf13/cobra"
)

// updateCmd represents the "update" command which updates a backward-compatible
// SFGA file to the latest SFGA version or silently keeps already latest one.
// Optionally it converts flat clssification to parent/child nested hierarchy.
// This command reads data from an input SFGA file and saves it to an output
// SFGA file with the latest schema.
//
// Usage:
//
//	sf update <input-sfga> [output-sfga] [options]
//
// Description:
// The command processes an input SFGA file, which can be provided as a file path
// or via standard input (pipe). If the output file path is not specified, the
// updated file will be saved in the current directory with the same name as the
// original file. This command updates SFGA schema to the latest available.
//
// Examples:
//
//	# Update a single SFGA file and save to a specific location
//	sf update paleodb.sqlite.zip ~/tmp/paleodb -z
//
//	# Update multiple SFGA files using a pipe
//	ls /path/to/*.sqlite.zip | xargs -I {} sf update {}
//
// Arguments:
//
//	<input-sfga>   The path to the input SFGA file. This argument is required.
//	[output-sfga]  The optional path to save the updated SFGA file.
//
// Options:
//
//	Additional options can be provided to customize the update process.
//
// Note:
// This command requires at least one argument (input-sfga) to be provided.
var updateCmd = &cobra.Command{
	Use:   "update <input-sfga> [output-sfga] [options]",
	Short: "Updates backward compatible SFGA file to the latest SFGA version.",
	Long: `The command processes an input SFGA file, which can be provided as a file path
or via standard input (pipe). If the output file path is not specified, the
updated file will be saved in the current directory with the same name as the
original file. This command updates SFGA schema to the latest available if
needed.

Examples:

    # Update a single SFGA file and save to a specific location
    sf update paleodb.sqlite.zip ~/tmp/paleodb -z

    # Update multiple SFGA files using a pipe
    ls /path/to/*.sqlite.zip | xargs -I {} sf update {}

	  # Update SFGA and convert its flat hierarchy to a nested parent/child
	  # tree.
	  sf update flat.sqlite /tmp/tree.sqlite --add-parents

Arguments:

 <input-sfga>   The path to the input SFGA file. This argument is required.
 [output-sfga]  The optional path to save the updated SFGA file.

Options:

 Additional options can be provided to customize the update process.

Note:
This command requires at least one argument (input-sfga) to be provided.
	`,
	Args: cobra.MinimumNArgs(1), // Ensures at least one argument (input-sfga) is provided
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

		oldSfgaPath := args[0]
		sfgaPath := args[1]

		slog.Info("Updatting SFGA.", "path", sfgaPath)

		// making sfga for old data source
		fs := fsfga.New(cfg)
		err := fs.Import(oldSfgaPath, sfgaPath)
		if err != nil {
			slog.Error("Cannot update SFGA to latest schema.", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().BoolP("add-parents", "p", false,
		"create nested hierarchy from flat one",
	)
}
