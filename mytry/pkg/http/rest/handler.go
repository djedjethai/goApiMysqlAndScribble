package rest

import (
	"encoding/json"
	"github.com/djedjethai/mytry/pkg/adding"
	"github.com/djedjethai/mytry/pkg/deleting"
	"github.com/djedjethai/mytry/pkg/listing"
	"github.com/djedjethai/mytry/pkg/reviewing"
	"github.com/djedjethai/mytry/pkg/updating"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Handler(a adding.Service, l listing.Service, rs reviewing.Service, d deleting.Service, u updating.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", GetAllBeersR(l))
	router.GET("/beer/:id", GetBeerR(l))
	router.POST("/beer", PostBeerR(a))
	router.GET("/beer/:id/review", GetBeerReviewR(l))
	router.POST("/beer/review/", PostBeerReviewR(rs))
	router.DELETE("/beer/:id/delreview", DeleteBeerReviewR(d))
	router.DELETE("/beer/:id/delbeer", DeleteBeerR(d))
	router.POST("/review/update", UpdateReviewR(u))
	router.POST("/beer/update", UpdateBeerR(u))

	return router
}

func UpdateBeerR(u updating.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var beerToUpdate listing.Beer
		if err := decoder.Decode(&beerToUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		str, err := u.UpdateBeerS(beerToUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(str)
	}
}

func UpdateReviewR(u updating.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var revToUpdate listing.Review
		if err := decoder.Decode(&revToUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		str, err := u.UpdateReviewS(revToUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(str)
	}
}

func DeleteBeerR(d deleting.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := d.DeleteBeerS(p.ByName("id")); err != nil {
			http.Error(w, "error in deleting beer", http.StatusInternalServerError)
		}

		w.Header().Set("Content-type", "application/json")
	}
}

func DeleteBeerReviewR(d deleting.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := d.DeleteReviewS(p.ByName("id")); err != nil {
			http.Error(w, "err occur during deleting", http.StatusInternalServerError)
		}

		w.Header().Set("Content-type", "application/json")
	}
}

func PostBeerReviewR(rs reviewing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var nRev reviewing.Review
		if err := decoder.Decode(&nRev); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		str, err := rs.PostBeerReviewS(nRev)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(str)
	}
}

func GetBeerReviewR(l listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-type", "application/json")
		rvs, err := l.GetBeerReviewS(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(rvs)
	}
}

func GetAllBeersR(l listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-type", "application/json")
		list, err := l.GetAllBeersS()
		if err != nil {
			http.Error(w, "Beers unfound", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(list)
	}
}

func GetBeerR(l listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		beer, err := l.GetBeerS(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The beer do not exist", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}

func PostBeerR(a adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var nbeer adding.Beer
		if err := decoder.Decode(&nbeer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		str, err := a.AddBeerS(nbeer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(str)
	}
}
