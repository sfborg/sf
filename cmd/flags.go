package cmd

import (
	"fmt"
	"os"

	sf "github.com/sfborg/sf/pkg"
	"github.com/spf13/cobra"
)

type flagFunc func(cmd *cobra.Command)

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
