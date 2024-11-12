package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	"hot-coffee/models"
	"io"
	"net/http"
	"strings"
)

type InventoryHandler struct {
	inventoryService service.InventoryServiceInterface
}

func NewInventoryHandler(inventoryService service.InventoryServiceInterface) *InventoryHandler {
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
		http.Error(w, "Failed to load inventory", http.StatusInternalServerError)
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
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	item, err := h.inventoryService.GetByID(path[2])
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
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
	if err = h.inventoryService.Update(path[2], updatedItem); err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.inventoryService.Delete(id); err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
