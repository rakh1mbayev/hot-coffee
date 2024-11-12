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
	Delete(string) error
	Update(string, models.InventoryItem) (*models.InventoryItem ,error)
}

type inventoryService struct {
	inventoryAccess dal.InventoryDalInterface
}

func NewInventoryService(inventoryDal dal.InventoryDalInterface) *inventoryService {
	return &inventoryService{inventoryAccess: inventoryDal}
}

func (s *inventoryService) Add(newInventoryItem models.InventoryItem) error {
	if newInventoryItem.IngredientID == "" {
		// error
		fmt.Println("Ingredient ID can not be empty")
		return nil
	}

	if newInventoryItem.Name == "" {
		// error
		fmt.Println("Ingredient name can not be empty")
		return nil
	}

	if newInventoryItem.Unit == "" {
		fmt.Println("Ingredient unit can not be empty")
		return nil
	}

	if newInventoryItem.Quantity <= 0 {
		fmt.Println("Ingredient quantity can not be equal or lower than 0 (quantity > 0)")
		return nil
	}

	items, err := s.inventoryAccess.GetAll()
	if err != nil {
		// error
		return err
	}

	for _, val := range items {
		if val.IngredientID == newInventoryItem.IngredientID {
			// error
			fmt.Println("Ingredient ID can not be same")
			return nil
		}
	}

	items = append(items, newInventoryItem)
	return s.inventoryAccess.SaveInventoryItems(items)
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

func (s *inventoryService) Update(id string, item models.InventoryItem) (*models.InventoryItem, error) {
	items, err := s.inventoryAccess.GetAll()
	if err != nil {
		return nil, err
	}
	for i, existingItem := range items {
		if existingItem.IngredientID == id {
			items[i] = item
			if err := s.inventoryAccess.SaveInventoryItems(items); err != nil {
				return nil, err
			}
			return &item, nil
		}
	}
	return nil, fmt.Errorf("menu item not found")
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
	return s.inventoryAccess.SaveInventoryItems(updatedItems)
}
