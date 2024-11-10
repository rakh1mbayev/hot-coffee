package dal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"hot-coffee/models"
)

type Dal_menu interface{
	PostMenuDal(putMenu models.MenuItem)
}

type Menu_dal struct{
	Datapath string
}

func NewDefault(datapath string)*Menu_dal{
	return &Menu_dal{Datapath: datapath}
}

func (dal *Menu_dal)PostMenuDal(putMenu models.MenuItem) {
	var menuItems []models.MenuItem

	if file, err := os.ReadFile(*models.Dir + "/menu.json"); err == nil && len(file) > 0 {
		if err := json.Unmarshal(file, &menuItems); err != nil {

			var singleItem models.MenuItem
			if err := json.Unmarshal(file, &singleItem); err != nil {
				log.Fatalf("Error unmarshaling JSON file in: menu_repository.go -> PostMenuDal %v", err)
			}

			menuItems = append(menuItems, singleItem)
		}
	}

	menuItems = append(menuItems, putMenu)

	writeJson, _ := json.MarshalIndent(menuItems, "", "\t")
	err := os.WriteFile(*models.Dir+"/menu.json", writeJson, 0o644)
	if err != nil {
		log.Fatalf("Error writing json file in: menu_repository.go -> PostMenuDal %v", err)
	}
}

func GetMenuDal(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
