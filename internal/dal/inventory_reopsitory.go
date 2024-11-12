package dal

type InventoryDalInterface interface {
}

type InventoryFileDataAccess struct {
	path string
}

func NewInventoryRepo(path string) *InventoryFileDataAccess {
	return &InventoryFileDataAccess{path: path}
}
