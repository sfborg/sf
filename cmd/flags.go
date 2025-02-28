package cmd

import (
	"fmt"
	"os"

	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/spf13/cobra"
)

type flagFunc func(cmd *cobra.Command)

func srcTaxonFlag(cmd *cobra.Command) {
	taxon, _ := cmd.Flags().GetString("source-taxon")
	if taxon != "" {
		opts = append(opts, config.OptCacheDir(taxon))
	}
}

func trgTaxonFlag(cmd *cobra.Command) {
	taxon, _ := cmd.Flags().GetString("target-taxon")
	if taxon != "" {
		opts = append(opts, config.OptCacheDir(taxon))
	}
}

func versionFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("version")
	if b {
		version := sf.GetVersion()
		fmt.Printf(
			"\nVersion: %s\nBuild:   %s\n",
			version.Version,
			version.Build,
		)
		os.Exit(0)
	}
}
