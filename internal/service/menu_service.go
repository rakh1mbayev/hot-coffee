package service

import (
	"fmt"
	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type MenuService interface {
	PostMenu(item model.MenuItem) error
	GetMenu() ([]model.MenuItem, error)
	GetMenuItemByID(id string) (*model.MenuItem, error)
	PutMenuItem(id string, item model.MenuItem) (*model.MenuItem, error)
	DeleteMenuItem(id string) error
}

type FileMenuService struct {
	dataAccess *dal.FileDataAccess
}

func NewFileMenuService(filePath string) *FileMenuService {
	return &FileMenuService{
		dataAccess: &dal.FileDataAccess{FilePath: filePath},
	}
}

func (f *FileMenuService) PostMenu(item model.MenuItem) error {
	items, err := f.dataAccess.LoadMenuItems()
	if err != nil {
		return err
	}
	items = append(items, item)
	return f.dataAccess.SaveMenuItems(items)
}

func (f *FileMenuService) GetMenu() ([]model.MenuItem, error) {
	return f.dataAccess.LoadMenuItems()
}

func (f *FileMenuService) GetMenuItemByID(id string) (*model.MenuItem, error) {
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

func (f *FileMenuService) PutMenuItem(id string, item model.MenuItem) (*model.MenuItem, error) {
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

func (f *FileMenuService) DeleteMenuItem(id string) error {
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
