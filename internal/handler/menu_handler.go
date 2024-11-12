package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	model "hot-coffee/models"
	"io"
	"net/http"
	"strings"
)

type MenuHandler struct {
	service service.MenuService
}

func NewMenuHandler(service service.MenuService) *MenuHandler {
	return &MenuHandler{service: service}
}

func (m *MenuHandler) PostMenu(w http.ResponseWriter, r *http.Request) {
	var newMenuItem model.MenuItem
	json.NewDecoder(r.Body).Decode(&newMenuItem)
	if err := m.service.PostMenu(newMenuItem); err != nil {
		http.Error(w, "Failed to add menu item", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (m *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	items, err := m.service.GetMenu()
	if err != nil {
		http.Error(w, "Failed to load menu", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(items)
}

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
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(item)
}

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
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (m *MenuHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := m.service.DeleteMenuItem(id)
	if err != nil {
		http.Error(w, "Menu item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
