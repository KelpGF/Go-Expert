package cmd

import (
	"github.com/KelpGF/Go-Expert/015-Cobra-CLI/internal/database"
	"github.com/spf13/cobra"
)

func CreateCategoryDBAdapter(categoryDB *database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		category, err := categoryDB.Create(name, description)
		if err != nil {
			return err
		}

		cmd.Println("Category created with ID:", category.ID)
		return nil
	}
}

func ListCategoryDBAdapter(categoryDB *database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetString("id")

		if id != "" {
			category, err := categoryDB.FindByID(id)
			if err != nil {
				return err
			}

			printCategory(cmd, category)
		} else {
			categories, err := categoryDB.FindAll()

			if err != nil {
				return err
			}

			for _, category := range categories {
				printCategory(cmd, &category)
			}
		}
		return nil
	}
}

func printCategory(cmd *cobra.Command, category *database.Category) {
	cmd.Println()
	cmd.Println("Category ID:", category.ID)
	cmd.Println("Category Name:", category.Name)
	cmd.Println("Category Description:", category.Description)
	cmd.Println()
}
