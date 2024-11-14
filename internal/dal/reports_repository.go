package dal

import (
	"encoding/json"
	"os"

	model "hot-coffee/models"
)

type ReportsDalInterface interface {
	GetAll() ([]model.Order, error)
}

type ReportsData struct {
	filePath string
}

func NewReportsRepo(filePath string) *ReportsData {
	return &ReportsData{filePath: filePath}
}

func (f *ReportsData) GetAll() ([]model.Order, error) {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, err
	}
	var items []model.Order
	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}
	return items, nil
}
