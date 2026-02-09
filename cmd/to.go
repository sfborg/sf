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
	"github.com/spf13/cobra"
)

// toCmd represents the to command
var toCmd = &cobra.Command{
	Use:   "to",
	Short: "Converts SFGA file to a variety of formats",
	Long: `Converts a Species File Group Archive (SFGA) file into a variety of
output formats.

This command offers several subcommands to facilitate the conversion from
SFGA to various targets:

coldp: Exports data to the Catalogue of Life Data Package (CoLDP) format,
	used for taxonomic data exchange.

dwca: Exports data to the Darwin Core Archive (DwCA) format, a standard
	format for sharing biodiversity data.

text: Exports a simple text file with one scientific name per line.

xsv: Exports data to Comma-Separated Value (CSV).

The source SFGA file can be a local file path or a remote URL.

    sf to coldp input.sfga /path_to/output_coldp/output.zip
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(toCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
