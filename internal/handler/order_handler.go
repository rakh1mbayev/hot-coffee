package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	service "hot-coffee/internal/service"
	models "hot-coffee/models"
)

func PostOrders(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	file, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reding file in: order_handler.go -> PostOrders")
	}

	json.Unmarshal(file, &order)
	service.PostOrder(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open(*models.Dir + "/files.json")
	defer file.Close()
	data, _ := io.ReadAll(file)
	service.GetOrder(data, w)
}

func GetOrdersID(w http.ResponseWriter, r *http.Request) {
}

func PutOrdersID(w http.ResponseWriter, r *http.Request) {
}

func DeleteOrdersID(w http.ResponseWriter, r *http.Request) {
}

func PostOrdersIDClose(w http.ResponseWriter, r *http.Request) {
}
