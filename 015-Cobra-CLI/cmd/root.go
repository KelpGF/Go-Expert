package cmd

import (
	"database/sql"
	"os"

	"github.com/KelpGF/Go-Expert/015-Cobra-CLI/internal/database"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS categories (id TEXT PRIMARY KEY, name TEXT, description TEXT)`)
	return db
}

func GetCategoryDB(db *sql.DB) *database.Category {
	return database.NewCategory(db)
}

var rootCmd = &cobra.Command{
	Use:   "015-Cobra-CLI",
	Short: "015-Cobra-CLI is a CLI application",
	Long:  `015-Cobra-CLI is a CLI application that demonstrates how to use Cobra.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
