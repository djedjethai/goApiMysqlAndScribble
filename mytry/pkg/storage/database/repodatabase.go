package database

import (
	// "encoding/json"
	"database/sql"
	"fmt"
	"github.com/djedjethai/mytry/pkg/adding"
	_ "github.com/go-sql-driver/mysql"
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

const (
	username = "root"
	password = "root"
	hostname = "mysql:3306"
	database = "goapi"
)

type Storage struct {
	db *sql.DB
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, hostname, database)
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = sql.Open("mysql", dsn())
	// s.db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/goapi?charset=utf8")
	fmt.Println(err)
	defer s.db.Close()

	err = s.db.Ping()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	query := `CREATE TABLE IF NOT EXISTS beer(beer_id int primary key auto_increment, beer_name VARCHAR(20), beer_brewery VARCHAR(20), beer_abv FLOAT(25), beer_shortdesc text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	stmt, err := s.db.Prepare(query)
	fmt.Printf("err at create table: %v", err)
	defer stmt.Close()

	r, err := stmt.Exec()
	fmt.Println(err)

	n, err := r.RowsAffected()
	fmt.Print(err)

	fmt.Printf("table beer created: %v", n)

	return s, nil
}

func (s *Storage) AddBeerDB(beer adding.Beer) (string, error) {

	// query
	fmt.Println("add some beer, cooool")

	return "fine", nil
}
