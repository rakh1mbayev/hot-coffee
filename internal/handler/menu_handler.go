package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"bufio"
	"os"
	"strings"

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
	file, err := os.Open(*model.Dir + "/menu.json")
	defer file.Close()
	if err != nil {
		slog.Error("Failed to open file "+*model.Dir+" /menu.json", err)
	}
	data, _ := io.ReadAll(file)
	service.GetMenuService(data, w)
}

func GetMenuID(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(*model.Dir + "/menu.json")
	defer file.Close()

	if err != nil {
		slog.Error("Failed to open file "+*model.Dir+" /menu.json", err)
	}

	path := strings.Split(r.URL.Path[1:], "/")
	id := path[1]
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	lastWord := "none"
	for scanner.Scan() {
		word := scanner.Text()
		if word == "\"product_id":""
	}
}

func PutMenuID(w http.ResponseWriter, r *http.Request) {
}

func DeleteMenuID(w http.ResponseWriter, r *http.Request) {
}
