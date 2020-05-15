package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func (db *sql.DB) Open() (*sql.DB, error) {
	dbHost,_ := rncConfig.GetRemoteConfig("DB_HOST")
	dbName,_ := rncConfig.GetRemoteConfig("DB_NAME")
	dbUser,_ := rncConfig.GetRemoteConfig("DB_USER")
	dbPwd,_ := rncConfig.GetRemoteConfig("DB_PWD")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", dbUser, dbPwd, dbHost, dbName)


	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		return nil,err
	}

	log.Info("Connecting to " + dbHost)
	err = db.Ping()
	if err != nil {
		return nil,err
	}

	return db, nil
}
