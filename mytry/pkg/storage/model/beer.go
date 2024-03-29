package model

import (
	"time"
)

type Beer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       float32   `json:"abv"`
	ShortDesc string    `json:"short_desc"`
	Created   time.Time `json:"created"`
}
