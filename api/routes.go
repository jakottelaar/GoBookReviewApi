package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jakottelaar/gobookreviewapp/internal/auth"
	"github.com/jakottelaar/gobookreviewapp/internal/book"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
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

	// Database
	db := database.GetDB()

	//JWT
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	// Setup book services
	bookRepository := book.NewBookRepository(db)
	bookService := book.NewBookService(bookRepository)
	bookHandler := book.NewBookHandler(bookService)

	userRepository := user.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository)
	authHandler := auth.NewAuthHandler(authService)

	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		common.WriteJSON(w, http.StatusOK, common.Envelope{"message": "Health Check OK"}, nil)
	})

	// API routes
	r.Route("/v1/api", func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			r.Post("/", bookHandler.CreateBook)
			r.Get("/{id}", bookHandler.GetBookById)
			r.Put("/{id}", bookHandler.UpdateBook)
			r.Delete("/{id}", bookHandler.DeleteBook)
		})
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
			r.Post("/register", authHandler.Register)
		})
		r.Route("/users", func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Get("/profile", userHandler.GetUserProfile)
			r.Put("/profile", userHandler.UpdateUser)
			r.Delete("/", userHandler.DeleteUser)
		})
	})

	return r
}
