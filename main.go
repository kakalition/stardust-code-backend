package main

import (
	"log"
	"net/http"
	"time"

	"stardustcode/backend/internal/database"
	internalMiddleware "stardustcode/backend/internal/middlewares"
	"stardustcode/backend/internal/projects/parcus/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Database{}
	dbpool := db.GetDatabaseConnection()
	defer dbpool.Close()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	categoryController := controllers.CategoryController{DbPool: dbpool}

	r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})

	r.Route("/parcus/v1", func(r chi.Router) {
		r.Use(internalMiddleware.JsonHeader)
		r.Get("/categories", categoryController.Get)
		r.Post("/categories", categoryController.Post)
		r.Put("/categories/{id}", categoryController.Put)
		r.Delete("/categories/{id}", categoryController.Delete)
	})

	transactionController := controllers.TransactionController{DbPool: dbpool}

	r.Route("/parcus/v1", func(r chi.Router) {
		r.Use(internalMiddleware.JsonHeader)
		r.Get("/transactions", transactionController.Get)
		r.Post("/transactions", transactionController.Post)
		r.Put("/transactions/{id}", transactionController.Put)
		r.Delete("/transactions/{id}", transactionController.Delete)
	})

	http.ListenAndServe(":3333", r)
}
