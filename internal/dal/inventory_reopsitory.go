package dal

import (
	"encoding/json"
	"os"

	model "hot-coffee/models"
)

type InventoryDalInterface interface {
	GetAll() ([]model.InventoryItem, error)
	Save([]model.InventoryItem) error
}

type InventoryData struct {
	path string
}

func NewInventoryRepo(path string) *InventoryData {
	return &InventoryData{path: path}
}

func (i *InventoryData) GetAll() ([]model.InventoryItem, error) {
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

func (i *InventoryData) Save(items []model.InventoryItem) error {
	fileData, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(i.path, fileData, 0o644)
}
