package main

import (
	"github.com/yawnak/fuel-record-crud/internal/adapters/repoadapt"
	"github.com/yawnak/fuel-record-crud/internal/config"
)

func initRepo() *repoadapt.Client {
	conf := Must(config.NewDBConnConfig(configPath))
	return Must(repoadapt.NewClientPSQL(
		conf.User,
		config.MustGetEnv("POSTGRES_PASSWORD"),
		conf.Host,
		conf.Port,
		conf.DbName,
	))
}
