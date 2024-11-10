package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	service "hot-coffee/internal/service"
	model "hot-coffee/models"
)

type Menu_handle interface {
	PostMenu(w http.ResponseWriter, r *http.Request)
}

type Menuhandler struct {
	Menu_serve service.Menu_serv
}

func PostMenu(w http.ResponseWriter, r *http.Request) {
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

	service.PostMenuService(putMenu, checkMenu)
}

// Erganat

func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(*model.Dir + "/menu.json")
	if err != nil {
		// error
	}
	defer file.Close()
	if err != nil {
		slog.Error("Failed to open file "+*model.Dir+" /menu.json", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		// error
	}
	service.GetMenuService(data, w)
}

type Getter interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type GetMenuID struct{}

func (m GetMenuID) Get(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile(*model.Dir + "/menu.json")
	if err != nil {
		// error
	}
	path := strings.Split(r.URL.Path[1:], "/")
	id := path[1]
	fmt.Println("ENTERED")

	var menuItems []model.MenuItem
	err = json.Unmarshal(file, &menuItems)
	if err != nil {
		// error
	}
	fmt.Println(id)
	for _, item := range menuItems {
		fmt.Println("ITEM", item.ID)
		if item.ID == id {
			result, err := json.Marshal(item)
			if err != nil {
				// error
			}
			service.GetMenuService(result, w)
			return
		}
	}
	fmt.Println("DOES NOT EXIST")
	// item does not existf
}

type PutMenuID struct{}

func (m PutMenuID) Put(w http.ResponseWriter, r *http.Request) {
}

type DeleteMenuID struct{}

func (m DeleteMenuID) Delete(w http.ResponseWriter, r *http.Request) {
}
