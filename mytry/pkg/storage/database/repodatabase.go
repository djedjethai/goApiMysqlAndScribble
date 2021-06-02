package database

import (
	// "encoding/json"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/djedjethai/mytry/pkg/adding"
	// "github.com/djedjethai/mytry/pkg/listing"
	// "github.com/djedjethai/mytry/pkg/reviewing"
	// "github.com/djedjethai/mytry/pkg/storage"
	// "github.com/djedjethai/mytry/pkg/updating"
	// "github.com/nanobox-io/golang-scribble"
	// "log"
	// "path"
	// "runtime"
	// "time"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/goapi?charset=utf8")
	fmt.Println(err)
	defer s.db.Close()

	err = s.db.Ping()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare(`CREATE TABLE beer (name VARCHAR(20));`)
	fmt.Printf("err at create table: %v", err)
	defer stmt.Close()

	r, err := stmt.Exec()
	fmt.Println(err)

	n, err := r.RowsAffected()
	fmt.Print(err)

	fmt.Printf("table beer created: %v", n)

	return s, nil
}
