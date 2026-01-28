package db

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entSQL "entgo.io/ent/dialect/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"api-go/ent"
)

var EntDB *ent.Client

func SetEntDB(urlDB string) {
	connectionDB, err := sql.Open("pgx", urlDB)
	if err != nil {
		panic("failed to open database connection")
	}

	err = connectionDB.Ping()
	if err != nil {
		panic("failed to ping database")
	}

	sqlDB := entSQL.OpenDB(dialect.Postgres, connectionDB)
	EntDB = ent.NewClient(ent.Driver(sqlDB))
}
