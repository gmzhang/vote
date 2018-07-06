package utils

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(host, user, password, port, dbName string) {
	connStr := user + ":" + password + "@tcp(" + host +
		":" + port + ")/" + dbName + "?parseTime=true"
	var err error
	DB, err = sqlx.Open("mysql", connStr)

	if err == nil {
		err = DB.Ping()
	}

	if err != nil {
		Logger.Fatalf("system db connect error: %#v", err)
	}
}

//RowExists check whether a row exist
func RowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := DB.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		Logger.Errorf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}

func RowCount(query string, args ...interface{}) int {
	var count int
	err := DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		Logger.WithField("err", err).Errorln("no rows or error")
		return count
	}

	return count
}
