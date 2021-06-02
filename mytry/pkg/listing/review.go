package listing

import "time"

type Review struct {
	ID        string    `json:"id"`
	BeerID    string    `json:"beerid"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Score     float32   `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}
