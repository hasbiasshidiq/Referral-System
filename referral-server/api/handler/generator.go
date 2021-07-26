package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"referral-server/api/presenter"
	"referral-server/entity"
	"referral-server/usecase/generator"
	"strings"

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

// LoginGenerator rest api handler
func LoginGenerator(service generator.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Login Generator"

		var input struct {
			ID       string `json:"generator_id"`
			Password string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		AccessToken, err := service.LoginGenerator(input.ID, input.Password)

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
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.LoginGenerator{
			AccessToken: AccessToken,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(toJ)
	})
}

// ExtendGenerator rest api handler
func ExtendGenerator(service generator.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error extending time"

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")

		if len(splitToken) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		AccessToken := splitToken[1]

		// AccessToken := r.Header.Get("Authorization")[7:]
		ExpirationTime, err := service.ExtendTime(AccessToken)

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

		toJ := &presenter.ExtendingTime{
			Status: "Success Extending Link till " + ExpirationTime,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(toJ)
	})
}

//MakeGeneratorHandlers make url handlers
func MakeGeneratorHandlers(r *mux.Router, n negroni.Negroni, service generator.UseCase) {

	r.Handle("/generator/register", n.With(
		negroni.Wrap(CreateGenerator(service)),
	)).Methods("POST", "OPTIONS").Name("createGenerator")

	r.Handle("/generator/login", n.With(
		negroni.Wrap(LoginGenerator(service)),
	)).Methods("POST", "OPTIONS").Name("loginGenerator")

	r.Handle("/generator/extend-time", n.With(
		negroni.Wrap(ExtendGenerator(service)),
	)).Methods("PUT", "OPTIONS").Name("extendGenerator")

}
