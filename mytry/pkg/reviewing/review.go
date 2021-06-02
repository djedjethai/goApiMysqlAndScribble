package reviewing

type Review struct {
	BeerID    string  `json:"beerid"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Score     float32 `json:"score"`
	Text      string  `json:"text"`
}
