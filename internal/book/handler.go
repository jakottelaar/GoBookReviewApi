package book

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type BookHandler struct {
	service BookService
}

func NewBookHandler(service BookService) *BookHandler {
	return &BookHandler{
		service: service,
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
	var req CreateBookRequest

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

	createdBook, err := h.service.Create(&req)

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
	}

	err = common.WriteJSON(w, http.StatusCreated, common.Envelope{"book": resp}, nil)
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}
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

	book, err := h.service.GetBookById(id)

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
	}

	err = common.WriteJSON(w, http.StatusOK, common.Envelope{"book": resp}, nil)
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}
}

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update a book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID" format(uuid)
// @Param book body UpdateBookRequest true "Book details"
// @Success 200 {object} CreateBookResponse
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

	book, err := h.service.Update(id, &req)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	resp := CreateBookResponse{
		ID:            book.ID.String(),
		Title:         book.Title,
		Author:        book.Author,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		CreatedAt:     book.CreatedAt,
	}

	err = common.WriteJSON(w, http.StatusOK, common.Envelope{"book": resp}, nil)
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}
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

	err = h.service.Delete(id)

	if err != nil {
		switch err {
		case common.ErrNotFound:
			common.NotFoundResponse(w, r)
		default:
			common.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = common.WriteJSON(w, http.StatusOK, common.Envelope{"message": "Successfully deleted book"}, nil)
	if err != nil {
		common.ServerErrorResponse(w, r, err)
		return
	}
}
