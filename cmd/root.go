package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "enumizer",
		Short: "enumizer is a command-line tool to generate enum helpers and check enum coverage.",
		// Example:      strings.Join([]string{exportExample, applyExample, createExample, diffExample}, "\n"),
		// SilenceUsage: true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
