package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5433 user=postgres dbname=postgres sslmode=disable password=goLANGn1nja")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECT TO DB")

	// call to db:

	err = insertUser(db, User{Name: "vova", Email: "vova@yohu.com", Password: "123"})
	if err != nil {
		log.Fatal(err)
	}

	err = updateUser(db, User{Email: "looooser@bobj.com"}, 3)
	if err != nil {
		log.Fatal(err)
	}

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

	user, err := getUserID(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

}

func getUsers(db *sql.DB) ([]User, error) {
	// select all row
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUserID(db *sql.DB, id int) (User, error) {
	//	select from qery plceholder
	var u User
	err := db.QueryRow("select * from users where id = $1", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	return u, err
}

func insertUser(db *sql.DB, u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into users (name, email, password) values ($1, $2, $3)",
		u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into logs (entity, action) values ($1, $2)",
		"user", "created")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func deletetUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1", id)
	return err
}

func updateUser(db *sql.DB, newUser User, id int) error {
	_, err := db.Exec("update users set email = $1 where id = $2", newUser.Email, id)
	return err
}
