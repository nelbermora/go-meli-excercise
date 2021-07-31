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
	querySellersByLocality = `SELECT l.*,(
									SELECT count(id)
	                            	  FROM sellers
									 WHERE sellers.locality_id = l.id) as cant 
								FROM localities l
							   WHERE l.id = IFNULL(?,L.ID)`
	queryCarriesByLocality = `SELECT l.*,(
								SELECT count(id)
								  FROM carries
								 WHERE carries.locality_id = l.id) as cant 
							FROM localities l
						   WHERE l.id = IFNULL(?,L.ID)`
)

type Repository interface {
	Save(ctx context.Context, l domain.Locality) (int, error)
	Exists(ctx context.Context, id int) bool
	GetSellersByLoc(ctx context.Context, id int) ([]domain.Locality, error)
	GetCarriesByLoc(ctx context.Context, id int) ([]domain.Locality, error)
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

func (r *repository) GetSellersByLoc(ctx context.Context, id int) ([]domain.Locality, error) {
	var localities []domain.Locality
	var rows *sql.Rows
	var err error
	if id > 0 {
		rows, err = r.db.Query(querySellersByLocality, id)
	} else {
		rows, err = r.db.Query(querySellersByLocality, nil)
	}

	for rows.Next() {
		l := domain.Locality{}
		err = rows.Scan(&l.ID, &l.Name, &l.Province, &l.Country, &l.Sellers)
		localities = append(localities, l)
	}

	if err != nil {
		return localities, err
	}
	return localities, nil
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

func (r *repository) GetCarriesByLoc(ctx context.Context, id int) ([]domain.Locality, error) {
	var localities []domain.Locality
	var rows *sql.Rows
	var err error
	if id > 0 {
		rows, err = r.db.Query(queryCarriesByLocality, id)
	} else {
		rows, err = r.db.Query(queryCarriesByLocality, nil)
	}

	for rows.Next() {
		l := domain.Locality{}
		err = rows.Scan(&l.ID, &l.Name, &l.Province, &l.Country, &l.Carries)
		localities = append(localities, l)
	}

	if err != nil {
		return localities, err
	}
	return localities, nil
}
