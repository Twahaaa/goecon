package orders

import (
	"context"

	repo "github.com/Twahaaa/goecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListOrders(ctx context.Context) ([]repo.Order, error)
	GetOrderById(ctx context.Context, id int64) ([]repo.FindOrderByIdRow, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListOrders(ctx context.Context) ([]repo.Order, error) {
	return s.repo.ListOrders(ctx)
}

func (s *svc) GetOrderById(ctx context.Context, id int64) ([]repo.FindOrderByIdRow, error) {
	return s.repo.FindOrderById(ctx, id)
}
