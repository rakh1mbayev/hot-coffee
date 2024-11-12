package service

import (
	"hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type OrdersService interface {
	PostOrders(item model.OrderItem) error
	GetOrders() ([]model.OrderItem, error)
	GetOrdersID(id string) (*model.OrderItem, error)
	PutOrdersID(id string, item model.OrderItem) (*model.OrderItem, error)
	DeleteOrdersID(id string) error
	PostOrdersIDnClose(id string) error
}

type FileOrderService struct {
	dataAccess *dal.OrderDataAccess
}

func NewFileOrderService(filePath string) *FileOrderService {
	return &FileOrderService{
		dataAccess: &dal.OrderDataAccess{FilePath: filePath},
	}
}

func (o *FileOrderService) PostOrders(item model.OrderItem) error {
	return nil
}

func (o *FileOrderService) GetOrders() ([]model.OrderItem, error) {
	return nil, nil
}

func (o *FileOrderService) GetOrdersID(id string) (*model.OrderItem, error) {
	return nil, nil
}

func (o *FileOrderService) PutOrdersID(id string, item model.OrderItem) (*model.OrderItem, error) {
	return nil, nil
}

func (o *FileOrderService) DeleteOrdersID(id string) error {
	return nil
}

func (o *FileOrderService) PostOrdersIDnClose(id string) error {
	return nil
}
