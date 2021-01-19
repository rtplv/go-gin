package connections

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() error {
	if DB != nil {
		return errors.New("DB instance already initialized")
	}

	connLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: 5 * time.Second,
			LogLevel:      gormLogger.Error,
			Colorful:      true,
		},
	)

	conn, err := openPgConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		&gorm.Config{
			Logger: connLogger,
		},
	)

	if conn != nil {
		sqlDB, err := conn.DB()

		if err != nil {
			logrus.Fatal(err, "Error on DB generic interface call")
		}

		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(300)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	DB = conn

	return err
}

func openPgConnection(host string, username string, password string, database string, port string, config *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host,
		username,
		password,
		database,
		port,
	)

	return gorm.Open(postgres.Open(dsn), config)
}
