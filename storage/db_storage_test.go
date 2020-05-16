package storage

import (
	"errors"
	"github.com/stretchr/testify/assert"
	sqlM "database/sql"
	"testing"
)


var sqlDBIntMockPingError error
var sqlDBIntMockCloseError error
type sqlDBIntMock struct {}
func (sqlDBIntMock) Ping() error {
	return sqlDBIntMockPingError
}
func (sqlDBIntMock) Close() error {
	return sqlDBIntMockCloseError
}


type rncConfigIntMock struct {}
func (rncConfigIntMock) GetRemoteConfig(s string) (string, error) {
	return s, nil
}

var sqlIntMockOpenErr error

type sqlIntMock struct {}
func (sqlIntMock) Open(driverName string, dataSourceName string) (sqlDBInt, error) {
	return nil, sqlIntMockOpenErr
}


func TestGetSqlStorage(t *testing.T) {
	rncConfig = rncConfigIntMock{}
	sql = sqlIntMock{}
	sqlIntMockOpenErr = nil

	db, err := GetSqlStorage()
	assert.Nil(t, err)
	assert.Equal(t, "DB_HOST", db.dbHost)
	assert.Equal(t, "DB_NAME", db.dbName)

	sqlIntMockOpenErr = errors.New("some error")
	db, err = GetSqlStorage()
	assert.NotNil(t, err)

}
func GetRemoteConfigMock(s string) (string, error){
	return s, nil
}
func TestGetRemoteConfigInterface(t *testing.T) {
	rncConfigModuleGetRemoteConfig = GetRemoteConfigMock
	rncConfig = rncConfigIntImpl{}
	res, err := rncConfig.GetRemoteConfig("string")
	assert.Equal(t, "string", res)
	assert.Nil(t, err)

}

func TestGetSqlStorageOpen(t *testing.T) {
	rncConfig = rncConfigIntMock{}
	sql = sqlIntMock{}
	sqlIntMockOpenErr = nil

	db, err := GetSqlStorage()
	assert.Nil(t, err)
	assert.Equal(t, "DB_HOST", db.dbHost)
	assert.Equal(t, "DB_NAME", db.dbName)

	sqlIntMockOpenErr = errors.New("some error")
	db, err = GetSqlStorage()
	assert.NotNil(t, err)

}


func TestOpenStorage(t *testing.T) {
	sdb := SDB {
		sqlDBIntMock {},
		"mock",
		"dbHost",
		"dbName",
	}
	sqlDBIntMockPingError = nil
	err := sdb.OpenStorage()
	assert.Nil(t, err)

	sqlDBIntMockPingError = errors.New("some error")
	err = sdb.OpenStorage()
	assert.NotNil(t, err)
}

func sqlModuleOpenMock(driverName string, dataSourceName string) (*sqlM.DB, error){
	return nil, nil
}
func TestOpenInterface(t *testing.T) {
	sqlModuleOpen = sqlModuleOpenMock
	sql = sqlIntImpl{}
	_, err := sql.Open("","")
	assert.Nil(t, err)

}



func TestClose(t *testing.T) {
	sdb := SDB {
		sqlDBIntMock {},
		"mock",
		"dbHost",
		"dbName",
	}
	sqlDBIntMockCloseError = errors.New("some error")
	err := sdb.Close()
	assert.NotNil(t, err)
}