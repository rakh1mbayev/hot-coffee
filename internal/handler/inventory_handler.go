package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
	"hot-coffee/models"
)

type InventoryHandler struct {
	inventoryService service.InventoryInterface
}

func NewInventoryHandler(inventoryService service.InventoryInterface) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newInventoryItem models.InventoryItem

	if err := json.NewDecoder(r.Body).Decode(&newInventoryItem); err != nil {
		return
	}

	if err := h.inventoryService.Add(newInventoryItem); err != nil {
		return
	}
}

func (h *InventoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	items, err := h.inventoryService.Get()
	if err != nil {
		dal.SendError("Failed to load inventory", http.StatusInternalServerError, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(items); err != nil {
		return
	}
}

func (h *InventoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	item, err := h.inventoryService.GetByID(path[2])
	if err != nil {
		dal.SendError("Inventory item not found", http.StatusNotFound, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
		return
	}
}

func (h *InventoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updatedItem models.InventoryItem
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
	if err = h.inventoryService.Update(path[2], updatedItem); err != nil {
		dal.SendError("Inventory item not found", http.StatusNotFound, w)
		return
	}
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.inventoryService.Delete(id); err != nil {
		dal.SendError("Inventory item not found", http.StatusNotFound, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
