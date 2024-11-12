package dal

import (
	"encoding/json"
	model "hot-coffee/models"
	"os"
)

type FileDataAccess struct {
	FilePath string
}

func (f *FileDataAccess) LoadMenuItems() ([]model.MenuItem, error) {
	file, err := os.ReadFile(f.FilePath)
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
	fileData, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return os.WriteFile(f.FilePath, fileData, 0644)
}
