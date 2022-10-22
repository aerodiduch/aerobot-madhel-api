package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slices"
)

var ctx = context.Background()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("./html/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error renderizando html", err)
	}

}

func retrieveEnvVariables(variable string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	return os.Getenv(variable)
}

func RetrieveData(w http.ResponseWriter, adKey string) {
	rdb := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     retrieveEnvVariables("REDIS_HOST"),
		Password: retrieveEnvVariables("REDIS_PASSWORD"),
		DB:       0,
	})

	val, err := rdb.Get(ctx, adKey).Result()
	switch {
	case err == redis.Nil:
		w.Header().Set("Content-Type", "application/json")
		jsonData := []byte(`{"detail":"Not found."}`)
		w.Write(jsonData)
	case err != nil:
		panic(err)
	}
	fmt.Fprintf(w, val)

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {

	currentTime := time.Now()

	auth_keys := []string{retrieveEnvVariables("API_KEY_1"), retrieveEnvVariables("API_KEY_2")}
	vars := mux.Vars(r)["key"]
	auth := r.Header.Get("Authorization")

	if !slices.Contains(auth_keys, auth) {
		fmt.Fprintf(w, "{'error':'Not authorized'}")
		return
	}

	if len(vars) == 3 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		RetrieveData(w, vars)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		jsonData := []byte(`{"detail":"Not found."}`)
		w.Write(jsonData)
		//fmt.Fprintf(w, "{'detail':'Not found.'}")

	}
	fmt.Println(currentTime.Format("2006-01-02 15:04:2"), r.Method, r.RequestURI, "-", r.UserAgent())

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/json/{key}", jsonHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Println("Listening...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe Error:\n", err)
	}
}

// if err := srv.ListenAndServe(); err != nil {
// 	log.Fatal("ListenAndServe: ", err)
