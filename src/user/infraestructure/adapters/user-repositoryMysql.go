package adapters

import (
	"api-go/src/database"
	"api-go/src/user/domain/entities"
	"database/sql"
)

type UserReposirtoyMysql struct {
	DB *sql.DB
}


func NewUserRepository() (*UserReposirtoyMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	return &UserReposirtoyMysql{DB: db}, nil
}


func (r *UserReposirtoyMysql) Create(user entities.User) (entities.User, error) {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.User{}, err
	}

	defer stmt.Close()

	result, err := r.DB.Exec(query, user.Email, user.Name, user.Password)
	if err != nil {
		return entities.User{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.User{}, err
	}

	user.ID = int(id)
	user.Password = ""

	return user, nil
}

func (r *UserReposirtoyMysql) GetByID(id int64) (entities.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = ?`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.User{}, err
	}

	defer stmt.Close()
	
	row := stmt.QueryRow(id)

	var user entities.User

	err = row.Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows{
		return entities.User{}, err
	} else if err != nil {
		return entities.User{}, err
	}

	return user, nil

}

func (r *UserReposirtoyMysql) Delete(id int64)(bool, error){
	query := `DELETE FROM users WHERE id = ?`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}

func (r *UserReposirtoyMysql) GetByEmail(email string) (entities.User, error){
	query := `SELECT id, name, email, password FROM users WHERE email = ?`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.User{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)

	var user entities.User

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows{
		return entities.User{}, err
	} else if err != nil {
		return entities.User{}, err
	}

	return user, nil
}