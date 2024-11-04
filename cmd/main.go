package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	handler "hot-coffee/internal/handler"
	flags "hot-coffee/models"
)

func main() {
	flag.Parse()

	if *flags.Help { // flag help I do not know where save it
	}

	if _, err := os.Stat(*flags.Dir); err != nil {
		os.Mkdir(*flags.Dir, 0o755)
	}

	if _, err := os.Stat(*flags.Dir + "/menu.json"); err != nil {
		file, err := os.Create(*flags.Dir + "/menu.json")
		if err != nil {
			fmt.Println("Error creating menu.json in main.go -> main:", err)
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/inventory.json"); err != nil {
		file, err := os.Create(*flags.Dir + "/inventory.json")
		if err != nil {
			fmt.Println("Error creating inventory.json in main.go -> main:", err)
		}
		defer file.Close()
	}

	if _, err := os.Stat(*flags.Dir + "/orders.json"); err != nil {
		file, err := os.Create(*flags.Dir + "/orders.json")
		if err != nil {
			fmt.Println("Error creating orders.json in main.go -> main:", err)
		}
		defer file.Close()
	}

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
	mux.HandleFunc("POST /menu", handler.PostMenu)
	mux.HandleFunc("GET /menu", handler.GetMenuHandler)
	mux.HandleFunc("GET /menu/{id}", handler.GetMenuID)
	mux.HandleFunc("PUT /menu/{id}", handler.PutMenuID)
	mux.HandleFunc("DELETE /menu/{id}", handler.DeleteMenuID)

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
}
