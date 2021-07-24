package main

import (
	"Referral-System/generator/api/handler"
	"Referral-System/generator/api/middleware"
	"Referral-System/generator/config"
	"Referral-System/generator/infrastructure/repository"
	"Referral-System/generator/usecase/generator"

	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {

	config.LoadEnv()

	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_NAME,
		config.DB_PASSWORD)

	db, err := sql.Open("postgres", DBURI)
	if err != nil {
		log.Fatal("Cannot connect to database : ", err)
	}
	defer db.Close()

	log.Println("We are connected to the database ", os.Getenv("DB_NAME"))

	r := mux.NewRouter().StrictSlash(true)

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	generatorRepo := repository.NewGeneratorSQL(db)
	generatorService := generator.NewService(generatorRepo)

	//Generator
	handler.MakeGeneratorHandlers(r, *n, generatorService)
	handler.MakeReferralHandlers(r, *n)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + config.API_PORT,
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	log.Printf("\nServer starting on port %s", config.API_PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
