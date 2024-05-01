package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "migrate database",
				Subcommands: []*cli.Command{
					{
						Name:  "up",
						Usage: "Migration up",
						Action: func(context *cli.Context) error {
							return MigrateUp()
						},
					},
				},
			},
			{
				Name:  "gen-model",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					return GenModel()
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func MigrateUp() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	service := os.Getenv("SERVICE")

	sourceUrl := fmt.Sprintf("file://%s/migration", service)
	databaseUrl := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	m, err := migrate.New(sourceUrl, databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return m.Up()
}

func GenModel() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)
	g := gen.NewGenerator(gen.Config{
		OutPath: "../internal/common/model",
		Mode:    gen.WithoutContext,
	})

	gormdb, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(gormdb) // reuse your gorm db

	//// Generate basic type-safe DAO API for struct `model.User` following conventions
	//
	g.ApplyBasic(
		g.GenerateModelAs("users", "User"),
	)
	//g.ApplyBasic(
	//	g.GenerateAllTable(gen.ModelOpt(gen.Onl))...,
	//)
	// Generate the code
	g.Execute()
	return nil
}
