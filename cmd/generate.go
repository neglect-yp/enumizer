package cmd

import (
	"errors"
	"github.com/neglect-yp/enumizer/internal/enumizer"
	"github.com/spf13/cobra"
)

var (
	generateExample = `* Generate enum helpers
  enumizer generate ./...`

	generateCmd = &cobra.Command{
		Use:     "generate PATH",
		Short:   "Generate enum helpers that include stringer and validation methods",
		Example: generateExample,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("must specify PATH")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}

			return enumizer.Generate(cmd.Context(), args[0], output)
		},
	}
)

func init() {
	generateCmd.Flags().String("output", "enumizer.gen.go", "output filename")

	rootCmd.AddCommand(generateCmd)
}
