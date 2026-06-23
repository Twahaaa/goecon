package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Twahaaa/goecom/internal/adapters/postgresql/sqlc"
	"github.com/Twahaaa/goecom/internal/orders"
	"github.com/Twahaaa/goecom/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	productService := products.NewService(repo.New(app.db))
	productsHandler := products.NewHandler(productService)
	r.Route("/products", func(r chi.Router) {
		r.Get("/", productsHandler.ListProducts)
		r.Get("/{id}", productsHandler.GetProductById)
		r.Post("/", productsHandler.CreateProduct)
	})

	orderService := orders.NewService(repo.New(app.db))
	ordersHandler := orders.NewHandler(orderService)
	r.Route("/orders", func(r chi.Router) {
		r.Get("/", ordersHandler.ListOrders)
		r.Get("/{id}", ordersHandler.GetOrderById)
	})
	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has strated at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	//logger
	db *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
