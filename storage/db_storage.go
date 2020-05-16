package storage

import (
	sqlModule "database/sql"
	"fmt"
	rncConfigModule "github.com/chernogorsky/rnc/config"

	//"fmt"
	//rncConfig "github.com/chernogorsky/rnc/config"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type sqlRows = sqlModule.Rows

type sqlDBInt interface {
	Ping() error
	Close() error
	Query(query string, args ...interface{}) (*sqlRows, error)
}



type sqlInt interface {
	Open(driverName string, dataSourceName string) (sqlDBInt, error)
}

type sqlIntImpl struct {}

var sql sqlInt = sqlIntImpl{}

var sqlModuleOpen = sqlModule.Open
func (sqlIntImpl) Open(driverName string, dataSourceName string) (sqlDBInt, error) {
	return sqlModuleOpen(driverName, dataSourceName)
}




type rncConfigInt interface {
	GetRemoteConfig(string) (string, error)
}

type rncConfigIntImpl struct {}

var rncConfig rncConfigInt = rncConfigIntImpl{}

var rncConfigModuleGetRemoteConfig = rncConfigModule.GetRemoteConfig
func (rncConfigIntImpl) GetRemoteConfig(s string) (string, error) {
	return rncConfigModuleGetRemoteConfig(s)
}


type SDB struct {
	sqlDBInt
	storageType string
	dbHost string
	dbName string

}


func GetSqlStorage() (*SDB, error) {
	dbHost,_ := rncConfig.GetRemoteConfig("DB_HOST")
	dbName,_ := rncConfig.GetRemoteConfig("DB_NAME")
	dbUser,_ := rncConfig.GetRemoteConfig("DB_USER")
	dbPwd,_ := rncConfig.GetRemoteConfig("DB_PWD")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", dbUser, dbPwd, dbHost, dbName)

	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Error("DB. Failed to create a connection to" + dbHost)
		return nil,err
	}
	return &SDB{db, "mysql", dbHost, dbName}, nil
}

func (db *SDB) OpenStorage() error {

	log.Info("Connecting to " + db.dbHost)
	err := db.Ping()
	if err != nil {
		log.Error("DB Ping error for host " + db.dbHost)
		return err
	}

	return nil
}

func (db *SDB) Close() error {
	return db.sqlDBInt.Close()
}

func (db *SDB) GetDevices() ([] Device, error){
	var result [] Device

	log.Info("GetDevices. query DB")
	rows, err := db.Query("select deviceId, name from devices")
	if err != nil {
		log.Error(fmt.Sprintf("GetDevices. Error during the query %s", err.Error()))
		return nil, err
	}
	defer rows.Close()
	log.Info("GetDevices. Scan rows")
	for rows.Next() {
		rawDev := Device{}
		err := rows.Scan(&rawDev.Id, &rawDev.Name)
		if err != nil {
			log.Error(fmt.Sprintf("GetDevices. Error during scanning the rows. %s", err.Error()))
			return nil, err
		}
		result = append(result, rawDev)

	}
	err = rows.Err()
	if err != nil {
		log.Error(fmt.Sprintf("GetDevices. Error during the last rows scan %s", err.Error()))
		return nil, err
	}

	log.Info(fmt.Sprintf("GetDevices. result len: %d", len(result)))
	return result, nil
}