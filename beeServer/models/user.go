package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

func ConnectDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateUser(db *sqlx.DB, user *User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	return db.QueryRowx(query, user.Username, user.Password).Scan(&user.ID)
}

func GetUserByUsername(db *sqlx.DB, username string) (*User, error) {
	var user User
	query := `SELECT id, username, password FROM users WHERE username=$1`
	err := db.Get(&user, query, username)
	return &user, err
}
