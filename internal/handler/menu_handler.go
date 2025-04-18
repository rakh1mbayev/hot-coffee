package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/service"
	model "hot-coffee/models"
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
		dal.SendError("Failed to add menu item", http.StatusInternalServerError, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (m *MenuHandler) Get(w http.ResponseWriter, r *http.Request) {
	items, err := m.service.Get()
	if err != nil {
		dal.SendError("Failed to load menu", http.StatusInternalServerError, w)
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
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	item, err := m.service.GetByID(path[2])
	if err != nil {
		dal.SendError("Menu item not found", http.StatusNotFound, w)
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
		dal.SendError("Error reading request body", http.StatusBadRequest, w)
		return
	}
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		dal.SendError("Invalid JSON", http.StatusBadRequest, w)
		return
	}
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		dal.SendError("Invalid request path", http.StatusBadRequest, w)
		return
	}
	if err := m.service.Update(path[2], updatedItem); err != nil {
		dal.SendError("Menu item not found", http.StatusNotFound, w)
		return
	}
}

func (m *MenuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := m.service.Delete(id)
	if err != nil {
		dal.SendError("Menu item not found", http.StatusNotFound, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
