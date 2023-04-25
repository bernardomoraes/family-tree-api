package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
	usecase "github.com/bernardomoraes/family-tree/internal/usecase/relationship"
)

type RelationshipHandler struct {
	RelationshipDB entity.RelationshipRepositoryInterface
}

func NewWebRelationshipHandler(db entity.RelationshipRepositoryInterface) *RelationshipHandler {
	return &RelationshipHandler{
		RelationshipDB: db,
	}
}

func (h *RelationshipHandler) CreateIsParent(rw http.ResponseWriter, r *http.Request) {
	var relationship dto.CreateParentRelationshipInput

	err := json.NewDecoder(r.Body).Decode(&relationship)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	if (relationship.ParentUUID == "") || (relationship.ChildUUID == "") {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "Fields parent and child with valid UUID are required",
		})
		return
	}

	fmt.Println(relationship)

	err = usecase.NewCreateRelationshipUseCase(h.RelationshipDB).Execute(r.Context(), &relationship)
	if err != nil {
		switch err {
		case entity.ErrRelationshipAlreadyExists:
			rw.WriteHeader(http.StatusForbidden)
		case entity.ErrStartAndEndIsRequired:
			rw.WriteHeader(http.StatusBadRequest)
		case entity.ErrInvalidType:
			rw.WriteHeader(http.StatusBadRequest)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "Relationship created",
	})
}
