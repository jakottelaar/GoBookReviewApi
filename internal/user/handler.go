package user

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
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

// GetUserProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} interface{}
// @Router /user/profile [get]
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

// UpdateUser godoc
// @Summary Update user profile
// @Description Update the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param user body UpdateUserRequest true "User details"
// @Success 200 {object} interface{}
// @Router /user/profile [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	id := claims["user_id"]

	var req UpdateUserRequest

	err = common.ReadJSON(w, r, &req)
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

	user, err := h.service.Update(id.(string), &req)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	resp := &UpdateUserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}

	common.WriteJSON(w, http.StatusOK, common.Envelope{"user": resp}, nil)

}

// DeleteUser godoc
// @Summary Delete user profile
// @Description Delete the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} interface{}
// @Router /user/profile [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	id := claims["user_id"]
	err = h.service.Delete(id.(string))
	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	common.WriteJSON(w, http.StatusOK, common.Envelope{"message": "user deleted"}, nil)

}
