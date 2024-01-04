package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var db *redis.Client        // for storing user athentication information
var sessionDB *redis.Client // for managing sessions
var ctx = context.Background()

func Init() {
	db = redis.NewClient(&redis.Options{
		Addr:     "redisDB:6379",
		Password: "",
		DB:       0,
	})
	sessionDB = redis.NewClient(&redis.Options{
		Addr:     "sessionsDB:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	Init()
	fmt.Println("Authentication with custom router in golang")
	router := NewRouter()

	router.POST("/api/v1/signup", SignupHandler)
	router.POST("/api/v1/signin", SigninHandler)
	router.POST("/api/v1/signout", SignoutHandler)

	fmt.Println("Server starting at port 4000")
	if err := checkDBConnection(); err != nil {
		panic("Database refused to connect, shutting down server")
	}
	log.Fatal(http.ListenAndServe(":4000", router))
}
