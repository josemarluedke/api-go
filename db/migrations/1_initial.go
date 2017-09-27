package main

import (
	"github.com/go-pg/migrations"
	"github.com/josemarluedke/api/models"
)

func up(db migrations.DB) error {
	return db.CreateTable(&models.User{}, nil)
}

func down(db migrations.DB) error {
	return db.DropTable(&models.User{}, nil)
}

func init() {
	migrations.Register(up, down)
}
