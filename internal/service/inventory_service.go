package service

import "hot-coffee/internal/dal"

type InventoryServiceInterface interface {
}

type inventoryService struct {
	inventoryAccess dal.InventoryDalInterface
}

func NewInventoryService(inventoryDal dal.InventoryDalInterface) *inventoryService {
	return &inventoryService{inventoryAccess: inventoryDal}
}
