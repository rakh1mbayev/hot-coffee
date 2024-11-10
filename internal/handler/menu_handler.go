package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	service "hot-coffee/internal/service"
	model "hot-coffee/models"
	dal "hot-coffee/internal/dal"
)

type Menu_handle interface {
	PostMenu(w http.ResponseWriter, r *http.Request)
}

type Menuhandler struct {
	Menu_serve service.Menu_serv
}

func NewMenuHandler(Menu_serve *service.Menu_serv) *Menuhandler {
	return &Menuhandler{Menu_serve: *Menu_serve}
}

func (m *Menuhandler) PostMenu(w http.ResponseWriter, r *http.Request) {
	var putMenu model.MenuItem
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
		fmt.Println("Error read body in: menu_handler.go -> PostMenu", err)
		return
	}

	if err := json.Unmarshal(file, &putMenu); err != nil {
		fmt.Println("Error unmarshaling putMenu in: menu_handler.go -> PostMenu")
		return
	}

	m.Menu_serve.PostMenuService(putMenu, checkMenu)
}

// GetMenuDal
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
	dal.GetMenuDal(data, w)
}

func GetMenuID(w http.ResponseWriter, r *http.Request) {
}

func PutMenuID(w http.ResponseWriter, r *http.Request) {
}

func DeleteMenuID(w http.ResponseWriter, r *http.Request) {
}
