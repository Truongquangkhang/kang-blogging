package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"strconv"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	cmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			action := args[1]

			dbHost := os.Getenv("DB_HOST")
			dbPort := os.Getenv("DB_PORT")
			dbUser := os.Getenv("DB_USER")
			dbPass := os.Getenv("DB_PASS")
			dbName := os.Getenv("DB_NAME")
			service := os.Getenv("SERVICE")

			sourceUrl := fmt.Sprintf("file://%s/migrations", service)
			databaseUrl := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

			m, err := migrate.New(sourceUrl, databaseUrl)
			if err != nil {
				log.Fatal(err)
			}

			switch action {
			case "step-up":
				steps, err := strconv.Atoi(args[2])
				if err != nil {
					panic(err)
				}
				err = m.Steps(steps)
			case "step-down":
				steps, err := strconv.Atoi(args[2])
				if err != nil {
					panic(err)
				}
				err = m.Steps(-steps)
			case "up":
				err = m.Up()
			case "down":
				err = m.Down()
			default:
				err = m.Up()
			}

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return cmd.Execute()
}
