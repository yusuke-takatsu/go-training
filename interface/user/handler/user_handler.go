package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/yusuke-takatsu/go-training/interface/user/dto"
	"github.com/yusuke-takatsu/go-training/service/user/usecase"
	"net/http"
)

type Handler struct {
	service *usecase.UseCase
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Image    string `json:"image,omitempty"`
}

type RegisterUserResponse struct {
	Message string `json:"message"`
}

func NewHandler(s *usecase.UseCase) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	input := dto.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Image:    req.Image,
	}
	if err := h.service.Register(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&RegisterUserResponse{
		Message: "ユーザー登録が完了しました。",
	})
}
