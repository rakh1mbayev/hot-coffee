package service

import (
	"fmt"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryServiceInterface interface {
	Add(models.InventoryItem) error
	Get() ([]models.InventoryItem, error)
	GetByID(id string) (models.InventoryItem, error)
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
	return s.inventoryAccess.GetAll()
}

func (s *inventoryService) GetByID(id string) (models.InventoryItem, error) {
	items, err := s.inventoryAccess.GetAll()
	if err != nil {
		return models.InventoryItem{}, err
	}
	for _, item := range items {
		if item.IngredientID == id {
			return item, nil
		}
	}
	return models.InventoryItem{}, fmt.Errorf("menu item not found")
}

func (s *inventoryService) Delete(id string) error {
	items, err := s.inventoryAccess.GetAll()
	if err != nil {
		return err
	}
	var updatedItems []models.InventoryItem
	for _, item := range items {
		if item.IngredientID != id {
			updatedItems = append(updatedItems, item)
		}
	}
	return s.inventoryAccess.SaveInventoryItems(items)
}

func (s *inventoryService) Update(string, models.InventoryItem) error {
	return nil
}
