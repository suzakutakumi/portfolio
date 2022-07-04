package db

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// SQL用のInsert,Update用の関数
func Push(sqlCommand string, args ...interface{}) error {
	dbx, dberr := sqlx.Open("mysql", os.Getenv("DBUser")+":"+os.Getenv("DBPass")+"@tcp(mysql:3306)/portfolio")
	if dberr != nil {
		return dberr
	}
	defer dbx.Close()
	dbx.SetConnMaxLifetime(time.Minute * 3)
	dbx.SetMaxOpenConns(10)
	dbx.SetMaxIdleConns(10)

	tx := dbx.MustBegin()
	tx.MustExec(sqlCommand, args...)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// SQL用のselect用の関数
func Select(rows interface{}, sqlCommand string, args ...interface{}) error {
	// dbx, dberr := sqlx.Open("mysql", "Curator:oldmap@tcp(localhost:3306)/old_map_explorer")
	dbx, dberr := sqlx.Open("mysql", "Curator:oldmap@tcp(mysql:3306)/old_map_explorer")
	if dberr != nil {
		return dberr
	}
	defer dbx.Close()
	dbx.SetConnMaxLifetime(time.Minute * 3)
	dbx.SetMaxOpenConns(10)
	dbx.SetMaxIdleConns(10)

	if err := dbx.Select(rows, sqlCommand, args...); err != nil {
		return err
	}
	return nil
}
