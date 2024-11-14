package main

import (
	"flag"
	"log"
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"

	"hot-coffee/internal/handler"

	flags "hot-coffee/models"
)

func main() {
	flag.Parse()

	if *flags.Help {
		dal.Helper()
		return
	}

	dal.Create()

	mux := http.NewServeMux()

	ordersService := service.NewFileOrderService(*flags.Dir + "/orders.json")
	ordersHandler := handler.NewOrdersHandler(ordersService)

	// orders
	mux.HandleFunc("POST /orders", ordersHandler.Add)
	mux.HandleFunc("GET /orders", ordersHandler.Get)
	mux.HandleFunc("GET /orders/{id}", ordersHandler.GetByID)
	mux.HandleFunc("PUT /orders/{id}", ordersHandler.Update)
	mux.HandleFunc("DELETE /orders/{id}", ordersHandler.Delete)
	mux.HandleFunc("POST /orders/{id}/close", ordersHandler.Close)

	// menu
	menuDal := dal.NewMenuRepo(*flags.Dir + "/menu.json")
	menuService := service.NewFileMenuService(menuDal)
	menuHandler := handler.NewMenuHandler(menuService)

	mux.HandleFunc("POST /menu", menuHandler.Add)
	mux.HandleFunc("GET /menu", menuHandler.Get)
	mux.HandleFunc("GET /menu/{id}", menuHandler.GetByID)
	mux.HandleFunc("PUT /menu/{id}", menuHandler.Update)
	mux.HandleFunc("DELETE /menu/{id}", menuHandler.Delete)

	// inventory
	inventoryDal := dal.NewInventoryRepo(*flags.Dir + "/inventory.json")
	inventoryService := service.NewInventoryService(inventoryDal)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	mux.HandleFunc("POST /inventory", inventoryHandler.Add)
	mux.HandleFunc("GET /inventory", inventoryHandler.Get)
	mux.HandleFunc("GET /inventory/{id}", inventoryHandler.GetByID)
	mux.HandleFunc("PUT /inventory/{id}", inventoryHandler.Update)
	mux.HandleFunc("DELETE /inventory/{id}", inventoryHandler.Delete)

	reportsDal := dal.NewMenuRepo(*flags.Dir + "/orders.json")
	reportsService := service.NewFileMenuService(reportsDal)
	reportsHandler := handler.NewMenuHandler(reportsService)

	// reports
	mux.HandleFunc("GET /reports/total-sales", reportsHandler.GetTotalSales)
	mux.HandleFunc("GET /reports/popular-items", reportsHandler.GetPopularItems)

	flags.Logger.Info("Host is started")
	log.Fatal(http.ListenAndServe(":"+*flags.Port, mux))
}
