package main

import (
	"flag"
	"fmt"
	"hot-coffee/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"

	handler "hot-coffee/internal/handler"

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

	// orders
	// need finish the last
	mux.HandleFunc("POST /orders", handler.PostOrders)
	mux.HandleFunc("GET /orders", handler.GetOrders)
	mux.HandleFunc("GET /orders/{id}", handler.GetOrdersID)
	mux.HandleFunc("PUT /orders/{id}", handler.PutOrdersID)
	mux.HandleFunc("DELETE /orders/{id}", handler.DeleteOrdersID)
	mux.HandleFunc("POST /orders/{id}/close", handler.PostOrdersIDClose)

	// menu
	// need finish first

	menuService := service.NewFileMenuService(*flags.Dir + "/menu.json")
	menuHandler := handler.NewMenuHandler(menuService)

	mux.HandleFunc("POST /menu", menuHandler.PostMenu)
	mux.HandleFunc("GET /menu", menuHandler.GetMenu)
	mux.HandleFunc("GET /menu/{id}", menuHandler.GetMenuItemByID)
	mux.HandleFunc("PUT /menu/{id}", menuHandler.PutMenuItem)
	mux.HandleFunc("DELETE /menu/{id}", menuHandler.DeleteMenuItem)

	// inventory
	// need finish second
	mux.HandleFunc("POST /inventory", handler.PostInventory)
	mux.HandleFunc("GET /inventory", handler.GetInventory)
	mux.HandleFunc("GET /inventory/{id}", handler.GetInventoryID)
	mux.HandleFunc("PUT /inventory/{id}", handler.PutInventoryID)
	mux.HandleFunc("DELETE /inventory/{id}", handler.DeleteInventoryID)

	// reports
	mux.HandleFunc("GET /reports/total-sales", handler.GetTotalSales)
	mux.HandleFunc("GET /reports/popular-items", handler.GetPopularItems)

	log.Fatal(http.ListenAndServe(":7070", mux))

	// WE NEED INTERFACE
}
