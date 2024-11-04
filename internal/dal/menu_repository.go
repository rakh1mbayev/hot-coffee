package dal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"hot-coffee/models"
)

func PostMenuDal(putMenu models.MenuItem) {
	writeJson, _ := json.MarshalIndent(putMenu, "", "\t")
	err := os.WriteFile(*models.Dir+"/menu.json", writeJson, 0o644)
	if err != nil {
		log.Fatalf("Error writing json file in: menu_repository.go -> ServiceMenuPost %v", err)
	}
}

func GetMenuDal(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
