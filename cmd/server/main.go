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
	relationshipDB := database.NewRelationship(neoDriver)
	println("Database configured")

	webserver := webserver.NewWebServer(config.WebserverPort)
	webPersonHandler := handlers.NewWebPersonHandler(personDB)
	webRelationshipHandler := handlers.NewWebRelationshipHandler(relationshipDB)

	// Person routes
	webserver.AddMethod("POST", "/person", webPersonHandler.Create)
	webserver.AddMethod("GET", "/person/{uuid}", webPersonHandler.FindOne)
	webserver.AddMethod("GET", "/person/name/{name}", webPersonHandler.FindOne)
	webserver.AddMethod("PUT", "/person/{uuid}", webPersonHandler.Update)
	webserver.AddMethod("DELETE", "/person/{uuid}", webPersonHandler.Delete)
	webserver.AddMethod("GET", "/person/{uuid}/ancestors", webPersonHandler.GetAncestors)
	webserver.AddMethod("GET", "/person/{uuid}/family", webPersonHandler.GetFamily)

	// Relationship routes
	webserver.AddMethod("POST", "/relationship", webRelationshipHandler.CreateIsParent)
	webserver.AddMethod("GET", "/relationship/{start}/bacon_number/{end}", webRelationshipHandler.GetBaconNumber)

	webserver.Start()
}
