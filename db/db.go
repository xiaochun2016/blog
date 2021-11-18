package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

var db *sqlx.DB

func Init() (err error) {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		return
	}
	dbUsername := cfg.Section("db").Key("username").String()
	dbPwd := cfg.Section("db").Key("password").String()
	dbHost := cfg.Section("db").Key("host").String()
	dbPort := cfg.Section("db").Key("port").String()
	databases := cfg.Section("db").Key("databases").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPwd, dbHost, dbPort, databases)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return
}
