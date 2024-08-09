package cmd

import (
	"github.com/KelpGF/Go-Expert/015-Cobra-CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCategoryCmd(categoryDb *database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category.",
		Long:  `Create a new category with a name and description.`,
		RunE:  CreateCategoryDBAdapter(categoryDb),
	}
}

func init() {
	createCmd := newCreateCategoryCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
