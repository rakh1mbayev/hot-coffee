package service

import (
	"errors"
	"fmt"
	"time"

	"hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type OrdersService interface {
	Add(item model.CreateOrderRequest) error
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

func (o *FileOrderService) Add(orderCheck model.CreateOrderRequest) error {
	// 1. Validate required fields
	if orderCheck.CustomerName == "" {
		model.Logger.Error("Invalid order data: missing fields")
		return fmt.Errorf("invalid order data")
	}

	// 2. Check if all items exist in the menu and validate inventory
	for _, orderItem := range orderCheck.Items {
		if orderItem.Quantity <= 0 {
			fmt.Println("Quantity can not be less or equal to 0 (quantity > 0)")
			model.Logger.Error("Quantity can not be less or equal to 0 (quantity > 0)")
			return fmt.Errorf("quantity can not be less or equal to 0")
		}

		// Check if item exists in the menu
		menuItem, err := o.menuAccess.GetByID(orderItem.ProductID)
		if err != nil {
			fmt.Printf("Product with ID %s not found in menu\n", orderItem.ProductID)
			model.Logger.Error(fmt.Sprintf("Product with ID %s not found in menu", orderItem.ProductID))
			return fmt.Errorf("product with ID %s not found in menu", orderItem.ProductID)
		}

		// Check each ingredient in the menu item for inventory availability
		for _, ingredient := range menuItem.Ingredients {
			inventoryItem, err := o.inventoryAcces.GetByID(ingredient.IngredientID)
			if err != nil {
				fmt.Printf("Ingredient with ID %s not found in inventory\n", ingredient.IngredientID)
				model.Logger.Error(fmt.Sprintf("Ingredient with ID %s not found in inventory", ingredient.IngredientID))
				return fmt.Errorf("ingredient with ID %s not found in inventory", ingredient.IngredientID)
			}

			// Check if enough quantity is available
			requiredQuantity := float64(orderItem.Quantity) * ingredient.Quantity
			if inventoryItem.Quantity < requiredQuantity {
				fmt.Printf("Not enough stock for ingredient %s\n", ingredient.IngredientID)
				model.Logger.Error(fmt.Sprintf("Not enough stock for ingredient %s", ingredient.IngredientID))
				return fmt.Errorf("not enough stock for ingredient %s", ingredient.IngredientID)
			}
		}
	}

	// 3. Check for duplicate Order ID
	generateID := generateOrderID()
	existingOrders, err := o.dataAccess.GetAll()
	if err != nil {
		return err
	}

	for {
		isUnique := true
		for _, existingOrder := range existingOrders {
			if existingOrder.ID == generateID {
				generateID = generateOrderID()
				isUnique = false
				break
			}
		}
		if isUnique {
			break // Exit outer loop when ID is unique
		}
	}

	// 4. Save the new order
	order := model.Order{
		ID:           generateID,
		CustomerName: orderCheck.CustomerName,
		Items:        orderCheck.Items,
		Status:       "open",
		CreatedAt:    time.Now().Format(time.RFC3339),
	}
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
	fmt.Println("order not found")
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
	fmt.Println("menu item not found")
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
		fmt.Println("order not found")
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
			if order.Status == "closed" {
				model.Logger.Info("already closed")
				return errors.New("already closed")
			}
			orderToClose = &orders[i]
			orderIndex = i
			break
		}
	}
	if orderToClose == nil {
		fmt.Println("order not found")
		model.Logger.Info("Order not found")
		return fmt.Errorf("order not found")
	}

	// 2. For each item in the order, verify against menu and update inventory
	for _, item := range orderToClose.Items {
		// Check if item exists in the menu
		menuItem, err := o.menuAccess.GetByID(item.ProductID)
		model.PriceList[menuItem.ID] = menuItem.Price
		if err != nil {
			fmt.Printf("Product with ID %s not found in menu", item.ProductID)
			model.Logger.Error(fmt.Sprintf("Product with ID %s not found in menu", item.ProductID))
			return fmt.Errorf("product with ID %s not found in menu", item.ProductID)
		}

		// Check each ingredient in the menu item and update inventory
		for _, ingredient := range menuItem.Ingredients {
			inventoryItem, err := o.inventoryAcces.GetByID(ingredient.IngredientID)
			if err != nil {
				fmt.Printf("Ingredient with ID %s not found in inventory", ingredient.IngredientID)
				model.Logger.Error(fmt.Sprintf("Ingredient with ID %s not found in inventory", ingredient.IngredientID))
				return fmt.Errorf("ingredient with ID %s not found in inventory", ingredient.IngredientID)
			}

			// Calculate required quantity
			requiredQuantity := float64(item.Quantity) * ingredient.Quantity
			if inventoryItem.Quantity < requiredQuantity {
				fmt.Printf("Not enough stock for ingredient %s", ingredient.IngredientID)
				model.Logger.Error(fmt.Sprintf("Not enough stock for ingredient %s", ingredient.IngredientID))
				return fmt.Errorf("not enough stock for ingredient %s", ingredient.IngredientID)
			}

			// Deduct the required quantity from inventory and update
			newQuantity := inventoryItem.Quantity - requiredQuantity
			err = o.inventoryAcces.Update(ingredient.IngredientID, inventoryItem, int(newQuantity))
			if err != nil {
				fmt.Printf("failed to update inventory for ingredient %s", ingredient.IngredientID)
				return fmt.Errorf("failed to update inventory for ingredient %s", ingredient.IngredientID)
			}
		}
	}

	// 3. Mark the order as closed and save it
	orders[orderIndex].Status = "closed"

	return o.dataAccess.Save(orders)
}

func generateOrderID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
