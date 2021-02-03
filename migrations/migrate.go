package main

import (
	"app/connections"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = connections.InitDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch os.Args[1] {
	case "up":
		up()
	case "down":
		down()
	}
}

func up() {
	// migrations up list
	ExampleMigration{}.up()
}

func down() {
	// migrations down list
	ExampleMigration{}.down()
}
