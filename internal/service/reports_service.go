package service

import (
	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type ReportsService interface {
	TotalPrice() (float64, error)
	PopularItems() (string, error)
}

type FileReportsService struct {
	dataAccess dal.ReportsDalInterface
}

func NewFileReportsService(dataAccess dal.ReportsDalInterface) *FileReportsService {
	return &FileReportsService{dataAccess: dataAccess}
}

func (f *FileReportsService) TotalPrice() (float64, error) {
	items, err := f.dataAccess.GetAll()
	if err != nil {
		return 0, err
	}
	var TotalPrice float64

	for _, order := range items {
		if order.Status != "closed" {
			continue
		}
		for _, items := range order.Items {
			TotalPrice += model.PriceList[items.ProductID] * float64(items.Quantity)
		}
	}
	return TotalPrice, nil
}

func (f *FileReportsService) PopularItems() (string, error) {
	items, err := f.dataAccess.GetAll()
	if err != nil {
		return "", err
	}
	var name string
	var maxPopularity int
	PopularItem := map[string]int{}

	for _, order := range items {
		if order.Status != "closed" {
			continue
		}
		for _, popularity := range order.Items {
			PopularItem[popularity.ProductID]++
			if maxPopularity < PopularItem[popularity.ProductID] {
				name = popularity.ProductID
			}
		}
	}
	return name, nil
}
