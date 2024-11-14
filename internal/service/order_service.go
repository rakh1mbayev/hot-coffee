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

	ordersMenu, err := o.menuAccess.Get()
	if err != nil {
		return err
	}

	isExists := false
	for _, val := range ordersMenu {
		for _, vol := range order.Items {
			if val.ID == vol.ProductID {
				isExists = true
			}
		}
	}

	if !isExists {
		fmt.Println("Product with this name does not exists")
		model.Logger.Error("Product with this name does not exists")
		return nil
	}

	// Loop through order items and check menu
	for _, item := range order.Items {
		file, err := o.menuAccess.GetByID(item.ProductID)
		if err != nil {
			fmt.Printf("Product with ID %s not found in menu\n", item.ProductID)
			model.Logger.Info(fmt.Sprintf("Product with ID %s not found in inventory", item.ProductID))
			return fmt.Errorf("product with ID %s not found in inventory", item.ProductID)
		}

		for _, vel := range file.Ingredients {
			fileInventory, err := o.inventoryAcces.GetByID(vel.IngredientID)
			if err != nil {
				fmt.Printf("Inventory with ID %s not found in inventory\n", vel.IngredientID)
				model.Logger.Info(fmt.Sprintf("Inventory with ID %s not found in inventory", vel.IngredientID))
				return fmt.Errorf("inventory with ID %s not found in inventory", vel.IngredientID)
			}

			for _, vil := range order.Items {
				for _, vol := range ordersMenu {
					for _, vel := range vol.Ingredients {
						fmt.Println("fileInventory:", fileInventory)
						fmt.Println("fileInventory.Quantity:", fileInventory.Quantity)
						fmt.Println("float64(vil.Quantity)*vel.Quantity:", float64(vil.Quantity)*vel.Quantity)
						if fileInventory.Quantity < float64(vil.Quantity)*vel.Quantity {
							fmt.Println("Not enought item for this product")
							model.Logger.Info("Not enought item for this product")
							return fmt.Errorf("not enought item for this product")
						} else {
							break
						}
					}
				}
			}
		}
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
			// Loop through order items and check inventory
			for _, item := range order.Items {
				inventoryItem, err := o.inventoryAcces.GetByID(item.ProductID)
				if err != nil {
					fmt.Printf("Product with ID %s not found in inventory\n", item.ProductID)
					model.Logger.Info(fmt.Sprintf("Product with ID %s not found in inventory", item.ProductID))
					return fmt.Errorf("product with ID %s not found in inventory", item.ProductID)
				}

				newQuantity := int(inventoryItem.Quantity) - item.Quantity
				o.inventoryAcces.Update(item.ProductID, inventoryItem, newQuantity)
			}

			// If all items are in stock, proceed to close the order
			orders[i].Status = "closed"
			return o.dataAccess.Save(orders)
		}
	}

	// If no matching order was found
	fmt.Println("Order not found")
	model.Logger.Info("Order not found")
	return fmt.Errorf("order not found")
}
