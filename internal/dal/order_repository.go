package dal

import (
	"encoding/json"
	"os"

	model "hot-coffee/models"
)

type OrderDalInterface interface {
	GetAll() ([]model.Order, error)
	Save(items []model.Order) error
}

type OrderData struct {
	FilePath string
}

func (f *OrderData) GetAll() ([]model.Order, error) {
	file, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, err
	}
	var items []model.Order
	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (f *OrderData) Save(items []model.Order) error {
	fileData, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return os.WriteFile(f.FilePath, fileData, 0o644)
}
