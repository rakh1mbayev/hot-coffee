package handler

import (
	"hot-coffee/internal/service"
	"net/http"
)

type InventoryHandler struct {
	inventoryService service.InventoryServiceInterface
}

func NewInventoryHandler(inventoryService service.InventoryServiceInterface) *InventoryHandler {
	return &InventoryHandler{inventoryService: inventoryService}
}

func (h *InventoryHandler) PostInventory(w http.ResponseWriter, r *http.Request) {
	
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) GetInventoryID(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) PutInventoryID(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) DeleteInventoryID(w http.ResponseWriter, r *http.Request) {

}
