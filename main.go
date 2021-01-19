package main

import (
	"app/routers/v1"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// @title App
// @version 1.0
// @description App description
// @BasePath /
func main() {
	// Bootstrap
	loadDotEnv()
	configureLogger()
	// Optional: Init database driver
	//err := connections.InitDB()

	// Set server mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Init routes
	router := gin.New()
	version1 := router.Group("/v1")
	v1.InitRoutes(version1)

	err := router.Run()

	if err != nil {
		logrus.Fatal(err)
	}
}

func loadDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}
}

func configureLogger() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	logrus.SetFormatter(customFormatter)
}
