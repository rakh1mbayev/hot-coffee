package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	model "hot-coffee/models"
	"io"
	"net/http"
	"strings"
)

type OrderHandler struct {
	service service.OrdersService
}

func NewOrdersHandler(service service.OrdersService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (o *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newOrderItem model.OrderItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Error reading request body")
		service.ErrorHandling("Error reading request body", w)
		return
	}
	if err := json.Unmarshal(body, &newOrderItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid JSON")
		service.ErrorHandling("Invalid JSON", w)
		return
	}
	if err := o.service.Add(newOrderItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Failed to add order item")
		service.ErrorHandling("Failed to add order item", w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
func (o *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	items, err := o.service.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		model.Logger.Error("Failed to load orders")
		service.ErrorHandling("Failed to load orders", w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(items); err != nil {
		return
	}
}

func (o *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid request path")
		service.ErrorHandling("Invalid request path", w)
		return
	}
	item, err := o.service.GetByID(path[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Order item not found")
		service.ErrorHandling("Order item not found", w)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
		return
	}
}

func (o *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updatedItem model.OrderItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Error reading request body")
		service.ErrorHandling("Error reading request body", w)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid JSON")
		service.ErrorHandling("Invalid JSON", w)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid request path")
		service.ErrorHandling("Invalid request path", w)
		return
	}

	if 	err := o.service.Update(path[2], updatedItem); err != nil {
		w.WriteHeader(http.StatusNotFound)
		model.Logger.Error("Order item not found")
		service.ErrorHandling("Order item not found", w)
		return
	}
}

func (o *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	err := o.service.Delete(path[2])
	if err != nil {
		http.Error(w, "Order item not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
func (o *OrderHandler) Close(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	err := o.service.Close(path[2])
	if err != nil {
		http.Error(w, "Order item not found", http.StatusNotFound)
		return
	}
}
