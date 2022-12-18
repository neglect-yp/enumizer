package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

var (
	rootCmd = &cobra.Command{
		Use:     "enumizer",
		Short:   "enumizer is a command-line tool to generate enum helpers and check enum coverage.",
		Example: strings.Join([]string{generateExample, coverExample}, "\n\n"),
	}
)

func Execute() error {
	return rootCmd.Execute()
}
