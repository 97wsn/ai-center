package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConf struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}

type CenterDb *sqlx.DB

func NewDBConf(conf *Config) CenterDb {
	db, err := sqlx.Connect("mysql", conf.DB.Dsn)
	if err != nil {
		log.Fatalf("[config] connect db err:%s", err.Error())
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db
}
