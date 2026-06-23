package orders

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Twahaaa/goecom/internal/json"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler{
	return &handler{
		service: service,
	}
}

func (h *handler) ListOrders(w http.ResponseWriter, r *http.Request){
	// 1. Call the service  -> ListProduct
	orders ,err := h.service.ListOrders(r.Context())
	
	if err!=nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError) 
		return 
	}
	json.Write(w, http.StatusOK, orders)
}

func (h *handler) GetOrderById(w http.ResponseWriter, r *http.Request){
	id, err  := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError) 
		return 
	}
	product, err := h.service.GetOrderById(r.Context(), id)
	json.Write(w, http.StatusOK, product)
}

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request){
	order, err := json.Read[CreateOrderInput](r)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError) 
		return 
	}
	orderItems, err := h.service.CreateOrder(r.Context(), order)
	json.Write(w, http.StatusCreated, orderItems)
}