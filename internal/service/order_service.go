package service

import (
	"fmt"

	"hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type OrdersService interface {
	Add(item model.Order) error
	Get() ([]model.Order, error)
	GetByID(id string) (*model.Order, error)
	Update(id string, item model.Order) error
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
	if order.ID == "" {
		fmt.Println("Order ID can not be empty. Please write it!")
		model.Logger.Error("Order ID can not be empty. Please write it!")
		return nil
	}

	if order.CustomerName == "" {
		fmt.Println("Customer name can not be empty. Please write it!")
		model.Logger.Error("Customer name can not be empty. Please write it!")
		return nil
	}

	if order.CreatedAt == "" {
		fmt.Println("Create time can not be empty. Please write it!")
		model.Logger.Error("Create time can not be empty. Please write it!")
		return nil
	}

	if order.Status != "open" {
		fmt.Println("Please write correctly order status")
		model.Logger.Error("Create time can not be empty. Please write it!")
		return nil
	}

	items, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}
	for _, val := range items {
		if val.ID == order.ID {
			fmt.Println("Order Id can not be same")
			model.Logger.Error("Order Id can not be same")
			return nil
		}
	}

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

func (o *FileOrderService) GetByID(id string) (*model.Order, error) {
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		if order.ID == id {
			return &order, nil
		}
	}
	model.Logger.Info("order not found")
	return nil, fmt.Errorf("order not found")
}

func (o *FileOrderService) Update(id string, item model.Order) error {
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}
	for i, order := range orders {
		if order.ID == id {
			orders[i] = item
			return o.dataAccess.Save(orders)
		}
	}
	model.Logger.Info("menu item not found")
	return fmt.Errorf("menu item not found")
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
		model.Logger.Info("order not found")
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
	model.Logger.Info("menu item not found")
	return fmt.Errorf("menu item not found")
}
