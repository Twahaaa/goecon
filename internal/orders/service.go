package orders

import (
	"context"
	"fmt"

	repo "github.com/Twahaaa/goecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListOrders(ctx context.Context) ([]repo.Order, error)
	GetOrderById(ctx context.Context, id int64) ([]repo.FindOrderByIdRow, error)
	CreateOrder(ctx context.Context, input CreateOrderInput) (CreateOrderResponse, error)
}
type svc struct {
	// repository
	repo repo.Querier
}

type CreateOrderInput struct {
	CustomerId int64       `json:"customer_id"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type CreateOrderResponse struct {
	Order repo.Order
	Items []repo.OrderItem
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

func (s *svc) CreateOrder(ctx context.Context, input CreateOrderInput) (CreateOrderResponse, error) {
	order, err := s.repo.CreateOrder(ctx, input.CustomerId)
	if err != nil {
		return CreateOrderResponse{}, err
	}

	items := []repo.OrderItem{}

	for _, item := range input.Items {
		product, err := s.repo.FetchPrice(ctx, item.ProductId)
		
		if err != nil {
			return CreateOrderResponse{}, err
		}

		if product.Quantity < int32(item.Quantity) {
			return CreateOrderResponse{}, fmt.Errorf("insufficient stock for product %d", item.ProductId)
		}
		
		orderItem, err := s.repo.CreateOrderItems(ctx, repo.CreateOrderItemsParams{
			OrderID:    order.ID,
			ProductID:  item.ProductId,
			Quantity:   int32(item.Quantity),
			PriceCents: product.PriceInCents,
		})

		if err != nil {
			return CreateOrderResponse{}, err
		}

		_, err = s.repo.DecrementProductQuantity(ctx, repo.DecrementProductQuantityParams{
			ID:       item.ProductId,
			Quantity: int32(item.Quantity),
		})
		items = append(items, orderItem)
	}

	return CreateOrderResponse{
		Order: order,
		Items: items,
	}, nil
}
