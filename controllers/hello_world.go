package controllers

import (
	"fmt"
	"net/http"

	"github.com/srfbogomolov/repository-db/models"
)

type BaseHandler struct {
	userRepo models.UserRepository
}

func NewBaseHandler(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		userRepo: userRepo,
	}
}

func (h *BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	if user, err := h.userRepo.FindByID(1); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, world"))
}
