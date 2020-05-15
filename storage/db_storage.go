package storage

import (
	"database/sql"
	"fmt"
	rncConfig "github.com/chernogorsky/rnc/config"

	//"fmt"
	//rncConfig "github.com/chernogorsky/rnc/config"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type StorageDB struct {
	*sql.DB
	storageType string
	dbHost string
	dbName string

}

func GetSqlStorage() (*StorageDB, error) {
	dbHost,_ := rncConfig.GetRemoteConfig("DB_HOST")
	dbName,_ := rncConfig.GetRemoteConfig("DB_NAME")
	dbUser,_ := rncConfig.GetRemoteConfig("DB_USER")
	dbPwd,_ := rncConfig.GetRemoteConfig("DB_PWD")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", dbUser, dbPwd, dbHost, dbName)

	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		return nil,err
	}
	return &StorageDB{db, "mysql", dbHost, dbName}, nil
}

func (db *StorageDB) OpenStorage() (*StorageDB, error) {

	log.Info("Connecting to " + db.dbHost)
	err := db.Ping()
	if err != nil {
		log.Error("DB Ping error for host " + db.dbHost)
		return nil,err
	}

	return db, nil
}
