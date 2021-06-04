package main

import (
	"fmt"
	"github.com/djedjethai/mytry/pkg/adding"
	"github.com/djedjethai/mytry/pkg/deleting"
	"github.com/djedjethai/mytry/pkg/http/rest"
	"github.com/djedjethai/mytry/pkg/listing"
	"github.com/djedjethai/mytry/pkg/reviewing"
	"github.com/djedjethai/mytry/pkg/storage/database"
	"github.com/djedjethai/mytry/pkg/storage/model"
	"github.com/djedjethai/mytry/pkg/updating"
	"log"
	"net/http"
)

type Type int

const JSON Type = iota

func main() {

	var adder adding.Service
	var lister listing.Service
	var reviewer reviewing.Service
	var deleter deleting.Service
	var updater updating.Service

	s, _ := model.NewStorage()
	sdb, _ := database.NewStorage()

	adder = adding.NewService(s, sdb)
	lister = listing.NewService(s)
	reviewer = reviewing.NewService(s)
	deleter = deleting.NewService(s)
	updater = updating.NewService(s)

	router := rest.Handler(adder, lister, reviewer, deleter, updater)

	fmt.Println("The beer server is on tap at: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
