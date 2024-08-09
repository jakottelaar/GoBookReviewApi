package book

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type BookHandler struct {
	bookService BookService
	userService user.UserService
}

func NewBookHandler(bookService BookService, userService user.UserService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
		userService: userService,
	}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body CreateBookRequest true "Book details"
// @Success 201 {object} CreateBookResponse
// @Router /books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	id := claims["user_id"]
	user, err := h.userService.FindById(id.(string))
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	var req CreateBookRequest

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

	createdBook, err := h.bookService.Create(&req, user.ID.String())

	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}

	resp := CreateBookResponse{
		ID:            createdBook.ID.String(),
		Title:         createdBook.Title,
		Author:        createdBook.Author,
		PublishedYear: createdBook.PublishedYear,
		ISBN:          createdBook.ISBN,
		CreatedAt:     createdBook.CreatedAt,
		UserID:        createdBook.UserId.String(),
	}

	common.WriteJSON(w, http.StatusCreated, common.Envelope{"book": resp}, nil)

}

// GetBookById godoc
// @Summary Get a book by ID
// @Description Get a book by the provided ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} GetBookResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r, "id")

	if err != nil {
		common.BadRequestResponse(w, r, err)
		return
	}

	book, err := h.bookService.GetBookById(id)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	resp := GetBookResponse{
		ID:            book.ID.String(),
		Title:         book.Title,
		Author:        book.Author,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
		UserID:        book.UserId.String(),
	}

	common.WriteJSON(w, http.StatusOK, common.Envelope{"book": resp}, nil)

}

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update a book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID" format(uuid)
// @Param book body UpdateBookRequest true "Book details"
// @Success 200 {object} UpdateBookResponse
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r, "id")

	if err != nil {
		common.BadRequestResponse(w, r, err)
		return
	}

	var req UpdateBookRequest

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

	book, err := h.bookService.Update(id, &req)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	resp := UpdateBookResponse{
		ID:            book.ID.String(),
		Title:         book.Title,
		Author:        book.Author,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		UpdatedAt:     book.UpdatedAt,
	}

	common.WriteJSON(w, http.StatusOK, common.Envelope{"book": resp}, nil)

}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book by the provided ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} interface{}
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := common.GetIdFromRequest(r, "id")

	if err != nil {
		common.BadRequestResponse(w, r, err)
		return
	}

	err = h.bookService.Delete(id)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	common.WriteJSON(w, http.StatusOK, common.Envelope{"message": "book deleted"}, nil)

}
