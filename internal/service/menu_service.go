package service

import (
	"fmt"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

type Serv_menu interface {
	PostMenuService(putMenu model.MenuItem, checkMenu []model.MenuItem)
}
type Menu_serv struct {
	menu_dal dal.Menu_dal
}

func NewDefault_servmenu(menu_dal *dal.Menu_dal) *Menu_serv {
	return &Menu_serv{menu_dal: *menu_dal}
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

	s.menu_dal.PostMenuDal(putMenu)
}

