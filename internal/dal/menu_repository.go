package dal

import (
	"encoding/json"
	model "hot-coffee/models"
	"os"
)

type MenuDalInterface interface {
	LoadMenuItems() ([]model.MenuItem, error)
	SaveMenuItems(items []model.MenuItem) error
}

type FileDataAccess struct {
	filePath string
}

func NewMenuRepo(filePath string) *FileDataAccess {
	return &FileDataAccess{filePath: filePath}
}

func (f *FileDataAccess) LoadMenuItems() ([]model.MenuItem, error) {
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

func (f *FileDataAccess) SaveMenuItems(items []model.MenuItem) error {
	fileData, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(f.filePath, fileData, 0644)
}
