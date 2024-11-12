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
		w.WriteHeader(http.StatusInternalServerError)
		models.Logger.Error("Failed to load inventory")
		service.ErrorHandling("Failed to load inventory", w)

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
		w.WriteHeader(http.StatusBadRequest)
		models.Logger.Error("Invalid request path")
		service.ErrorHandling("Invalid request path", w)
		return
	}
	item, err := h.inventoryService.GetByID(path[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		models.Logger.Error("Inventory item not found")
		service.ErrorHandling("Inventory item not found", w)
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
		w.WriteHeader(http.StatusBadRequest)
		models.Logger.Error("Error reading request body")
		service.ErrorHandling("Error reading request body", w)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		models.Logger.Error("Invalid JSON")
		service.ErrorHandling("Invalid JSON", w)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		models.Logger.Error("Invalid request path")
		service.ErrorHandling("Invalid request path", w)
		return
	}
	item, err := h.inventoryService.Update(path[2], updatedItem)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		models.Logger.Error("Menu item not found")
		service.ErrorHandling("Menu item not found", w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
		return
	}
}

func (h *InventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.inventoryService.Delete(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		models.Logger.Error("Inventory item not found")
		service.ErrorHandling("Inventory item not found", w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
