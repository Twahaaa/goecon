package products

import (
	"context"

	repo "github.com/Twahaaa/goecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductById(ctx context.Context, id int64) (repo.Product, error)
	CreateProduct(ctx context.Context, body CreateProductInput) (repo.Product, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

type CreateProductInput struct {
    Name        string `json:"name"`
    PriceInCents int32 `json:"price_in_cents"`
    Quantity    int32  `json:"quantity"`
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductsById(ctx, id)
}

func (s *svc) CreateProduct(ctx context.Context, input CreateProductInput) (repo.Product, error){
	return s.repo.CreateProduct(ctx, repo.CreateProductParams{
		Name: input.Name,
		PriceInCents: input.PriceInCents,
		Quantity: input.Quantity,
	})
}