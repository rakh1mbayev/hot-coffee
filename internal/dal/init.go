package dal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	flags "hot-coffee/models"
)

func Create() {

	if _, err := os.Stat(*flags.Dir); err != nil { // Check if dir exists if not create
		os.Mkdir(*flags.Dir, 0o755)
	}

	if _, err := os.Stat(*flags.Dir + "/report.log"); err != nil {
		file, err := os.Create(*flags.Dir + "/report.log")
		if err != nil {
			fmt.Println("Error creating report.log", err)
			return
		}
		defer file.Close()
	}

	file, err := os.OpenFile(*flags.Dir+"/report.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println("Error opening report.log", err)
		return
	}

	flags.Logger = slog.New(slog.NewTextHandler(file, nil))

	if _, err := os.Stat(*flags.Dir + "/menu.json"); err != nil { // Check if /menu.json exists if not create
		file, err := os.Create(*flags.Dir + "/menu.json")
		if err != nil {
			fmt.Println("Error creating menu.json:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Println("Error writing '[]' in '/menu.json'", err)
			return
		}
	}

	if _, err := os.Stat(*flags.Dir + "/inventory.json"); err != nil { // Check if inventory.json exists if not create
		file, err := os.Create(*flags.Dir + "/inventory.json")
		if err != nil {
			fmt.Println("Error creating inventory.json", err)
		}
		defer file.Close()

		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Println("Error writing '[]' in '/inventory.json'", err)
			return
		}
	}

	if _, err := os.Stat(*flags.Dir + "/orders.json"); err != nil { // Check if orders.json exists if not create
		file, err := os.Create(*flags.Dir + "/orders.json")
		if err != nil {
			fmt.Println("Error creating orders.json", err)
		}
		defer file.Close()

		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Println("Error writing '[]' in '/orders.json'", err)
			return
		}
	}

	if _, err := os.Stat(*flags.Dir + "/report.log"); err != nil { // Check if /report.log exists if not create
		file, err := os.Create(*flags.Dir + "/report.log")
		if err != nil {
			fmt.Println("Error creating report.log in main.go -> main:", err)
		}
		defer file.Close()
	}
}

func Helper() {
	fmt.Println(
		`Coffee Shop Management System

Usage:
hot-coffee [--port <N>] [--dir <S>] 
hot-coffee --help

Options:
--help       Show this screen.
--port N     Port number.
--dir S      Path to the data directory.`)
}

func SendError(message string, status int, w http.ResponseWriter) {
	flags.Logger.Error(message)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // Set the status code in the response header

	// Create the error response using the struct
	errorResponse := flags.Error{
		Messege: message,
		Status:  int64(status),
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(errorResponse)
	if err != nil {
		// If JSON marshalling fails, send a generic internal server error response
		http.Error(w, `{"Error": "Internal Server Error", "Status": 500}`, http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonData)
}