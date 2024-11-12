package dal

import (
	"encoding/json"
	model "hot-coffee/models"
	"os"
)

type MenuDalInterface interface {
	GetAll() ([]model.MenuItem, error)
	Save(items []model.MenuItem) error
}

type MenuData struct {
	filePath string
}

func NewMenuRepo(filePath string) *MenuData {
	return &MenuData{filePath: filePath}
}

func (f *MenuData) GetAll() ([]model.MenuItem, error) {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, err
	}
	var items []model.MenuItem
	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (f *MenuData) Save(items []model.MenuItem) error {
	fileData, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(f.filePath, fileData, 0644)
}
