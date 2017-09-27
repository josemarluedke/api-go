package db

import (
	"github.com/go-pg/pg"
	"github.com/josemarluedke/api/config"
)

var Conn *pg.DB

func init() {
	Conn = connect()
}

func connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     config.Config.Database.Addr,
		User:     config.Config.Database.Username,
		Password: config.Config.Database.Password,
		Database: config.Config.Database.DatabaseName,
	})
}
