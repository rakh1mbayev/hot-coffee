package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

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

func PutMenuItem(id string, item model.MenuItem) ([]byte, error) {
	file, err := os.ReadFile(*model.Dir + "/menu.json")
	if err != nil {
		return []byte{}, err
	}
	MenuItems := []model.MenuItem{}
	if err = json.Unmarshal(file, &MenuItems); err != nil {
		return []byte{}, err
	}
	found := false
	for i := 0; i < len(MenuItems); i++ {
		if MenuItems[i].ID == id {
			MenuItems[i] = item
			found = true
			break
		}
	}
	if !found {
		return []byte{}, errors.New("ERROR: didn't found any item with given id")
	}
	file, err = json.Marshal(MenuItems)
	if err != nil {
		return []byte{}, err
	}
	return file, nil
}

func DeleteMenuItem(id string, item model.MenuItem) ([]byte, error) {
	file, err := os.ReadFile(*model.Dir + "/menu.json")
	if err != nil {
		return []byte{}, err
	}
	MenuItems := []model.MenuItem{}
	if err = json.Unmarshal(file, &MenuItems); err != nil {
		return []byte{}, err
	}
	found := false
	NewMenuItems := []model.MenuItem{}
	for _, MenuItem := range MenuItems {
		if MenuItem.ID == id {
			found = true
			continue
		}
		NewMenuItems = append(NewMenuItems, MenuItem)
	}
	if !found {
		return []byte{}, errors.New("ERROR: didn't found any item with given id")
	}
	file, err = json.Marshal(NewMenuItems)
	if err != nil {
		return []byte{}, err
	}
	return file, nil
}
