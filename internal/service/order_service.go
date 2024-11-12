package service

import (
	"fmt"
	"hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type OrdersService interface {
	Add(item model.OrderItem) error
	Get() ([]model.OrderItem, error)
	GetByID(id string) (*model.OrderItem, error)
	Update(id string, item model.OrderItem) error
	Delete(id string) error
	Close(id string) error
}

type FileOrderService struct {
	dataAccess *dal.OrderData
}

func NewFileOrderService(filePath string) *FileOrderService {
	return &FileOrderService{
		dataAccess: &dal.OrderData{FilePath: filePath},
	}
}

func (o *FileOrderService) Add(order model.Order) error {
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}
	orders = append(orders, order)
	return o.dataAccess.Save(orders)
}

func (o *FileOrderService) Get() ([]model.Order, error) {
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *FileOrderService) GetByID(id string) (*model.OrderItem, error) {
	return nil, nil
}

func (o *FileOrderService) Update(id string, item model.OrderItem) error {
	return nil
}

func (o *FileOrderService) Delete(id string) error {
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}
	found := false
	var newOrders []model.Order
	for _, order := range orders {
		if order.ID != id {
			newOrders = append(newOrders, order)
		} else {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("order not found")
	}
	return o.dataAccess.Save(newOrders)
}

func (o *FileOrderService) Close(id string) error {
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}
	for i, order := range orders {
		if order.ID == id {
			orders[i].Status = "closed"
			return o.dataAccess.Save(orders)
		}
	}
	return fmt.Errorf("menu item not found")
}
