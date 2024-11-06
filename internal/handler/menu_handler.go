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

	reader, err := os.Open(*model.Dir + "/menu.json")
	if err != nil {
		// error
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		// error
	}

	file, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading file in: menu_handler -> PostMenu")
		return
	}
	json.Unmarshal(file, &putMenu)

	service.PostMenuService(putMenu, data)
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
