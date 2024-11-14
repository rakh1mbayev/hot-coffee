package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
	model "hot-coffee/models"
)

type OrderHandler struct {
	service service.OrdersService
}

func NewOrdersHandler(service service.OrdersService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (o *OrderHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newOrderItem model.Order
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dal.SendError("Error reading request body", http.StatusBadRequest, w)
		return
	}
	if err := json.Unmarshal(body, &newOrderItem); err != nil {
		dal.SendError("Invalid JSON", http.StatusBadRequest, w)
		return
	}
	if err := o.service.Add(newOrderItem); err != nil {
		dal.SendError("Failed to add order item", http.StatusInternalServerError, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (o *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	items, err := o.service.Get()
	if err != nil {
		dal.SendError("Failed to load orders", http.StatusInternalServerError, w)
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
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	item, err := o.service.GetByID(path[2])
	if err != nil {
		dal.SendError("Order item not found", http.StatusNotFound, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
		return
	}
}

func (o *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updatedItem model.Order
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dal.SendError("Error reading request body", http.StatusBadRequest, w)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		dal.SendError("Invalid JSON", http.StatusBadRequest, w)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	if err := o.service.Update(path[2], updatedItem); err != nil {
		dal.SendError("Order item not found", http.StatusNotFound, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
}

func (o *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	err := o.service.Delete(path[2])
	if err != nil {
		dal.SendError("Order item not found", http.StatusNotFound, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (o *OrderHandler) Close(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	err := o.service.Close(path[2])
	if err != nil {
		dal.SendError("Order item not found", http.StatusNotFound, w)
		return
	}
}
