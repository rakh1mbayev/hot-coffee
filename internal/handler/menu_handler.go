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

func (m *MenuHandler) Add(w http.ResponseWriter, r *http.Request) {
	var newMenuItem model.MenuItem
	json.NewDecoder(r.Body).Decode(&newMenuItem)
	if err := m.service.Add(newMenuItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		model.Logger.Error("Failed to add menu item")
		service.ErrorHandling("Failed to add menu item", w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (m *MenuHandler) Get(w http.ResponseWriter, r *http.Request) {
	items, err := m.service.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		model.Logger.Error("Failed to load menu")
		service.ErrorHandling("Failed to load menu", w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(items); err != nil {
		return
	}
}

func (m *MenuHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid request path")
		service.ErrorHandling("Invalid request path", w)
		return
	}
	item, err := m.service.GetByID(path[2])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		model.Logger.Error("Menu item not found")
		service.ErrorHandling("Menu item not found", w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(item); err != nil {
		return
	}
}

func (m *MenuHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updatedItem model.MenuItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Error reading request body")
		service.ErrorHandling("Error reading request body", w)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid JSON")
		service.ErrorHandling("Invalid JSON", w)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		model.Logger.Error("Invalid request path")
		service.ErrorHandling("Invalid request path", w)
		return
	}
	item, err := m.service.Update(path[2], updatedItem)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		model.Logger.Error("Menu item not found")
		service.ErrorHandling("Menu item not found", w)
		return
	}
}

func (m *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := m.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		model.Logger.Error("Menu item not found")
		service.ErrorHandling("Menu item not found", w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
