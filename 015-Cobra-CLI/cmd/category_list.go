package cmd

import (
	"github.com/KelpGF/Go-Expert/015-Cobra-CLI/internal/database"
	"github.com/spf13/cobra"
)

func newListCategoryCmd(categoryDb *database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List categories.",
		Long:  `List all categories or filter by id.`,
		RunE:  ListCategoryDBAdapter(categoryDb),
	}
}

func init() {
	createCmd := newListCategoryCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("id", "i", "", "ID of the category")
}
