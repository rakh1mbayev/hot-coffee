package service

import (
	"fmt"
	"net/http"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

func PostMenuService(putMenu model.MenuItem, put model.MenuItemIngredient, checkMenu []model.MenuItem) {
	// logic

	if putMenu.ID == "" {
		model.Logger.Error("Menu ID can not be empty")
		return
	}

	if putMenu.Name == "" {
		fmt.Println("Name can not be empty. Please write name")
		return
	}

	if putMenu.Price < 0 {
		fmt.Println("Price can not be lower than 0 (price >= 0)")
		return
	}

	if put.IngredientID == "" {
		fmt.Println("Ingredient ID can not be empty. Please write ingredient ID name")
		return
	}
	
	if put.Quantity <= 0 {
		fmt.Println("Quantity of ingredient cao not be equal or lesser than 0 (quantity > 0)")
		return
	}

	for _, val := range checkMenu {
		if val.ID == putMenu.ID {
			fmt.Println("Id can not be same")
			return
		}
	}

	dal.PostMenuDal(putMenu)
}

func GetMenuService(data []byte, w http.ResponseWriter) {
	// logic check

	dal.GetMenuDal(data, w)
}
