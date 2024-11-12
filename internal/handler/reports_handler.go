package handler

import (
	"hot-coffee/internal/service"
	"net/http"
)

type ReportsHandler struct {
	reports service.MenuService
}

func (r *ReportsHandler) GetTotalSales(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	sales, err := r.reports.TotalPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(sales)
}

func (r *ReportsHandler) GetPopularItems(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	popularItem, err := r.reports.PopularItems()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(popularItem))
}
