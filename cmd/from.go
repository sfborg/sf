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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Process all persistent flags for all child commands
		flags := []flagFunc{
			zipFlag, detailsFlag, codeFlag, jobsFlag,
			withQuotesFlag, addParentsFlag,
		}
		for _, v := range flags {
			v(cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(fromCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	fromCmd.PersistentFlags().BoolP("zip-output", "z", false, "compress output with zip")
	fromCmd.PersistentFlags().BoolP(
		"no-parse-details", "n", false,
		"do not use name parsing to populate SFGA name fields",
	)
	fromCmd.PersistentFlags().IntP("jobs-number", "j", 0, "number of concurrent jobs")
	fromCmd.PersistentFlags().
		BoolP("quotes-allowed", "q", false, "fields in pipe or tsv file might be escaped by quotes")

	nomCodeHelp := `Sets the nomenclatural code for entries where code is not
provided and changes name parsing rules accordingly

Accepted values are:
  - 'bact', 'icnp', 'bacterial' for bacterial code
  - 'bot', 'icn', 'botanical' for botanical code
  - 'cult', 'icncp', 'cultivar' for cultivar code
  - 'vir', 'virus', 'viral', 'ictv', 'icvcn' for viral code
  - 'zoo', 'iczn', 'zoological' for zoological code
`

	fromCmd.PersistentFlags().StringP("code-of-nomenclature", "c", "",
		nomCodeHelp)

	fromCmd.PersistentFlags().BoolP("add-parents", "p", false,
		"convert flat hierarchy to parent/child tree")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fromCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
