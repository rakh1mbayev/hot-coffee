package dal

type Dal_inventory interface {

}

type Inventory_dal struct {
	DatapathInventory string
}

func NewInventoryRepo(datapath string) *Inventory_dal {
	return &Inventory_dal{DatapathInventory: datapath}
}