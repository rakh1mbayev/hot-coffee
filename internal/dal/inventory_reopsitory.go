package dal

import (
	"encoding/json"
	model "hot-coffee/models"
	"os"
)

type InventoryDalInterface interface {
	GetAll() ([]model.InventoryItem, error)
	SaveInventoryItems([]model.InventoryItem) error
}

type InventoryFileDataAccess struct {
	path string
}

func NewInventoryRepo(path string) *InventoryFileDataAccess {
	return &InventoryFileDataAccess{path: path}
}

func (i *InventoryFileDataAccess) GetAll() ([]model.InventoryItem, error) {
	file, err := os.ReadFile(i.path)
	if err != nil {
		return nil, err
	}
	var items []model.InventoryItem
	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (i *InventoryFileDataAccess) SaveInventoryItems(items []model.InventoryItem) error {
	fileData, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(i.path, fileData, 0644)
}
