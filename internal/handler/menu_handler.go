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

	file, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading file in: menu_handler -> PostMenu", err)
		return
	}
	json.Unmarshal(file, &putMenu)

	service.PostMenuService(putMenu)
}

func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(*model.Dir + "/menu.json")
	defer file.Close()
	data, _ := io.ReadAll(file)
	service.GetMenuService(data, w)
}

func GetMenuID(w http.ResponseWriter, r *http.Request) {
}

func PutMenuID(w http.ResponseWriter, r *http.Request) {
}

func DeleteMenuID(w http.ResponseWriter, r *http.Request) {
}
