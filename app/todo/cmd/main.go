package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/DragonPow/todo_project/app/todo/config"
	"github.com/DragonPow/todo_project/app/todo/internal/service"
	"github.com/DragonPow/todo_project/library/log"
	"github.com/DragonPow/todo_project/library/server"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v3"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Printf("Error when run service: %s\n", err)
	}
}

func run(args []string) error {
	cmd := cli.Command{
		Commands: []*cli.Command{
			{
				Name:   "server",
				Usage:  "Start service",
				Action: startServer,
			},
			{
				Name:   "migrate-up",
				Usage:  "Migrate database",
				Action: migrateDb,
			},
		},
	}

	if err := cmd.Run(context.Background(), args); err != nil {
		return err
	}
	return nil
}

func startServer(_ context.Context, cmd *cli.Command) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logger, err := cfg.BuildLogger()
	if err != nil {
		return err
	}

	serviceInstance, err := service.NewService(cfg, logger)
	if err != nil {
		return err
	}
	logger.Info("Service created")

	srv, err := server.NewServer(
		server.WithService(serviceInstance),
		server.WithGrpcListener(cfg.Server.GrpcConfig),
		server.WithHttpListener(cfg.Server.HttpConfig),
		server.GrpcUnaryInterceptors(
			log.UnaryServerInterceptor(logger),
			log.UnaryServerPanicInterceptor(logger),
		),
	)
	if err != nil {
		return err
	}

	return srv.Start()
}

func migrateDb(_c context.Context, cmd *cli.Command) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	fmt.Println("Migrate database")

	db, err := sql.Open("postgres", cfg.Database.DSN())
	if err != nil {
		fmt.Println("Error when open database", err)
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println("Error when create driver", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		cfg.MigrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println("Error when create migrate instance", err)
		return err
	}
	return m.Up()
}
