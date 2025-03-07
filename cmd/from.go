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

// fromCmd represents the from command
var fromCmd = &cobra.Command{
	Use:   "from",
	Short: "Converts a variety of formats to SFGA format",
	Long: `Converts a variety of data formats into the Species File
Group Archive (SFGA) format.

This command offers several subcommands to facilitate the conversion from
various sources:

text: Imports a simple text file, where each line contains a scientific name.

xsv: Imports data from Comma-Separated Value (CSV), Tab-Separated
	Value (TSV), or Pipe-Separated Value (PSV) files. These files must contain
	a header row with column names that correspond to Darwin Core or CoLDP terms.
	At a minimum, the "ScientificName" field is required.

coldp: Imports data structured in the Catalogue of Life Data Package
	(CoLDP) format.

dwca: Imports data from Darwin Core Archive (DwCA) files, a common
	format for sharing biodiversity data.

tw: Imports data from a TaxonWorks project.

The resulting SFGA files can be used for further analysis, comparison with
other datasets, or conversion to other formats supported by this program.

    sf from xsv mydata.csv output.sfga
`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	_ = cmd.Help()
	// },
}

func init() {
	rootCmd.AddCommand(fromCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	fromCmd.PersistentFlags().BoolP("zip-output", "z", false, "compress output with zip")
	fromCmd.PersistentFlags().BoolP("parse-details", "p", false, "use detailed parsing of names")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fromCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
