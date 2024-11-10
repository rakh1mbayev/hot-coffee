package service

import "hot-coffee/internal/dal"

type Serv_inventory interface{

}

type Inventory_serv struct {
	inventory_dal dal.Inventory_dal
}

func NewDefault_servinventory(inventory_dal *dal.Inventory_dal) *Inventory_serv {
	return &Inventory_serv{inventory_dal: *inventory_dal}
}

