package main

import (
	"flag"
	"fmt"
	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
	"hot-coffee/models"
	"log"
	"log/slog"
	"net/http"
	"os"

	"hot-coffee/internal/handler"

	flags "hot-coffee/models"
)

func main() {
	flag.Parse()

	if *flags.Help { // flag help I do not know where save it
		fmt.Println(
			`Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.`)
		return
	}

	if _, err := os.Stat(*flags.Dir); err != nil { // Check if dir exists if not create
		os.Mkdir(*flags.Dir, 0o755)
	}

	if _, err := os.Stat(*flags.Dir + "/menu.json"); err != nil { // Check if /menu.json exists if not create
		file, err := os.Create(*flags.Dir + "/menu.json")
		if err != nil {
			fmt.Println("Error creating menu.json in main.go -> main:", err)
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/inventory.json"); err != nil { // Check if inventory.json exists if not create
		file, err := os.Create(*flags.Dir + "/inventory.json")
		if err != nil {
			fmt.Println("Error creating inventory.json in main.go -> main:", err)
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/orders.json"); err != nil { // Check if orders.json exists if not create
		file, err := os.Create(*flags.Dir + "/orders.json")
		if err != nil {
			fmt.Println("Error creating orders.json in main.go -> main:", err)
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/report.log"); err != nil { // Check if /report.log exists if not create
		file, err := os.Create(*flags.Dir + "/report.log")
		if err != nil {
			fmt.Println("Error creating report.log in main.go -> main:", err)
		}
		defer file.Close()
	}

	file, err := os.OpenFile(*flags.Dir+"/report.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println("Error opening file in: main.go.go -> main", err)
		return
	}

	flags.Logger = slog.New(slog.NewTextHandler(file, nil))

	mux := http.NewServeMux()

	ordersService := service.NewFileOrderService(*flags.Dir + "/menu.json")
	ordersHandler := handler.NewOrdersHandler(ordersService)

	// orders
	// need finish the last
	mux.HandleFunc("POST /orders", ordersHandler.PostOrders)
	mux.HandleFunc("GET /orders", ordersHandler.GetOrders)
	mux.HandleFunc("GET /orders/{id}", ordersHandler.GetOrdersID)
	mux.HandleFunc("PUT /orders/{id}", ordersHandler.PutOrdersID)
	mux.HandleFunc("DELETE /orders/{id}", ordersHandler.DeleteOrdersID)
	mux.HandleFunc("POST /orders/{id}/close", ordersHandler.PostOrdersIDnClose)

	// menu
	// need finish first
	menuDal := dal.NewMenuRepo(*models.Dir + "/menu.json")
	menuService := service.NewFileMenuService(menuDal)
	menuHandler := handler.NewMenuHandler(menuService)

	mux.HandleFunc("POST /menu", menuHandler.PostMenu)
	mux.HandleFunc("GET /menu", menuHandler.GetMenu)
	mux.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItemByID)
	mux.HandleFunc("PUT /menu/{id}", menuHandler.PutMenuItem)
	mux.HandleFunc("DELETE /menu/{id}", menuHandler.DeleteMenuItem)

	// inventory
	// need finish second
	inventoryDal := dal.NewInventoryRepo(*models.Dir + "/inventory.json")
	inventoryService := service.NewInventoryService(inventoryDal)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	mux.HandleFunc("POST /inventory", inventoryHandler.PostInventory)
	mux.HandleFunc("GET /inventory", inventoryHandler.GetInventory)
	mux.HandleFunc("GET /inventory/{id}", inventoryHandler.GetInventoryID)
	mux.HandleFunc("PUT /inventory/{id}", inventoryHandler.PutInventoryID)
	mux.HandleFunc("DELETE /inventory/{id}", inventoryHandler.DeleteInventoryID)

	// reports
	mux.HandleFunc("GET /reports/total-sales", handler.GetTotalSales)
	mux.HandleFunc("GET /reports/popular-items", handler.GetPopularItems)

	log.Fatal(http.ListenAndServe(":7070", mux))

	// WE NEED INTERFACE
}
