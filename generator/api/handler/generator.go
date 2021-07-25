package handler

import (
	"Referral-System/generator/api/presenter"
	"Referral-System/generator/entity"
	"Referral-System/generator/usecase/generator"
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// CreateGenerator rest api handler
func CreateGenerator(service generator.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding Generator"

		var input struct {
			ID       string `json:"generator_id"`
			Name     string `json:"generator_name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		generatedLink, err := service.CreateGenerator(input.ID, input.Name, input.Email, input.Password)

		if err == entity.ErrAlreadyExist {
			toJ := &presenter.AdditionalStatus{
				StatusCode:    entity.ErrCodeMapper[err],
				StatusMessage: err.Error(),
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(toJ)
			return
		}

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Generator{
			GeneratedLink: generatedLink,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(toJ)
	})
}

//MakeGeneratorHandlers make url handlers
func MakeGeneratorHandlers(r *mux.Router, n negroni.Negroni, service generator.UseCase) {

	r.Handle("/generator/register", n.With(
		negroni.Wrap(CreateGenerator(service)),
	)).Methods("POST", "OPTIONS").Name("createGenerator")

}