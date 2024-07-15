package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"stardustcode/backend/internal/database"
	internalMiddleware "stardustcode/backend/internal/middlewares"
	"stardustcode/backend/internal/projects/parcus/controllers"
	"stardustcode/backend/internal/projects/parcus/services"

	"github.com/antonlindstrom/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
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

	validate := validator.New(validator.WithRequiredStructEnabled())

	gob.Register(pgtype.Timestamp{})
	gob.Register(time.Time{})

	store, err := pgstore.NewPGStore(os.Getenv("DATABASE_URL"), []byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		log.Fatalf(err.Error())
	}

	store.Path = "/"
	defer store.Close()
	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	authService := services.AuthService{DbPool: dbpool}
	accountService := services.AccountService{DbPool: dbpool}
	budgetService := services.BudgetService{DbPool: dbpool}
	categoryService := services.CategoryService{DbPool: dbpool}
	transactionService := services.TransactionService{DbPool: dbpool}
	recurringTransactionService := services.RecurringTransactionService{DbPool: dbpool}

	authController := controllers.AuthController{Service: &authService, Validator: validate, Store: store}
	accountController := controllers.AccountController{Service: &accountService, Validator: validate}
	budgetController := controllers.BudgetController{Service: &budgetService, Validator: validate}
	categoryController := controllers.CategoryController{Service: &categoryService, Validator: validate, Store: store}
	transactionController := controllers.TransactionController{Service: &transactionService, Validator: validate}
	recurringTransactionController := controllers.RecurringTransactionController{Service: &recurringTransactionService, Validator: validate}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})

	r.Route("/parcus/v1", func(r chi.Router) {
		r.Use(internalMiddleware.JsonHeader)
		r.Use(internalMiddleware.GetUser(store))

		r.Get("/auth", authController.Get)
		r.Post("/login", authController.Login)

		r.Group(func(r chi.Router) {
			r.Use(internalMiddleware.Authenticated)

			r.Post("/logout", authController.Logout)

			r.Get("/accounts", accountController.Get)
			r.Post("/accounts", accountController.Post)
			r.Put("/accounts/{id}", accountController.Put)
			r.Delete("/accounts/{id}", accountController.Delete)

			r.Get("/budgets", budgetController.Get)
			r.Post("/budgets", budgetController.Post)
			r.Put("/budgets/{id}", budgetController.Put)
			r.Delete("/budgets/{id}", budgetController.Delete)

			r.Get("/categories", categoryController.Get)
			r.Post("/categories", categoryController.Post)
			r.Put("/categories/{id}", categoryController.Put)
			r.Delete("/categories/{id}", categoryController.Delete)

			r.Get("/transactions", transactionController.Get)
			r.Post("/transactions", transactionController.Post)
			r.Put("/transactions/{id}", transactionController.Put)
			r.Delete("/transactions/{id}", transactionController.Delete)

			r.Get("/recurring_transactions", recurringTransactionController.Get)
			r.Post("/recurring_transactions", recurringTransactionController.Post)
			r.Put("/recurring_transactions/{id}", recurringTransactionController.Put)
			r.Delete("/recurring_transactions/{id}", recurringTransactionController.Delete)
		})

	})

	http.ListenAndServe(":3333", r)
}
