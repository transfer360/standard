package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/transfer360/standard/secret"
	"os"
	"strings"
	"time"
)

var ErrGettingDataCredentials = errors.New("unexpected error reading credentials from secret")
var ErrDatabaseConnection = errors.New("unexpected error connecting to the database")
var ErrDatabaseReading = errors.New("unexpected error reading from the database")
var ErrDatabaseInsert = errors.New("unexpected error saving to the database")
var ErrDatabaseDelete = errors.New("unexpected error deleting from the database")

type Configuration struct {
	Host      string `json:"host"`
	PrivateIP string `json:"private"`
	PublicIP  string `json:"public"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
}

// CredentialsFromSecretManagerPath  - Read data from Google Secrets Manager
func CredentialsFromSecretManagerPath(ctx context.Context, secretPath string) (dbc Configuration, err error) {

	data, err := secret.Get(ctx, secretPath)
	if err != nil {
		return dbc, err
	}

	err = json.Unmarshal(data, &dbc)

	if err != nil {
		log.Errorf("CredentialsFromSecretManagerPath:%v", err)
		return dbc, err
	}

	return dbc, nil
}

// GetCredentialsFromSecretEnvironmentVariable --------------------------------------------------------------
func GetCredentialsFromSecretEnvironmentVariable() (Configuration, error) {

	dbc := Configuration{}
	if len(os.Getenv("SECRET_PATH")) == 0 {
		return dbc, fmt.Errorf("missing SECRET_PATH Environment Variable")
	}

	err := json.Unmarshal([]byte(os.Getenv("SECRET_PATH")), &dbc)
	if err != nil {
		return dbc, fmt.Errorf("error parsing SECRET_PATH: %w", err)
	}

	return dbc, nil

}

// Connect - via CloudSQL
func Connect(dbc Configuration) (*sql.DB, error) {

	link, err := sql.Open("mysql", _mysqlPath(dbc))
	if err != nil {
		return nil, fmt.Errorf("connect:1: %w", err)
	}
	_sqlConnectionConfig(link)

	return link, nil

}

// ConnectProxy - via CloudSQL
func ConnectProxy(dbc Configuration) (*sql.DB, error) {

	sqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, "localhost", dbc.Database)
	link, err := sql.Open("mysql", sqlConnectionString)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			log.Warnf("Connection refused to Proxy")
			time.Sleep(1 * time.Second)
			return ConnectProxy(dbc)
		} else {
			return nil, fmt.Errorf("connect:1: %w", err)
		}
	}
	_sqlConnectionConfig(link)

	return link, nil

}

// ConnectIP via IP Address
func ConnectIP(dbc Configuration) (*sql.DB, error) {
	return _mysqlConnect(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, dbc.PublicIP, dbc.Database))
}

// ConnectPrivateIP via IP Address
func ConnectPrivateIP(dbc Configuration) (*sql.DB, error) {
	return _mysqlConnect(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, dbc.PrivateIP, dbc.Database))
}

func _mysqlPath(dbc Configuration) string {

	sqlPath := "/cloudsql"

	if len(os.Getenv("SQLPROXY")) > 0 {
		sqlPath = os.Getenv("SQLPROXY")
	}

	return fmt.Sprintf("%s:%s@unix(%s/%s)/%s?parseTime=true&timeout=5s", dbc.Username, dbc.Password, sqlPath, dbc.Host, dbc.Database)
}

// _mysqlConnect - Connect via SQL string using TCP/IP
func _mysqlConnect(dbString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbString)

	if err != nil {
		return nil, fmt.Errorf("ConnectIP:1: %w", err)
	}

	_, err = db.Exec("SET SESSION time_zone = 'europe/london'")

	if err != nil {
		return nil, fmt.Errorf("Connect:sql.Open:2: %v", err)
	}

	return db, nil
}

// _sqlConnectionConfig ----------------------------------------------------------------------------------
func _sqlConnectionConfig(link *sql.DB) {
	_, err := link.Exec("SET time_zone = 'Europe/London'")
	if err != nil {
		log.Errorln(err)
	}

	// source: https://www.alexedwards.net/blog/configuring-sqldb
	//link.SetMaxOpenConns(5) // Caused bottle neck on larger traffic site
	link.SetConnMaxIdleTime(2)
	link.SetConnMaxLifetime(1 * time.Hour)
}
