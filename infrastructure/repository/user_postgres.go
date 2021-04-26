package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/FindIdols/findidols-back/entity"
)

//UserPostgres postgres repo
type UserPostgres struct {
	db *sql.DB
}

//NewUserPostgres create new repository
func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

//Create repo
func (r *UserPostgres) Create(u *entity.User) (*entity.User, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO direct_users (user_id, first_name, last_name, email, phone, genre, category, created_at) 
		VALUES( $1,$2,$3,$4,$5,$6,$7,$8 )`)

	if err != nil {
		fmt.Println("erro criar user no banco")
		return u, err
	}
	_, err = stmt.Exec(
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		u.Genre,
		u.Category,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return u, err
	}

	err = stmt.Close()
	if err != nil {
		return u, err
	}

	return u, nil
}

//Get a user
func (r *UserPostgres) Get(id entity.ID) (*entity.User, error) {
	stmt, err := r.db.Prepare(`select id, name, last_name, email, phone, 
	usage, subject, instruction, created_at from idols where id = ?`)

	if err != nil {
		return nil, err
	}

	var user entity.User

	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Genre,
			&user.Email,
			&user.Phone,
			&user.Category,
			&user.CreatedAt,
		)
	}

	return &user, nil
}
