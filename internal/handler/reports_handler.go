package handler

import (
	"encoding/json"
	"net/http"

	"hot-coffee/internal/service"
)

type ReportsHandler struct {
	service service.ReportsService
}

func NewReportsHandler(service service.ReportsService) *ReportsHandler {
	return &ReportsHandler{service: service}
}

func (m *ReportsHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sales, err := m.service.TotalPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	maershal, err := json.Marshal(sales)
	if err != nil {
		http.Error(w, `{"Error": "Internal Server Error", "Status": 500}`, http.StatusInternalServerError)
	}
	w.Write(maershal)
}

func (m *ReportsHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	popularItem, err := m.service.PopularItems()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	maershal, err := json.Marshal(popularItem)
	if err != nil {
		http.Error(w, `{"Error": "Internal Server Error", "Status": 500}`, http.StatusInternalServerError)
	}
	w.Write(maershal)
}
