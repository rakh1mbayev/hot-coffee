package handler

import (
	"net/http"
)

func (m *MenuHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sales, err := m.service.TotalPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(sales)
}

func (m *MenuHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	popularItem, err := m.service.PopularItems()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(popularItem))
}
