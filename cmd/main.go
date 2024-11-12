package main

import (
	"flag"
	"fmt"
	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
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
			flags.Logger.Error("Error creating menu.json in main.go -> main")
			fmt.Println("Error creating menu.json in main.go -> main:", err)
			return
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/inventory.json"); err != nil { // Check if inventory.json exists if not create
		file, err := os.Create(*flags.Dir + "/inventory.json")
		if err != nil {
			flags.Logger.Error("Error creating inventory.json in main.go -> main")
			fmt.Println("Error creating inventory.json in main.go -> main:", err)
			return
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/orders.json"); err != nil { // Check if orders.json exists if not create
		file, err := os.Create(*flags.Dir + "/orders.json")
		if err != nil {
			flags.Logger.Error("Error creating orders.json in main.go -> main")
			fmt.Println("Error creating orders.json in main.go -> main:", err)
			return
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/report.log"); err != nil { // Check if /report.log exists if not create
		file, err := os.Create(*flags.Dir + "/report.log")
		if err != nil {
			flags.Logger.Error("Error creating report.log in main.go -> main")
			fmt.Println("Error creating report.log in main.go -> main:", err)
			return
		}
		defer file.Close()
	}

	file, err := os.OpenFile(*flags.Dir+"/report.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		flags.Logger.Error("Error opening file in: main.go.go -> main")
		fmt.Println("Error opening file in: main.go.go -> main", err)
		return
	}

	flags.Logger = slog.New(slog.NewTextHandler(file, nil))

	mux := http.NewServeMux()

	ordersService := service.NewFileOrderService(*flags.Dir + "/orders.json")
	ordersHandler := handler.NewOrdersHandler(ordersService)

	// orders
	// need finish the last
	mux.HandleFunc("POST /orders", ordersHandler.Add)
	mux.HandleFunc("GET /orders", ordersHandler.Get)
	mux.HandleFunc("GET /orders/{id}", ordersHandler.GetByID)
	mux.HandleFunc("PUT /orders/{id}", ordersHandler.Update)
	mux.HandleFunc("DELETE /orders/{id}", ordersHandler.Delete)
	mux.HandleFunc("POST /orders/{id}/close", ordersHandler.Close)

	// menu
	// need finish first
	menuDal := dal.NewMenuRepo(*flags.Dir + "/menu.json")
	menuService := service.NewFileMenuService(menuDal)
	menuHandler := handler.NewMenuHandler(menuService)

	mux.HandleFunc("POST /menu", menuHandler.Add)
	mux.HandleFunc("GET /menu", menuHandler.Get)
	mux.HandleFunc("GET /menu/{id}", menuHandler.GetByID)
	mux.HandleFunc("PUT /menu/{id}", menuHandler.Update)
	mux.HandleFunc("DELETE /menu/{id}", menuHandler.Delete)

	// inventory
	// need finish second
	inventoryDal := dal.NewInventoryRepo(*flags.Dir + "/inventory.json")
	inventoryService := service.NewInventoryService(inventoryDal)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	mux.HandleFunc("POST /inventory", inventoryHandler.Add)
	mux.HandleFunc("GET /inventory", inventoryHandler.Get)
	mux.HandleFunc("GET /inventory/{id}", inventoryHandler.GetByID)
	mux.HandleFunc("PUT /inventory/{id}", inventoryHandler.Update)
	mux.HandleFunc("DELETE /inventory/{id}", inventoryHandler.Delete)

	// reports
	mux.HandleFunc("GET /reports/total-sales", handler.GetTotalSales)
	mux.HandleFunc("GET /reports/popular-items", handler.GetPopularItems)

	log.Fatal(http.ListenAndServe(*flags.Port, mux))

	// WE NEED INTERFACE
}
