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
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &newOrderItem); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := o.service.Add(newOrderItem); err != nil {
		http.Error(w, "Failed to add order item", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
func (o *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	items, err := o.service.Get()
	if err != nil {
		http.Error(w, "Failed to load orders", http.StatusInternalServerError)
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
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	item, err := o.service.GetByID(path[2])
	if err != nil {
		http.Error(w, "Order item not found", http.StatusNotFound)
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
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	item, err := o.service.Update(path[2], updatedItem)
	if err != nil {
		http.Error(w, "Order item not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
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
