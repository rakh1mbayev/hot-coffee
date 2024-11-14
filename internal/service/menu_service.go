package service

import (
	"errors"
	"fmt"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type MenuService interface {
	Add(item model.MenuItem) error
	Get() ([]model.MenuItem, error)
	GetByID(id string) (*model.MenuItem, error)
	Update(id string, item model.MenuItem) error
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
		model.Logger.Error("Menu ID can not be empty")
		return errors.New("menu ID can not be empty")
	}

	if item.Name == "" {
		fmt.Println("Name can not be empty. Please write name")
		model.Logger.Error("Name can not be empty. Please write name")
		return errors.New("name can not be empty. Please write name")
	}

	if item.Price < 0 {
		fmt.Println("Price can not be lower than 0 (price >= 0)")
		model.Logger.Error("Price can not be lower than 0 (price >= 0)")
		return errors.New("pice can not be lower than 0 (price >= 0)")
	}

	for _, val := range item.Ingredients {
		if val.IngredientID == "" {
			fmt.Println("Ingredient ID can not be empty. Please write ingredient ID")
			model.Logger.Error("Ingredient ID can not be empty. Please write ingredient ID")
			return errors.New("ingredient ID can not be empty. Please write ingredient ID")
		}

		if val.Quantity <= 0 {
			fmt.Println("Quantity of ingredient can not be equal or lesser than 0 (quantity > 0)")
			model.Logger.Error("Quantity of ingredient can not be equal or lesser than 0 (quantity > 0)")
			return errors.New("quantity of ingredient can not be equal or lesser than 0 (quantity > 0)")
		}
	}

	items, err := f.dataAccess.GetAll()
	if err != nil {
		return err
	}

	for _, vol := range items {
		if vol.ID == item.ID {
			fmt.Println("Menu Id can not be same")
			model.Logger.Error("Menu Id can not be same")
			return errors.New("menu Id can not be same")
		}
	}

	items = append(items, item)
	return f.dataAccess.Save(items)
}

func (f *FileMenuService) Get() ([]model.MenuItem, error) {
	return f.dataAccess.GetAll()
}

func (f *FileMenuService) GetByID(id string) (*model.MenuItem, error) {
	items, err := f.dataAccess.GetAll()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.ID == id {
			return &item, nil
		}
	}
	model.Logger.Info("menu item not found")
	return nil, fmt.Errorf("menu item not found")
}

func (f *FileMenuService) Update(id string, item model.MenuItem) error {
	items, err := f.dataAccess.GetAll()
	if err != nil {
		return err
	}
	for i, existingItem := range items {
		if existingItem.ID == id {
			items[i] = item
			return f.dataAccess.Save(items)
		}
	}
	model.Logger.Info("menu item not found")
	return fmt.Errorf("menu item not found")
}

func (f *FileMenuService) Delete(id string) error {
	items, err := f.dataAccess.GetAll()
	if err != nil {
		return err
	}
	found := false
	var updatedItems []model.MenuItem
	for _, item := range items {
		if item.ID != id {
			updatedItems = append(updatedItems, item)
		} else {
			found = true
		}
	}
	if !found {
		model.Logger.Info("menu item not found")
		return fmt.Errorf("menu item not found")
	}
	return f.dataAccess.Save(updatedItems)
}
