package locality

import (
	"context"
	"database/sql"
	"log"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

const (
	insertLocality         = `INSERT INTO localities(id,locality_name,province_name,country_name) VALUES (?,?,?,?)`
	queryExists            = `SELECT id from localities where id = ?`
	queryGetById           = `SELECT * from localities where id = ?`
	querySellersByLocality = `SELECT locality_id,locality_name,province_name, country_name, count(locality_id) as sellers FROM localities INNER JOIN sellers ON localities.id = sellers.locality_id
								 AND localities.id = ?
							GROUP BY locality_name, locality_id, province_name, country_name`
)

type Repository interface {
	Save(ctx context.Context, l domain.Locality) (int, error)
	Exists(ctx context.Context, id int) bool
	GetSellersByLoc(ctx context.Context, id int) (domain.Locality, error)
	Get(ctx context.Context, id int) (domain.Locality, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, l domain.Locality) (int, error) {
	stmt, err := r.db.Prepare(insertLocality)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&l.ID, &l.Name, &l.Province, &l.Country)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(queryExists, id)
	err := row.Scan(&id)
	log.Println(err)
	return err == nil
}

func (r *repository) GetSellersByLoc(ctx context.Context, id int) (domain.Locality, error) {
	row := r.db.QueryRow(querySellersByLocality, id)
	l := domain.Locality{}
	err := row.Scan(&l.ID, &l.Name, &l.Province, &l.Country, &l.Sellers)
	if err != nil {
		return domain.Locality{}, err
	}
	return l, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Locality, error) {
	row := r.db.QueryRow(queryGetById, id)
	l := domain.Locality{}
	err := row.Scan(&l.ID, &l.Name, &l.Province, &l.Country)
	if err != nil {
		return domain.Locality{}, err
	}
	return l, nil
}
