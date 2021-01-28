package migrations

import (
	"app/core/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ReportsMigration struct {
	DB *gorm.DB
}

func (m ReportsMigration) Up() {
	model := &models.Example{}
	tableExist := m.DB.Migrator().HasTable(model)

	if tableExist {
		logrus.Fatalf("Table %s already exist", model.TableName())
	}

	err := m.DB.Migrator().AutoMigrate(&model)

	if err != nil {
		logrus.Fatal("Error on autoMigrate", err)
	}

	logrus.Info("Successfully migrate!")
}

func (m ReportsMigration) Down() {
	model := models.Example{}
	tableExist := m.DB.Migrator().HasTable(model)

	if !tableExist {
		logrus.Fatalf("Table %s does not exist", model.TableName())
	}

	err := m.DB.Migrator().DropTable(model)

	if err != nil {
		logrus.Fatal("Error on autoMigrate", err.Error())
	}

	logrus.Info("Successfully migrate down!")
}
