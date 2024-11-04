package service

import (
	"fmt"
	"net/http"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

func PostMenuService(putMenu model.MenuItem) {
	// logic

	fmt.Println(putMenu)
	dal.PostMenuDal(putMenu)
}

func GetMenuService(data []byte, w http.ResponseWriter) {
	// logic check

	dal.GetMenuDal(data, w)
}
