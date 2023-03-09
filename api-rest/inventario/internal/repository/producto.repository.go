package repository

import (
	context "context"

	"github.com/manuonda/go-projects/inventario/internal/entitys"
)

const (
	queryInsertProduct = `
	  insert into PRODUCTS(
		name, 
		description,
		price,
		created_by,
	  )
	  values(?,?,?,?)
	`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_
}

func (r *repo) GetProducts(ctx context.Context) ([]entitys.Producto, error) {
	return nil, nil
}

func (r *repo) GetProduct(ctx context.Context, ID int64) (entitys.Producto, error) {
	return nil, nil
}
