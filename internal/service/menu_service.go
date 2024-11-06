package service

import (
	"fmt"
	"net/http"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

func PostMenuService(putMenu model.MenuItem, data []byte) {
	// logic

	// fmt.Println(string(data))

	if putMenu.ID == "" {
		model.Logger.Error("Menu ID can not be empty")
		return
	}

	if putMenu.ID == string(data) {
		fmt.Println("ID can not be same")
		return
	}

	fmt.Println(putMenu)
	dal.PostMenuDal(putMenu)
}

func GetMenuService(data []byte, w http.ResponseWriter) {
	// logic check

	dal.GetMenuDal(data, w)
}
