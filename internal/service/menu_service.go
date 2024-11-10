package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type Serv_menu interface {
	PostMenuService(putMenu model.MenuItem, checkMenu []model.MenuItem)
}

type Menu_serv struct {
	menu_dal dal.Dal_menu
}

func NewDefault_servmenu(menu_dal dal.Dal_menu) *Menu_serv {
	return &Menu_serv{menu_dal: menu_dal}
}

func (s *Menu_serv) PostMenuService(putMenu model.MenuItem, checkMenu []model.MenuItem) {
	// logic

	if putMenu.ID == "" {
		fmt.Println("Menu ID can not be empty")
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

	for _, val := range putMenu.Ingredients {
		if val.IngredientID == "" {
			fmt.Println("Ingredient ID can not be empty. Please write ingredient ID")
			return
		}

		if val.Quantity <= 0 {
			fmt.Println("Quantity of ingredient can not be equal or lesser than 0 (quantity > 0)")
			return
		}
	}

	for _, vol := range checkMenu {
		if vol.ID == putMenu.ID {
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

func GetMenuItem(id string) (model.MenuItem, error) {
	file, err := os.ReadFile(*model.Dir + "/menu.json")
	if err != nil {
		// error
	}

	var menuItems []model.MenuItem
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		// error
	}
	for _, item := range menuItems {
		fmt.Println("ITEM", item.ID)
		if item.ID == id {
			return item, nil
		}
	}
	return model.MenuItem{}, errors.New("ERROR: didn't found any item with given id")
}
