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
	dataAccess     *dal.OrderData
	menuAccess     MenuService
	inventoryAcces InventoryInterface
}

func NewFileOrderService(filePath string, menuService MenuService, inventoryService InventoryInterface) *FileOrderService {
	return &FileOrderService{
		dataAccess:     &dal.OrderData{FilePath: filePath},
		menuAccess:     menuService,
		inventoryAcces: inventoryService,
	}
}

func (o *FileOrderService) Add(order model.Order) error {
	// 1. Validate required fields
	if order.ID == "" || order.CustomerName == "" || order.CreatedAt == "" || order.Status != "open" {
		model.Logger.Error("Invalid order data: missing fields or incorrect status")
		return fmt.Errorf("invalid order data")
	}

	// 2. Check if all items exist in the menu and validate inventory
	for _, orderItem := range order.Items {
		// Check if item exists in the menu
		menuItem, err := o.menuAccess.GetByID(orderItem.ProductID)
		if err != nil {
			model.Logger.Error(fmt.Sprintf("Product with ID %s not found in menu", orderItem.ProductID))
			return fmt.Errorf("product with ID %s not found in menu", orderItem.ProductID)
		}

		// Check each ingredient in the menu item for inventory availability
		for _, ingredient := range menuItem.Ingredients {
			inventoryItem, err := o.inventoryAcces.GetByID(ingredient.IngredientID)
			if err != nil {
				model.Logger.Error(fmt.Sprintf("Ingredient with ID %s not found in inventory", ingredient.IngredientID))
				return fmt.Errorf("ingredient with ID %s not found in inventory", ingredient.IngredientID)
			}

			// Check if enough quantity is available
			requiredQuantity := float64(orderItem.Quantity) * ingredient.Quantity
			if inventoryItem.Quantity < requiredQuantity {
				model.Logger.Error(fmt.Sprintf("Not enough stock for ingredient %s", ingredient.IngredientID))
				return fmt.Errorf("not enough stock for ingredient %s", ingredient.IngredientID)
			}
		}
	}

	// 3. Check for duplicate Order ID
	existingOrders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}
	for _, existingOrder := range existingOrders {
		if existingOrder.ID == order.ID {
			model.Logger.Error("Duplicate order ID")
			return fmt.Errorf("order ID already exists")
		}
	}

	// 4. Save the new order
	existingOrders = append(existingOrders, order)
	return o.dataAccess.Save(existingOrders)
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
	// 1. Retrieve all orders and check if the order exists
	orders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}

	var orderToClose *model.Order
	orderIndex := -1
	for i, order := range orders {
		if order.ID == id {
			orderToClose = &orders[i]
			orderIndex = i
			break
		}
	}
	if orderToClose == nil {
		model.Logger.Info("Order not found")
		return fmt.Errorf("order not found")
	}

	// 2. For each item in the order, verify against menu and update inventory
	for _, item := range orderToClose.Items {
		// Check if item exists in the menu
		menuItem, err := o.menuAccess.GetByID(item.ProductID)
		if err != nil {
			model.Logger.Error(fmt.Sprintf("Product with ID %s not found in menu", item.ProductID))
			return fmt.Errorf("product with ID %s not found in menu", item.ProductID)
		}

		// Check each ingredient in the menu item and update inventory
		for _, ingredient := range menuItem.Ingredients {
			inventoryItem, err := o.inventoryAcces.GetByID(ingredient.IngredientID)
			if err != nil {
				model.Logger.Error(fmt.Sprintf("Ingredient with ID %s not found in inventory", ingredient.IngredientID))
				return fmt.Errorf("ingredient with ID %s not found in inventory", ingredient.IngredientID)
			}

			// Calculate required quantity
			requiredQuantity := float64(item.Quantity) * ingredient.Quantity
			if inventoryItem.Quantity < requiredQuantity {
				model.Logger.Error(fmt.Sprintf("Not enough stock for ingredient %s", ingredient.IngredientID))
				return fmt.Errorf("not enough stock for ingredient %s", ingredient.IngredientID)
			}

			// Deduct the required quantity from inventory and update
			newQuantity := inventoryItem.Quantity - requiredQuantity
			err = o.inventoryAcces.Update(ingredient.IngredientID, inventoryItem, int(newQuantity))
			if err != nil {
				return fmt.Errorf("failed to update inventory for ingredient %s", ingredient.IngredientID)
			}
		}
	}

	// 3. Mark the order as closed and save it
	orders[orderIndex].Status = "closed"
	return o.dataAccess.Save(orders)
}
