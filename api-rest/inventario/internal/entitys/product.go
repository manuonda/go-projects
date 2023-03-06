package entitys

type Producto struct {
	ID          int64   `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	Created_by  int64   `db:"created_by" json:"created_by"`
}
