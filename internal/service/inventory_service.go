package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryServiceInterface interface {
	Add(models.InventoryItem) error
	Get() ([]models.InventoryItem, error)
	GetByID() (models.InventoryItem, error)
	Delete(models.InventoryItem) error
	Update(string, models.InventoryItem) error
}

type inventoryService struct {
	inventoryAccess dal.InventoryDalInterface
}

func NewInventoryService(inventoryDal dal.InventoryDalInterface) *inventoryService {
	return &inventoryService{inventoryAccess: inventoryDal}
}

func (s *inventoryService) Add(models.InventoryItem) error {
	return nil
}

func (s *inventoryService) Get() ([]models.InventoryItem, error) {
	return []models.InventoryItem{}, nil
}

func (s *inventoryService) GetByID() (models.InventoryItem, error) {
	return models.InventoryItem{}, nil
}

func (s *inventoryService) Delete(models.InventoryItem) error {
	return nil
}

func (s *inventoryService) Update(string, models.InventoryItem) error {
	return nil
}
