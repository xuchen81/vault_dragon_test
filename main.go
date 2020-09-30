package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"vault_dragon_test/controllers"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"

	"vault_dragon_test/config"
	"vault_dragon_test/database"
	"vault_dragon_test/models"
	"vault_dragon_test/repositories"
)

func initDBClient() error {
	if _, present := os.LookupEnv("VD_ENV"); !present {
		return errors.New("missing environment setting")
	}
	configFile, err := config.GetConfigFilePath()
	if err != nil {
		return err
	}

	config.Vi = viper.New()
	config.Vi.SetConfigFile(configFile)
	config.Vi.ReadInConfig()
	// Connecting to database
	_, err = database.InitDBClient(
		config.Vi.GetString("database.db_user"),
		config.Vi.GetString("database.db_password"),
		config.Vi.GetString("database.db_name"),
	)
	if err != nil {
		return err
	}

	// migration
	database.Client.AutoMigrate(&models.Key{})
	database.Client.AutoMigrate(&models.Value{})

	return nil

}

func main() {
	err := initDBClient()
	if err != nil {
		panic("failed to connect to database")
	}

	// initialize repositories
	repositories.InitRepositories()
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!!")
	})
	r.HandleFunc("/object", controllers.CreateKeyValueHandler).Methods("POST")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	http.ListenAndServe(":3000", n)

}
