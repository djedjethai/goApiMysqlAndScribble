package deleting

type Service interface {
	DeleteReviewS(string) error
	DeleteBeerS(string) error
}

type Repository interface {
	DeleteReview(string) error
	DeleteBeer(string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteBeerS(bid string) error {
	return s.r.DeleteBeer(bid)
}

func (s *service) DeleteReviewS(bid string) error {
	return s.r.DeleteReview(bid)
}
