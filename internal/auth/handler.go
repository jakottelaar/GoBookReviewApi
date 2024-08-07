package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest

	err := common.ReadJSON(w, r, &req)
	if err != nil {
		common.BadRequestResponse(w, r, err)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(req)

	if err != nil {
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}

		common.FailedValidationResponse(w, r, errors)
		return
	}

	createdUser, err := h.service.Register(&req)
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	resp := RegisterResponse{
		ID:        createdUser.ID.String(),
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}

	common.WriteJSON(w, http.StatusCreated, common.Envelope{"user": resp}, nil)

}
