package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"referral-server/api/presenter"
	"referral-server/entity"
	"referral-server/usecase/contributor"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// CreateContribution rest api handler
func CreateContribution(service contributor.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error make contribution"

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		AccessToken := splitToken[1]

		// AccessToken := r.Header.Get("Authorization")[7:]

		var input struct {
			Email string `json:"email"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		err = service.Contribute(input.Email, AccessToken)

		if val, ok := entity.ErrCodeMapper[err]; ok {
			toJ := &presenter.AdditionalStatus{
				StatusCode:    val,
				StatusMessage: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toJ)
			return
		}
		if err == entity.ErrUnauthorizedAccess {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(entity.ErrUnauthorizedAccess.Error()))
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Contribution{
			Status: "Contribution is counted",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(toJ)
	})
}

// ListContributor handler
func ListContributor(service contributor.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading Contributor"

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		AccessToken := splitToken[1]

		data, err := service.ListContributor(AccessToken)

		if val, ok := entity.ErrCodeMapper[err]; ok {
			toJ := &presenter.AdditionalStatus{
				StatusCode:    val,
				StatusMessage: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toJ)
			return
		}
		if err == entity.ErrUnauthorizedAccess {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(entity.ErrUnauthorizedAccess.Error()))
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		var Contributors []*presenter.Contributor
		for _, d := range data {

			Contributors = append(Contributors, &presenter.Contributor{
				Email:        d.Email,
				ReferralLink: d.ReferralLink,
				Contribution: d.Contribution,
			})
		}

		if err := json.NewEncoder(w).Encode(Contributors); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakeContributorHandlers make url handlers
func MakeContributorHandlers(r *mux.Router, n negroni.Negroni, service contributor.UseCase) {

	r.Handle("/contributor/contribute", n.With(
		negroni.Wrap(CreateContribution(service)),
	)).Methods("POST", "OPTIONS").Name("createContribution")

	r.Handle("/contributor/", n.With(
		negroni.Wrap(ListContributor(service)),
	)).Methods("GET", "OPTIONS").Name("listContributor")

}
