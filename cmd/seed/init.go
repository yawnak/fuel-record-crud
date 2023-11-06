package main

import (
	dbsql "database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/yawnak/fuel-record-crud/ent"
	"github.com/yawnak/fuel-record-crud/internal/config"
)

func initRepo() *ent.Client {
	conf := Must(config.NewDBConnConfig(configPath))
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", conf.User, config.MustGetEnv("POSTGRES_PASSWORD"), conf.Host, conf.Port, conf.DbName)
	db, err := dbsql.Open("pgx", databaseURL)
	if err != nil {
		log.Panicln("error opening db:", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
