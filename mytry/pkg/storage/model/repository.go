package model

import (
	"encoding/json"
	"fmt"
	"github.com/djedjethai/mytry/pkg/adding"
	"github.com/djedjethai/mytry/pkg/listing"
	"github.com/djedjethai/mytry/pkg/reviewing"
	"github.com/djedjethai/mytry/pkg/storage"
	"github.com/djedjethai/mytry/pkg/updating"
	"github.com/nanobox-io/golang-scribble"
	"log"
	"path"
	"runtime"
	"time"
)

const (
	dir              = "/data/"
	CollectionBeer   = "beers"
	CollectionReview = "reviews"
)

type Storage struct {
	db *scribble.Driver
}

func NewStorage() (*Storage, error) {

	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) UpdateBeer(updatedBeer listing.Beer) (string, error) {
	oldBeer, err := s.GetBeer(updatedBeer.ID)
	if err != nil {
		return "", updating.ErrRegister
	}

	// delete old beer
	if err := s.DeleteBeer(oldBeer.ID); err != nil {
		return "", updating.ErrRegister
	}

	// save new Beer (using UpdateBeerAction)
	if err := s.AddUpdatedBeer(updatedBeer); err != nil {
		return "", updating.ErrRegister
	}

	return "beer updated", nil
}

func (s *Storage) AddUpdatedBeer(upBeer listing.Beer) error {
	// save one beer( from listing.Beer )
	if err := s.db.Write(CollectionBeer, upBeer.ID, upBeer); err != nil {
		return updating.ErrRegister
	}

	return nil
}

func (s *Storage) UpdateReview(oldRvs listing.Review) (string, error) {
	var newRvs listing.Review
	// find rv
	listingReviews, err := s.GetBeerReviews(oldRvs.BeerID)
	if err != nil {
	}

	for _, rv := range listingReviews {
		if rv.ID == oldRvs.ID {
			newRvs.ID = oldRvs.ID
			newRvs.BeerID = oldRvs.BeerID
			newRvs.FirstName = oldRvs.FirstName
			newRvs.LastName = oldRvs.LastName
			newRvs.Score = oldRvs.Score
			newRvs.Text = oldRvs.Text
			newRvs.Created = time.Now()

			if err := s.DeleteSingleReview(newRvs.ID); err != nil {
				return "", err
			}

			if err := s.AddUpdatedReview(newRvs); err != nil {
				return "", err
			}

			return "updated", nil
		}

		return "", updating.ErrRegister
	}

	return "", updating.ErrRegister
}

func (s *Storage) AddUpdatedReview(r listing.Review) error {

	if err := s.db.Write(CollectionReview, r.ID, r); err != nil {
		return updating.ErrRegister
	}

	return nil
}

func (s *Storage) DeleteSingleReview(rid string) error {
	if err := s.db.Delete(CollectionReview, rid); err != nil {
		return updating.ErrRegister
	}

	return nil
}

func (s *Storage) DeleteBeer(bid string) error {
	if err := s.DeleteReview(bid); err != nil {
		fmt.Printf("in delete beer, err from delete reviews: %v", err)
		return err
	}

	if err := s.db.Delete(CollectionBeer, bid); err != nil {
		fmt.Printf("err during deleting beer: %v", err)
		return err
	}

	fmt.Println("beer deleted succefully")
	return nil
}

func (s *Storage) DeleteReview(bid string) error {

	reviews, err := s.GetBeerReviews(bid)
	if err != nil {
		fmt.Printf("In delete Review, err occur at GetBeerReviews: %v", err)
		return err
	}

	for i := range reviews {
		if err := s.db.Delete(CollectionReview, reviews[i].ID); err != nil {
			fmt.Printf("an err occur happend at delete time: %v", err)
			return err
		}
	}

	fmt.Println("deleted successfully")
	return nil
}

func (s *Storage) GetBeerReviews(bid string) ([]listing.Review, error) {
	var reviews []listing.Review

	allReviews, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		return reviews, err
	}

	for _, rv := range allReviews {
		var nRev listing.Review
		var rev Review
		if err := json.Unmarshal([]byte(rv), &rev); err != nil {
			return reviews, err
		}

		if rev.BeerID == bid {
			nRev.ID = rev.ID
			nRev.BeerID = rev.BeerID
			nRev.FirstName = rev.FirstName
			nRev.LastName = rev.LastName
			nRev.Score = rev.Score
			nRev.Text = rev.Text
			nRev.Created = rev.Created

			reviews = append(reviews, nRev)
		}
	}

	return reviews, nil
}

func (s *Storage) AddBeerReview(r reviewing.Review) (string, error) {
	id, err := storage.CreateID("review")
	if err != nil {
		log.Fatal(err)
	}

	nReview := Review{
		ID:        id,
		BeerID:    r.BeerID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Score:     r.Score,
		Text:      r.Text,
		Created:   time.Now(),
	}

	if err := s.db.Write(CollectionReview, nReview.ID, nReview); err != nil {
		return "", reviewing.ErrRegister
	}

	return nReview.ID, nil
}

func (s *Storage) AddBeer(barg adding.Beer) (string, error) {
	id, err := storage.CreateID("beer")
	if err != nil {
		log.Fatal(err)
	}

	nBeer := Beer{
		ID:        id,
		Name:      barg.Name,
		Brewery:   barg.Brewery,
		Abv:       barg.Abv,
		ShortDesc: barg.ShortDesc,
		Created:   time.Now(),
	}

	if err := s.db.Write(CollectionBeer, nBeer.ID, nBeer); err != nil {
		return "", err
	}

	return nBeer.ID, nil
}

func (s *Storage) GetBeer(bid string) (listing.Beer, error) {

	var beerToRet listing.Beer
	var beerFromDB Beer

	if err := s.db.Read(CollectionBeer, bid, &beerFromDB); err != nil {
		return beerToRet, listing.ErrNotFound
	}

	beerToRet = listing.Beer{
		ID:        beerFromDB.ID,
		Name:      beerFromDB.Name,
		Brewery:   beerFromDB.Brewery,
		Abv:       beerFromDB.Abv,
		ShortDesc: beerFromDB.ShortDesc,
		Created:   beerFromDB.Created,
	}

	return beerToRet, nil
}

func (s *Storage) GetAllBeers() ([]listing.Beer, error) {
	beers := []listing.Beer{}

	allBeers, err := s.db.ReadAll(CollectionBeer)
	if err != nil {
		return beers, err
	}

	for _, r := range allBeers {
		var nbeer listing.Beer
		var b Beer

		if err := json.Unmarshal([]byte(r), &b); err != nil {
			return beers, err
		}

		nbeer.ID = b.ID
		nbeer.Name = b.Name
		nbeer.Brewery = b.Brewery
		nbeer.Abv = b.Abv
		nbeer.ShortDesc = b.ShortDesc
		nbeer.Created = b.Created

		beers = append(beers, nbeer)
	}

	return beers, nil
}

// func (s *Storage) ModifBeer(bid string) error {
//
// }
//
// func (s *Storage) DeleteBeer(bid string) error {
//
// }
