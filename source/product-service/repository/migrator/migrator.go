package migrator

import (
	"clean-code-structure/repository/pgsql"
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   pgsql.Config
	migrations *migrate.FileMigrationSource
	db         *sql.DB
}

func New(dbConfig pgsql.Config) Migrator {
	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/pgsql/migrations",
	}
	db := pgsql.New(dbConfig)

	return Migrator{
		dbConfig:   dbConfig,
		dialect:    "postgres",
		migrations: migrations,
		db:         db.Conn(),
	}
}

func (m Migrator) Up() {
	n, err := migrate.Exec(m.db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	n, err := migrate.Exec(m.db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %v", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}

func (m Migrator) Refresh() {
	m.Down()
	m.Up()
}
