package main

import (
	"fmt"
	"go-rest-ddd/pkg/config"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CT version v1.0.0.0")
	},
}

func commandMigrationStatus(cfg *config.Migration) *cobra.Command {
	c := &cobra.Command{
		Use:   "migration:status",
		Short: "Chek migration status",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("status")
		},
	}
	return c
}

func commandMigrationCreate(cfg *config.Migration) *cobra.Command {
	var table, create string
	c := &cobra.Command{
		Use:   "migration:create",
		Short: "Create SQL file migration",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("arg 0 is required")
				os.Exit(1)
			}
			if len(table) < 1 && len(create) < 1 {
				fmt.Println("create or table param required")
				os.Exit(1)
			}
			param := args[0]
			version := int(time.Now().Unix())
			filename := strconv.Itoa(version) + "_" + param
			f1, _ := os.Create(cfg.Path + "/" + filename + ".up.sql")
			f2, _ := os.Create(cfg.Path + "/" + filename + ".down.sql")
			if len(table) > 0 {
				f1.WriteString(fmt.Sprintf("ALTER TABLE %s\nALTER COLUMN %s TYPE %s;", table, "column_name", "column_definition"))
				f2.WriteString(fmt.Sprintf("ALTER TABLE %s\nALTER COLUMN %s TYPE %s;", table, "column_name", "column_definition"))
			} else {
				f1.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s \n(\n);", create))
				f2.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s;", create))
			}
		},
	}
	c.Flags().StringVar(&table, "table", "", "table name of migration")
	c.Flags().StringVar(&create, "create", "table_name", "table name of migration")
	return c
}

func commandMigrationUp(cfg *config.Migration) *cobra.Command {
	c := &cobra.Command{
		Use:   "migration:up",
		Short: "Migrate up database",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New(fmt.Sprintf("file://%s", cfg.Path), cfg.Database)
			if err != nil {
				log.Fatal(err)
			}
			if err := m.Up(); err != nil {
				log.Fatal(err)
			}
		},
	}
	return c
}

func commandMigrationDown(cfg *config.Migration) *cobra.Command {
	var input int
	c := &cobra.Command{
		Use:   "migration:down",
		Short: "Migrate down database",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New(fmt.Sprintf("file://%s", cfg.Path), cfg.Database)
			if err != nil {
				log.Fatal(err)
			}
			err = m.Down()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	c.Flags().IntVar(&input, "step", 1, "Number of step down migration")
	return c
}

func init() {
	cfg, err := config.Migrations()
	if err != nil {
		panic(err)
	}

	command.AddCommand(commandMigrationStatus(cfg))
	command.AddCommand(commandMigrationCreate(cfg))
	command.AddCommand(commandMigrationUp(cfg))
	command.AddCommand(commandMigrationDown(cfg))
}

func main() {
	err := command.Execute()
	if err != nil {
		log.Println(err.Error())
	}
}
