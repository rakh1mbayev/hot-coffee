package service

import (
	"net/http"

	dal "hot-coffee/internal/dal"
	model "hot-coffee/models"
)

func PostMenuService(putMenu model.MenuItem) {
	// logic

	dal.PostMenuDal(putMenu)
}

func GetMenuService(data []byte, w http.ResponseWriter) {
	// logic check

	dal.GetMenuDal(data, w)
}
