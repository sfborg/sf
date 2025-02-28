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
	"os"

	"github.com/sfborg/sf/internal/ent/io/diffio"
	"github.com/sfborg/sf/pkg/config"
	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff sfga1 sfga2",
	Short: "Compares data from two SFGA files",
	Long: `Compares data from two SFGA files. It is possible to do comparison
between specific taxon in the files, providing either name or taxon.id,
for example:

  sf diff sfga1.sqlite.zip sfga2.sqlite.zip --taxon1 Plantae --taxon2 Plantae

  or

  sf diff sfga1.sqlite.zip sfga2.sqlite.zip --taxon1 2938 --taxon2 taxon-3343

Files can be local or remote. Remote files can be accessed via HTTP URL.
`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag(cmd)
		flags := []flagFunc{
			srcTaxonFlag, trgTaxonFlag,
		}
		for _, v := range flags {
			v(cmd)
		}

		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}

		src := args[0]
		dst := args[1]

		cfg := config.New(opts...)
		diff := diffio.New(cfg)
		diff.Compare(src, dst)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	diffCmd.Flags().StringP("source-taxon", "s", "",
		"source's highest taxon to compare",
	)

	diffCmd.Flags().StringP("target-taxon", "t", "",
		"target's highest taxon to compare",
	)
}
