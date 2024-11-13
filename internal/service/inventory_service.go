package service

import (
	"fmt"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryInterface interface {
	Add(models.InventoryItem) error
	Get() ([]models.InventoryItem, error)
	GetByID(id string) (models.InventoryItem, error)
	Update(string, models.InventoryItem) error
	Delete(string) error
}

type FileInventoryService struct {
	dataAccess dal.InventoryDalInterface
}

func NewInventoryService(inventoryDal dal.InventoryDalInterface) *FileInventoryService {
	return &FileInventoryService{dataAccess: inventoryDal}
}

func (s *FileInventoryService) Add(models.InventoryItem) error {
	return nil
}

func (s *FileInventoryService) Get() ([]models.InventoryItem, error) {
	return s.dataAccess.GetAll()
}

func (s *FileInventoryService) GetByID(id string) (models.InventoryItem, error) {
	items, err := s.dataAccess.GetAll()
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

func (s *FileInventoryService) Update(string, models.InventoryItem) error {
	return nil
}

func (s *FileInventoryService) Delete(id string) error {
	items, err := s.dataAccess.GetAll()
	if err != nil {
		return err
	}
	var updatedItems []models.InventoryItem
	for _, item := range items {
		if item.IngredientID != id {
			updatedItems = append(updatedItems, item)
		}
	}
	return s.dataAccess.Save(items)
}
