package cmd

import (
	"fmt"
	"os"

	"github.com/sfborg/sf/config"
	"github.com/sfborg/sf/pkg/sf"
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

func withQuotesFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("with-quotes")
	if b {
		opts = append(opts, config.OptWithQuotes(true))
	}
}

func zipFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("zip-output")
	if b {
		opts = append(opts, config.OptWithZipOutput(true))
	}
}

func coldpNameUsageFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("name-usage")
	if b {
		opts = append(opts, config.OptColdpNameUsage(true))
	}
}

func jobsFlag(cmd *cobra.Command) {
	i, _ := cmd.Flags().GetInt("jobs-number")
	if i > 0 {
		opts = append(opts, config.OptJobsNum(i))
	}
}

func detailsFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("no-parse-details")
	if b {
		opts = append(opts, config.OptNoParser(true))
	}
}
