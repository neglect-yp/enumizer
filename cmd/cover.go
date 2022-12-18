package cmd

import (
	"errors"
	"github.com/neglect-yp/enumizer/analyzer/enumcover"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/analysis/singlechecker"
	"os"
)

var (
	coverExample = `  enumizer cover ./...`

	coverCmd = &cobra.Command{
		Use:     "cover PATH",
		Short:   "Check that switch statements cover all variants of passed enum",
		Example: coverExample,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("must specify PATH")
			}

			// HACK: remove subcommand for singlechecker
			os.Args = append(os.Args[:1], args...)

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			singlechecker.Main(enumcover.Analyzer)
		},
	}
)

func init() {
	rootCmd.AddCommand(coverCmd)
}
