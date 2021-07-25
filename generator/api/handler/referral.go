package handler

import (
	"Referral-System/generator/api/presenter"
	"Referral-System/generator/config"
	"Referral-System/generator/entity"
	"Referral-System/generator/usecase/token"
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func GetReferral(service token.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error providing referral"
		vars := mux.Vars(r)

		ReferralLink := config.SHARED_LINK_ENDPOINT + vars["id"]
		fmt.Println("ReferralLink", ReferralLink)

		accessToken, err := service.CreateContributorToken(ReferralLink)

		if val, ok := entity.ErrCodeMapper[err]; ok {
			toJ := &presenter.AdditionalStatus{
				StatusCode:    val,
				StatusMessage: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toJ)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// if data == nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	w.Write([]byte(errorMessage))
		// 	return
		// }

		toJ := &presenter.Referral{
			AccessToken: accessToken,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(toJ)
	})
}

//MakeReferralHandlers make url handlers
func MakeReferralHandlers(r *mux.Router, n negroni.Negroni, service token.UseCase) {

	r.Handle("/referral/{id}", n.With(
		negroni.Wrap(GetReferral(service)),
	)).Methods("GET", "OPTIONS").Name("getReferral")

}
