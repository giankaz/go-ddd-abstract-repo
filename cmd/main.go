package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	starters "mongodb.com/cmd/starters"
	common "mongodb.com/common/application"
)

func main() {
	err := run()

	if err != nil {
		common.ErrorLog.Println(err)
	}
}

func run() error {
	starters.StartLoggers()

	if err := godotenv.Load(); err != nil {
		common.ErrorLog.Fatal(".env file not found")
	}

	client, err := starters.MongoConnection()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return err
	}

	dbName := os.Getenv("MONGO_DB_NAME")

	if dbName == "" {
		common.ErrorLog.Fatal("missing environment database name")
	}

	database := client.Database(dbName)

	starters.StartRepositories(database)

	starters.StartHandlers()

	go starters.StartRoutines()

	routes()

	cancel := server()

	defer cancel()

	return nil
}
