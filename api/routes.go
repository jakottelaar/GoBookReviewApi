package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jakottelaar/gobookreviewapp/internal/book"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
	"github.com/jakottelaar/gobookreviewapp/pkg/database"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition
	))

	// Setup book services
	db := database.GetDB()
	bookRepository := book.NewBookRepository(db)
	bookService := book.NewBookService(bookRepository)
	bookHandler := book.NewBookHandler(bookService)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		err := common.WriteJSON(w, http.StatusOK, common.Envelope{"message": "Health Check OK"}, nil)
		if err != nil {
			common.ServerErrorResponse(w, r, err)
			return
		}
	})

	// API routes
	r.Route("/v1/api", func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Post("/", bookHandler.CreateBook)
			r.Get("/{id}", bookHandler.GetBookById)
			r.Put("/{id}", bookHandler.UpdateBook)
			r.Delete("/{id}", bookHandler.DeleteBook)
		})
	})

	return r
}
