package main

import (
	"database/sql"
	"database_external/brokers"
	cache "database_external/cahce"
	"database_external/entity"
	"database_external/repository"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	db    *sql.DB
	dbErr error
)

func main() {

	// preparing db connecting driver
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbName := "godb"
	psgInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, dbErr = sql.Open("postgres", psgInfo)
	if dbErr != nil {
		panic(dbErr)
	}

	cityRepo := repository.NewRepo(db)
	rabbitmq := brokers.NewRabbitMQ()
	redisCache := cache.NewRedisCahce()

	http.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if r.URL.Query().Has("id") {
				cityIdStr := r.URL.Query().Get("id")
				cityId, _ := strconv.Atoi(cityIdStr)
				city := cityRepo.GetById(cityId)

				json.NewEncoder(w).Encode(city)
				return
			}

			cityCache := redisCache.Get()
			if cityCache != nil {
				fmt.Println("city cachede var")
				w.Write(cityCache)
				return
			}

			fmt.Println("city cachede yok")

			cityList := cityRepo.List()
			cityBytes, err := json.Marshal(cityList)
			if err != nil {
				fmt.Println(err)
			}

			go func(data []byte) {
				fmt.Println("city cacheleniyor")
				redisCache.Put(data)
			}(cityBytes)

			w.Write(cityBytes)

		} else if r.Method == http.MethodPost {
			var city entity.City

			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = json.Unmarshal(bodyBytes, &city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			cityRepo.Insert(city)
			rabbitmq.Publish([]byte("city created with name " + city.Name))
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, "unsupported", http.StatusMethodNotAllowed)
		}

	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	<-make(chan struct{})
}
