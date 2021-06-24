package product

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO "main"."products"("description","expiration_rate","freezing_rate","height","lenght","netweight","product_code","recommended_freezing_temperature","width","id_product_type","id_seller") VALUES (?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil{
		return 0,err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)
	if err != nil{
		return 0,err
	}

	id, err := res.LastInsertId()
	if err != nil{
		return 0,err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	fmt.Println(p)

	return nil
}