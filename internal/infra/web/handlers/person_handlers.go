package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
	usecase "github.com/bernardomoraes/family-tree/internal/usecase/person"
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

func (h *PersonHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var person dto.CreatePersonInputDTO

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := usecase.NewCreatePersonUseCase(h.PersonDB).Execute(r.Context(), &person)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&output)
}

func (h *PersonHandler) FindOne(rw http.ResponseWriter, r *http.Request) {
	person := dto.FindPersonInputDTO{
		Person: dto.Person{
			Name: chi.URLParam(r, "name"),
			UUID: chi.URLParam(r, "uuid"),
		},
	}

	personFinded, err := usecase.NewFindOnePersonUseCase(h.PersonDB).Execute(r.Context(), &person)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if personFinded == nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "Person not found",
		})
		return
	}

	rw.WriteHeader(http.StatusFound)
	json.NewEncoder(rw).Encode(&personFinded)
}

func (h *PersonHandler) Update(rw http.ResponseWriter, r *http.Request) {
	person := dto.FindPersonInputDTO{
		Person: dto.Person{
			UUID: chi.URLParam(r, "uuid"),
		},
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

	var personInput dto.UpdatePersonInputDTO
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

	output := dto.UpdatePersonOutputDTO{
		Person: dto.Person{
			Name: person.Name,
			UUID: person.UUID,
		},
		AuditTrail: dto.AuditTrail{
			CreatedAt: personUpdated.CreatedAt,
			UpdatedAt: personUpdated.UpdatedAt,
		},
	}

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

func (h *PersonHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	person := dto.DeletePersonInputDTO{
		UUID: chi.URLParam(r, "uuid"),
	}

	err := usecase.NewDeletePersonUseCase(h.PersonDB).Execute(r.Context(), &person)
	if err == nil {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "Person deleted successfully",
		})
		return
	}

	var errMessage string
	switch err.Error() {
	case "person not found":
		rw.WriteHeader(http.StatusNotFound)
		errMessage = "Person not found with UUID: " + person.UUID
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		errMessage = err.Error()
	}

	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": errMessage,
	})
}

func (h *PersonHandler) GetAncestors(rw http.ResponseWriter, r *http.Request) {
	person := dto.GetAncestorsInput{
		UUID: chi.URLParam(r, "uuid"),
	}
	if person.UUID == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": errors.New("field uuid is required").Error(),
		})
		return
	}

	ancestors, err := usecase.NewGetAncestorsUseCase(h.PersonDB).Execute(r.Context(), &person)
	if err != nil {
		switch err {
		case entity.ErrIDIsRequired:
			rw.WriteHeader(http.StatusBadRequest)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(ancestors)
}

func (h *PersonHandler) GetFamily(rw http.ResponseWriter, r *http.Request) {
	person := dto.GetAncestorsInput{
		UUID: chi.URLParam(r, "uuid"),
	}
	if person.UUID == "" {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": errors.New("field uuid is required").Error(),
		})
		return
	}

	family, err := usecase.NewGetFamilyUseCase(h.PersonDB).Execute(r.Context(), &person)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err,
		})
		return
	}

	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(family)

}
