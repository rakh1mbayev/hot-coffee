package service

import (
	"fmt"
	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type MenuService interface {
	Add(item model.MenuItem) error
	Get() ([]model.MenuItem, error)
	GetByID(id string) (*model.MenuItem, error)
	Update(id string, item model.MenuItem) (*model.MenuItem, error)
	Delete(id string) error
}

type FileMenuService struct {
	dataAccess dal.MenuDalInterface
}

func NewFileMenuService(dataAccess dal.MenuDalInterface) *FileMenuService {
	return &FileMenuService{dataAccess: dataAccess}
}

func (f *FileMenuService) Add(item model.MenuItem) error {
	if item.ID == "" {
		fmt.Println("Menu ID can not be empty")
		return nil
	}

	if item.Name == "" {
		fmt.Println("Name can not be empty. Please write name")
		return nil
	}

	if item.Price < 0 {
		fmt.Println("Price can not be lower than 0 (price >= 0)")
		return nil
	}

	for _, val := range item.Ingredients {
		if val.IngredientID == "" {
			fmt.Println("Ingredient ID can not be empty. Please write ingredient ID")
			return nil
		}

		if val.Quantity <= 0 {
			fmt.Println("Quantity of ingredient can not be equal or lesser than 0 (quantity > 0)")
			return nil
		}
	}

	items, err := f.dataAccess.LoadMenuItems()
	if err != nil {
		return err
	}

	for _, vol := range items {
		if vol.ID == item.ID {
			fmt.Println("Menu Id can not be same")
			return nil
		}
	}

	items = append(items, item)
	return f.dataAccess.SaveMenuItems(items)
}

func (f *FileMenuService) Get() ([]model.MenuItem, error) {
	return f.dataAccess.LoadMenuItems()
}

func (f *FileMenuService) GetByID(id string) (*model.MenuItem, error) {
	items, err := f.dataAccess.LoadMenuItems()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("menu item not found")
}

func (f *FileMenuService) Update(id string, item model.MenuItem) (*model.MenuItem, error) {
	items, err := f.dataAccess.LoadMenuItems()
	if err != nil {
		return nil, err
	}
	for i, existingItem := range items {
		if existingItem.ID == id {
			items[i] = item
			if err := f.dataAccess.SaveMenuItems(items); err != nil {
				return nil, err
			}
			return &item, nil
		}
	}
	return nil, fmt.Errorf("menu item not found")
}

func (f *FileMenuService) Delete(id string) error {
	items, err := f.dataAccess.LoadMenuItems()
	if err != nil {
		return err
	}
	var updatedItems []model.MenuItem
	for _, item := range items {
		if item.ID != id {
			updatedItems = append(updatedItems, item)
		}
	}
	return f.dataAccess.SaveMenuItems(updatedItems)
}
