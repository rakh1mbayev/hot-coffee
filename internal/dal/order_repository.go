package dal

import (
	"encoding/json"
	model "hot-coffee/models"
	"os"
)

type OrderDataAccess struct {
	FilePath string
}

func (f *OrderDataAccess) LoadOrderItems() ([]model.OrderItem, error) {
	file, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, err
	}
	var items []model.OrderItem
	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (f *OrderDataAccess) SaveOrderItems(items []model.OrderItem) error {
	fileData, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return os.WriteFile(f.FilePath, fileData, 0644)
}
