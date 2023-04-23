package handlers

import "github.com/bernardomoraes/family-tree/internal/entity"

type UserHandler struct {
	UserDB entity.UserRepositoryInterface
}

func NewUserHandler(userDB entity.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) Create() {

}
