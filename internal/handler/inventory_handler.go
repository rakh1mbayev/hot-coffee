package handler

import (
	"net/http"

	"hot-coffee/internal/service"
)

type Inventory_handle interface{}

type Inventoryhandler struct {
	Inventory_serve service.Inventory_serv
}

func NewInventoryHandler(Inventory_serve *service.Inventory_serv) *Inventoryhandler {
	return &Inventoryhandler{Inventory_serve: *Inventory_serve}
}

func (i Inventoryhandler) PostInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
}

func GetInventoryID(w http.ResponseWriter, r *http.Request) {
}

func PutInventoryID(w http.ResponseWriter, r *http.Request) {
}

func DeleteInventoryID(w http.ResponseWriter, r *http.Request) {
}
