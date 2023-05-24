package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
)

func (d *DataBase) migrateDB() error {
	fmt.Println("Migrating our DB")
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Could not create the postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("Could not run up migrations: %w", err)
	}
	fmt.Println("Succesfully migrated to the database!")
	return nil

}
