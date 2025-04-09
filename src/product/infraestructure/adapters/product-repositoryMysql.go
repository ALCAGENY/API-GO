package adapters

import (
	"api-go/src/database"
	"api-go/src/product/domain/entities"
	"database/sql"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}


func NewProductRepositoryMysql() (*ProductRepositoryMysql, error){
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	return &ProductRepositoryMysql{DB: db}, nil
}

func (r *ProductRepositoryMysql) Create(product entities.Product) (entities.Product, error) {
	query := `INSERT INTO products (name, fecha_adquisicion) VALUES (?, ?)`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.Product{}, err
	}

	defer stmt.Close()

	result, err := r.DB.Exec(query, product.Name, product.Fecha_Adquisicion)
	if err != nil {
		return entities.Product{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.Product{}, err
	}

	product.ID = int(id)

	return product, nil
}

func (r *ProductRepositoryMysql) GetByID(id int64) (entities.Product, error) {
	query := `SELECT id, name, fecha_adquisicion FROM products WHERE id = ?`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.Product{}, err
	}

	defer stmt.Close()

	var product entities.Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Fecha_Adquisicion)

	if err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func (r *ProductRepositoryMysql) Delete(id int64) (bool, error) {
	query := `DELETE FROM products WHERE id = ?`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	return true, nil
}