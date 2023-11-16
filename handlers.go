package main

import (
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Login struct {
	Email    string `json"email"`
	Password string `json"password"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
	Email     string `json:"email"`
	Message   string `json:"message"`
}

func SignupHandler(res http.ResponseWriter, req *http.Request) {

}

func SigninHandler(res http.ResponseWriter, req *http.Request) {

}

func SignoutHandler(res http.ResponseWriter, req *http.Request) {
}
