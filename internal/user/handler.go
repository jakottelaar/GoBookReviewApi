package user

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil { // Better to not return the error to the client
		common.ServerErrorResponse(w, r, err)
		return
	}

	id := claims["user_id"]
	user, err := h.service.FindById(id.(string))
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	resp := &GetUserProfileResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	common.WriteJSON(w, http.StatusOK, common.Envelope{"user": resp}, nil)

}
