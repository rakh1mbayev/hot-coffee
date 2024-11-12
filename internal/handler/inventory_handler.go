package handler

import (
	"encoding/json"
	"fmt"
	"hot-coffee/internal/service"
	"hot-coffee/models"
	"io"
	"net/http"
)

type InventoryHandler struct {
	inventoryService service.InventoryServiceInterface
}

func NewInventoryHandler(inventoryService service.InventoryServiceInterface) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) PostInventory(w http.ResponseWriter, r *http.Request) {
	var newInventoryItem models.InventoryItem
	json.NewDecoder(r.Body).Decode(&newInventoryItem)
	if err := h.inventoryService.Add(newInventoryItem); err != nil {
		// error
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	items, err := h.inventoryService.Get()
	if err != nil {
		// error
		http.Error(w, "Failed to load menu", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(items); err != nil {
		return
	}
}

func (h *InventoryHandler) GetInventoryID(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) PutInventoryID(w http.ResponseWriter, r *http.Request) {
	var updateItem models.InventoryItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &updateItem); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	path := r.PathValue("/inventory/{id}")

	item, err := h.inventoryService.Update(path, updateItem)
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
		return
	}
}

func (h *InventoryHandler) DeleteInventoryID(w http.ResponseWriter, r *http.Request) {
}
