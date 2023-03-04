package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode string `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "mkzero:0147@tcp(192.168.11.11:3306)/owncloud")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return

	}
	Db = database
}

func main() {
	conn, err := Db.Begin()
	if err != nil {
		fmt.Println("begin failed : ", err)
		return
	}

	r, err := conn.Exec("insert into person(username, sex,email)values(?,?,?)", "agem", "man", "stumail")
	if err != nil {
		fmt.Println("exec failed , ", err)
		conn.Rollback()
		return

	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)
	r, err = conn.Exec("insert into person(username,sex,email)values(?,?,?)", "mk1", "man", "mk@gmail.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("succ:", id)
	conn.Commit()
}
