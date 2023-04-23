package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
	"github.com/go-chi/chi/v5"
)

type PersonHandler struct {
	PersonDB entity.PersonRepositoryInterface
}

func NewWebPersonHandler(db entity.PersonRepositoryInterface) *PersonHandler {
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
	fmt.Println("person:", person.UUID)
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
func (h *PersonHandler) DeletePerson(rw http.ResponseWriter, r *http.Request) {
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

	err = h.PersonDB.Delete(r.Context(), personFinded.UUID)
	if err != nil {
		fmt.Println("err:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "Person deleted successfully",
	})
}
