package service

import (
	"errors"
	"fmt"

	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type InventoryInterface interface {
	Add(models.InventoryItem) error
	Get() ([]models.InventoryItem, error)
	GetByID(id string) (models.InventoryItem, error)
	Update(string, models.InventoryItem, int) error
	Delete(string) error
}

type FileInventoryService struct {
	dataAccess dal.InventoryDalInterface
}

func NewInventoryService(inventoryDal dal.InventoryDalInterface) *FileInventoryService {
	return &FileInventoryService{dataAccess: inventoryDal}
}

func (s *FileInventoryService) Add(item models.InventoryItem) error {
	if item.IngredientID == "" {
		models.Logger.Error("Ingredient ID can not be empty")
		fmt.Println("Ingredient ID can not be empty")
		return errors.New("ingredient ID can not be empty")
	}

	if item.Name == "" {
		models.Logger.Error("Ingredient name can not be empty")
		fmt.Println("Ingredient name can not be empty")
		return errors.New("ingredient name can not be empty")
	}

	if item.Unit == "" {
		models.Logger.Error("Ingredient unit can not be empty")
		fmt.Println("Ingredient unit can not be empty")
		return errors.New("ingredient unit can not be empty")
	}

	if item.Quantity <= 0 {
		models.Logger.Error("Ingredient quantity can not be equal or lower than 0 (quantity > 0)")
		fmt.Println("Ingredient quantity can not be equal or lower than 0 (quantity > 0)")
		return errors.New("ingredient quantity can not be equal or lower than 0 (quantity > 0)")
	}

	items, err := s.dataAccess.GetAll()
	if err != nil {
		return errors.New("coul do not get data")
	}

	for _, val := range items {
		if val.IngredientID == item.IngredientID {
			models.Logger.Error("Ingredient ID can not be same")
			fmt.Println("Ingredient ID can not be same")
			return errors.New("ingredient ID can not be same")
		}
	}

	items = append(items, item)
	return s.dataAccess.Save(items)
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
	models.Logger.Info("menu item not found")
	return models.InventoryItem{}, fmt.Errorf("menu item not found")
}

func (s *FileInventoryService) Update(id string, item models.InventoryItem, newQuantity int) error {
	items, err := s.dataAccess.GetAll()
	if err != nil {
		return err
	}
	for i, existingItem := range items {
		if existingItem.IngredientID == id {
			items[i] = item
			items[i].Quantity = float64(newQuantity)
			return s.dataAccess.Save(items)
		}
	}
	models.Logger.Info("inventory item not found")
	return fmt.Errorf("inventory item not found")
}

func (s *FileInventoryService) Delete(id string) error {
	items, err := s.dataAccess.GetAll()
	if err != nil {
		return err
	}
	found := false
	var updatedItems []models.InventoryItem
	for _, item := range items {
		if item.IngredientID != id {
			updatedItems = append(updatedItems, item)
		} else {
			found = true
		}
	}
	if !found {
		models.Logger.Info("inventory item not found")
		return fmt.Errorf("inventory item not found")
	}
	return s.dataAccess.Save(updatedItems)
}
