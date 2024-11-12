package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	"hot-coffee/models"
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
	h.inventoryService.Add(newInventoryItem)
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) GetInventoryID(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) PutInventoryID(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) DeleteInventoryID(w http.ResponseWriter, r *http.Request) {

}
