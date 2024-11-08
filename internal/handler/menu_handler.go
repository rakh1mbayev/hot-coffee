package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	service "hot-coffee/internal/service"
	model "hot-coffee/models"
)

func PostMenu(w http.ResponseWriter, r *http.Request) {
	var putMenu model.MenuItem
	var put model.MenuItemIngredient
	var checkMenu []model.MenuItem

	if data, err := os.ReadFile(*model.Dir + "/menu.json"); err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, &checkMenu); err != nil {

			var checkMenuSingle model.MenuItem
			if err := json.Unmarshal(data, &checkMenuSingle); err != nil {
				// error
				fmt.Println("Error unmarshal file in: menu_handler.go -> PostMenu")
				return
			}

			checkMenu = append(checkMenu, checkMenuSingle)
		}
	}

	file, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading file in: menu_handler -> PostMenu")
		return
	}
	json.Unmarshal(file, &putMenu)
	json.Unmarshal(file, &put)


	service.PostMenuService(putMenu, put, checkMenu)
}

func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(*model.Dir + "/menu.json")
	if err != nil {
		// error
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		// error
	}
	service.GetMenuService(data, w)
}

func GetMenuID(w http.ResponseWriter, r *http.Request) {
}

func PutMenuID(w http.ResponseWriter, r *http.Request) {
}

func DeleteMenuID(w http.ResponseWriter, r *http.Request) {
}
