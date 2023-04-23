package main

import (
	"context"
	"fmt"

	"github.com/bernardomoraes/family-tree/configs"
	"github.com/bernardomoraes/family-tree/internal/infra/database"
	"github.com/bernardomoraes/family-tree/internal/infra/web/handlers"
	"github.com/bernardomoraes/family-tree/internal/infra/web/webserver"
	"github.com/bernardomoraes/family-tree/pkg/helpers"
)

func main() {
	config, _ := configs.LoadConfig(".")
	ctx := context.Background()
	fmt.Println("Server configured")

	neoDriver := helpers.Driver(config.DBUri, config.DBUser, config.DBPassword)
	defer neoDriver.Close(ctx)
	println("Driver configured")

	personDB := database.NewPerson(neoDriver)
	println("Database configured")

	webserver := webserver.NewWebServer(config.WebserverPort)
	webPersonHandler := handlers.NewWebPersonHandler(personDB)

	// webserver.AddMethod("GET", "/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Add("Content-Type", "application/json")
	// 	rw.Header().Set("Access-Control-Allow-Origin", "*")
	// 	rw.WriteHeader(http.StatusOK)

	// 	json.NewEncoder(rw).Encode(map[string]string{
	// 		"message": "Welcome to the Family Tree API",
	// 	})
	// })
	webserver.AddMethod("POST", "/person", webPersonHandler.CreatePerson)
	webserver.AddMethod("GET", "/person/{uuid}", webPersonHandler.GetPerson)
	webserver.AddMethod("PUT", "/person/{uuid}", webPersonHandler.UpdatePerson)
	webserver.AddMethod("DELETE", "/person/{uuid}", webPersonHandler.DeletePerson)

	webserver.Start()
}
