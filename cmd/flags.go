package cmd

import (
	"fmt"
	"log/slog"
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

func jobsFlag(cmd *cobra.Command) {
	i, _ := cmd.Flags().GetInt("jobs-number")
	if i > 0 {
		if i > 100 {
			slog.Error("Jobs number should be between 1 and 100")
			slog.Info("Setting jobs number to 100")
			i = 100
		}
		opts = append(opts, config.OptJobsNum(i))
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

func zipFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("zip-output")
	if b {
		opts = append(opts, config.OptWithZipOutput(true))
	}
}

func detailsFlag(cmd *cobra.Command) {
	b, _ := cmd.Flags().GetBool("parse-details")
	if b {
		opts = append(opts, config.OptWithDetails(true))
	}

}
