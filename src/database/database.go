package database

import (
	"database/sql"
	"fmt"
	"os"

	c "github.com/braulioinf/QRLauncher/src/constants"
	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

// Connect func
func Connect() (*sql.DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?multiStatements=true&parseTime=true",
		c.MYSQL_USER,
		c.MYSQL_PASSWORD,
		c.MYSQL_HOST,
		c.MYSQL_PORT,
		c.MYSQL_DATABASE,
	)

	instance, err := sql.Open(c.DRIVER, sourceName)
	if err != nil {
		return nil, err
	}

	if err := migration(instance); err != nil {
		return instance, err
	}

	return instance, nil
}

func migration(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/src/database/migrations", dir),
		c.DRIVER,
		driver,
	)

	if err != nil {
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("Running db migrations %v\n")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := migration.Version()
	if err != nil {
		return err
	}

	migration.Log.Printf("Active db version %d\n", version)

	return nil
}
