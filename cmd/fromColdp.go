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

	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/from/fcoldp"
	"github.com/spf13/cobra"
)

// fromColdpCmd represents the fromColdp command
var fromColdpCmd = &cobra.Command{
	Use:   "coldp",
	Short: "Converts CoLDP format to SFGA format",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		src := args[0]
		out := args[1]

		flags := []flagFunc{}
		// append opts using flags input
		for _, v := range flags {
			v(cmd)
		}
		cfg := config.New(opts...)
		fc := fcoldp.New(cfg)

		err := fc.Import(src, out)
		if err != nil {
			file := filepath.Base(src)
			slog.Error("Cannot import CoLDP file", "file", file, "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	fromCmd.AddCommand(fromColdpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fromColdpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fromColdpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
