package repository

import (
	context "context"

	"github.com/manuonda/go-projects/inventario/internal/entitys"
	entity "github.com/manuonda/go-projects/inventario/internal/entitys"
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
	queryAll = `select name, description,price, created_by from PRODUCTS;`

	queryGetById = `select id, name, description, price,created_by from PRODUCTS
	where id=?`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entitys.Producto, error) {
	pp := []entity.Producto{}
	err := r.db.SelectContext(ctx, &pp, queryAll)
	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (r *repo) GetProduct(ctx context.Context, ID int64) (*entitys.Producto, error) {
	producto := &entity.Producto{}
	err := r.db.GetContext(ctx, producto, queryGetById, ID)
	if err != nil {
		return nil, err
	}
	return producto, nil
}
