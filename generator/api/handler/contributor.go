package handler

import (
	"Referral-System/generator/api/presenter"
	"Referral-System/generator/entity"
	"Referral-System/generator/usecase/contributor"
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// CreateContribution rest api handler
func CreateContribution(service contributor.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error make contribution"

		AccessToken := r.Header.Get("Authorization")[7:]

		var input struct {
			Email string `json:"email"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println("CreateContribution error handle input : ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
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

		toJ := &presenter.Contributor{
			Status: "Contribution is counted",
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(toJ)
	})
}

//MakeContributorHandlers make url handlers
func MakeContributorHandlers(r *mux.Router, n negroni.Negroni, service contributor.UseCase) {

	r.Handle("/contributor/contribute", n.With(
		negroni.Wrap(CreateContribution(service)),
	)).Methods("POST", "OPTIONS").Name("createContribution")

}
