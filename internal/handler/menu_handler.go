package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	model "hot-coffee/models"
)

// model

// MenuService Interface (SRP, OCP)
type MenuService interface {
	PostMenu(item model.MenuItem, existingItems []model.MenuItem) error
	GetMenu() ([]model.MenuItem, error)
	GetMenuItemByID(id string) (*model.MenuItem, error)
	PutMenuItem(id string, item model.MenuItem) (*model.MenuItem, error)
	DeleteMenuItem(id string) error
}

// Service Implementation
type FileMenuService struct {
	filePath string
}

func (f *FileMenuService) PostMenu(item model.MenuItem, existingItems []model.MenuItem) error {
	existingItems = append(existingItems, item)
	return f.saveMenuItems(existingItems)
}

func (f *FileMenuService) GetMenu() ([]model.MenuItem, error) {
	return f.loadMenuItems()
}

func (f *FileMenuService) GetMenuItemByID(id string) (*model.MenuItem, error) {
	items, err := f.loadMenuItems()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("menu item not found")
}

func (f *FileMenuService) PutMenuItem(id string, item model.MenuItem) (*model.MenuItem, error) {
	items, err := f.loadMenuItems()
	if err != nil {
		return nil, err
	}
	for i, existingItem := range items {
		if existingItem.ID == id {
			items[i] = item
			if err := f.saveMenuItems(items); err != nil {
				return nil, err
			}
			return &item, nil
		}
	}
	return nil, fmt.Errorf("menu item not found")
}

<<<<<<< HEAD
func (f *FileMenuService) DeleteMenuItem(id string) error {
	items, err := f.loadMenuItems()
=======
// Erganat

func GetMenuHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(*model.Dir + "/menu.json")
>>>>>>> d7d1c89df6ca3c2baddc084f6e6f330e4ae84b57
	if err != nil {
		return err
	}
	var updatedItems []model.MenuItem
	for _, item := range items {
		if item.ID != id {
			updatedItems = append(updatedItems, item)
		}
	}
	return f.saveMenuItems(updatedItems)
}

// Helper functions for file I/O
func (f *FileMenuService) loadMenuItems() ([]model.MenuItem, error) {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, err
	}
	var items []model.MenuItem
	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (f *FileMenuService) saveMenuItems(items []model.MenuItem) error {
	fileData, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, fileData, 0644)
}

// MenuHandler (SRP, DIP)
type MenuHandler struct {
	service MenuService
}

func NewMenuHandler(service MenuService) *MenuHandler {
	return &MenuHandler{service: service}
}

func (m *MenuHandler) PostMenu(w http.ResponseWriter, r *http.Request) {
	var newMenuItem model.MenuItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &newMenuItem); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	existingItems, _ := m.service.GetMenu() // Can handle error if necessary
	if err := m.service.PostMenu(newMenuItem, existingItems); err != nil {
		http.Error(w, "Failed to add menu item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get Menu Handler (SRP)
func (m *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	items, err := m.service.GetMenu()
	if err != nil {
		http.Error(w, "Failed to load menu", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

// Get Menu Item by ID Handler (SRP)
func (m *MenuHandler) GetMenuItemByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	item, err := m.service.GetMenuItemByID(path[2])
	if err != nil {
		http.Error(w, "Menu item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Put Menu Item by ID Handler (SRP)
func (m *MenuHandler) PutMenuItem(w http.ResponseWriter, r *http.Request) {
	var updatedItem model.MenuItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}

	item, err := m.service.PutMenuItem(path[2], updatedItem)
	if err != nil {
		http.Error(w, "Menu item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

// Delete Menu Item by ID Handler (SRP)
func (m *MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}

	err := m.service.DeleteMenuItem(path[2])
	if err != nil {
		http.Error(w, "Menu item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
