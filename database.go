package main

import (
	"go-api-template/ent"
	"go-api-template/settings"
	//_ "github.com/mattn/go-sqlite3" UNCOMMENT THIS LINE IF YOU WANT TO USE SQLITE3
	//_ "github.com/lib/pq" UNCOMMENT THIS LINE IF YOU WANT TO USE POSTGRESQL
	//_ "github.com/go-sql-driver/mysql" UNCOMMENT THIS LINE IF YOU WANT TO USE MYSQL or MARIADB
)

func NewDatabase(databaseSettings *settings.DatabaseSettings) *ent.Client {
	client, err := ent.Open(databaseSettings.Driver, databaseSettings.ConnectionString)
	if err != nil {
		panic(err)
	}

	return client
}
