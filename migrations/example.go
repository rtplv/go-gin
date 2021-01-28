package main

import (
	"app/connections"
	"app/core/models"
	"github.com/sirupsen/logrus"
)

type ExampleMigration struct {
}

func (m ExampleMigration) up() {
	model := &models.Example{}
	tableExist := connections.DB.Migrator().HasTable(model)

	if tableExist {
		logrus.Fatalf("Table %s already exist", model.TableName())
	}

	err := connections.DB.Migrator().AutoMigrate(&model)

	if err != nil {
		logrus.Fatal("Error on autoMigrate", err)
	}

	logrus.Info("Successfully migrate!")
}

func (m ExampleMigration) down() {
	model := models.Example{}
	tableExist := connections.DB.Migrator().HasTable(model)

	if !tableExist {
		logrus.Fatalf("Table %s does not exist", model.TableName())
	}

	err := connections.DB.Migrator().DropTable(model)

	if err != nil {
		logrus.Fatal("Error on autoMigrate", err.Error())
	}

	logrus.Info("Successfully migrate down!")
}
