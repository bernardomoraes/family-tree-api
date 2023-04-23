package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bernardomoraes/family-tree/configs"
	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
	"github.com/bernardomoraes/family-tree/internal/infra/database"
	"github.com/bernardomoraes/family-tree/pkg/helpers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// enum of basic environment variables fallbacks
const (
	DefaultPort = "8080"
)

func configureEndpoints(router chi.Router, driver helpers.AvailableDatabaseDrivers) {
	personDB := database.NewPerson(driver)
	personHandler := NewPersonHandler(personDB)

	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.WriteHeader(http.StatusOK)

		json.NewEncoder(rw).Encode(map[string]string{
			"message": "Welcome to the Family Tree API",
		})
	})

	router.Post("/person", func(rw http.ResponseWriter, r *http.Request) {
		acceptHeader := r.Header.Values("Accept")
		fmt.Println("acceptHeader:", acceptHeader)

		personHandler.CreatePerson(rw, r)
	})

	router.Get("/person/{uuid}", func(rw http.ResponseWriter, r *http.Request) {
		// rw.Header().Add("Content-Type", "application/json")
		personHandler.GetPerson(rw, r)
	})
	router.Put("/person/{uuid}", func(rw http.ResponseWriter, r *http.Request) {
		// rw.Header().Add("Content-Type", "application/json")
		personHandler.UpdatePerson(rw, r)
	})
}

func configureRouter(portNumber string, driverDatabase helpers.AvailableDatabaseDrivers) (string, *chi.Mux) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	// r.Use(middleware.Compress(5, "application/json"))
	if portNumber == "" {
		portNumber = DefaultPort
	}

	configureEndpoints(r, driverDatabase)

	return ":" + portNumber, r
}

type PersonHandler struct {
	PersonDB database.PersonInterface
}

func NewPersonHandler(db database.PersonInterface) *PersonHandler {
	return &PersonHandler{
		PersonDB: db,
	}
}

func (h *PersonHandler) CreatePerson(rw http.ResponseWriter, r *http.Request) {
	var person dto.CreatePersonInput

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewPerson(person.Name)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	personCreated, err := h.PersonDB.Create(r.Context(), p)

	if err != nil {
		fmt.Println("err:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&personCreated)
}
func (h *PersonHandler) GetPerson(rw http.ResponseWriter, r *http.Request) {
	person := dto.FindPersonInput{
		UUID: chi.URLParam(r, "uuid"),
	}

	personFinded, err := h.PersonDB.FindByUUID(r.Context(), person.UUID)
	if err != nil {
		fmt.Println("Database Error:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if personFinded == nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "Person not found with UUID: " + person.UUID,
		})
		return
	}
	fmt.Println("personFinded:", personFinded)

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusFound)
	json.NewEncoder(rw).Encode(&personFinded)
}
func (h *PersonHandler) UpdatePerson(rw http.ResponseWriter, r *http.Request) {
	person := dto.FindPersonInput{
		UUID: chi.URLParam(r, "uuid"),
	}
	personFinded, err := h.PersonDB.FindByUUID(r.Context(), person.UUID)
	if err != nil {
		fmt.Println("Database Error:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if personFinded == nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "Person not found with UUID: " + person.UUID,
		})
		return
	}
	fmt.Println("personFinded:", personFinded)

	var personInput dto.UpdatePersonInput
	err = json.NewDecoder(r.Body).Decode(&personInput)
	if err != nil {
		fmt.Println("err:", err)
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "Missing required fields",
		})
		return
	}

	personFinded.Name = personInput.Name

	personUpdated, err := h.PersonDB.Update(r.Context(), personFinded)

	// output := dto.UpdatePersonOutput{...personUpdated}
	var output dto.UpdatePersonOutput
	output.Name = personUpdated.Name
	output.UUID = personUpdated.UUID
	output.CreatedAt = personUpdated.CreatedAt
	output.UpdatedAt = personUpdated.UpdatedAt

	// fmt.Println("personCreated:", personCreated)

	if err != nil {
		fmt.Println("err:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&output)
}

func main() {
	config, _ := configs.LoadConfig(".")
	ctx := context.Background()
	fmt.Println("Server configured")

	neoDriver := helpers.Driver(config.DBUri, config.DBUser, config.DBPassword)
	defer neoDriver.Close(ctx)
	println("Driver configured")

	err := http.ListenAndServe(configureRouter(config.WebserverPort, neoDriver))
	if err != nil {
		panic(err)
	}
}
