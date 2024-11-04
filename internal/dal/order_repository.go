package dal

import (
	"encoding/json"
	"log"
	"os"

	"hot-coffee/models"
)

func OrderPost(order models.Order) {
	writeJson, _ := json.MarshalIndent(order, "", "\t")
	err := os.WriteFile(*models.Dir + "/orders.json", writeJson, 0o644)
	if err != nil {
		log.Fatalf("Error writing json file in: order_repository.go -> OrderPost %v", err)
	}
}