package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const secret = "my_secret_key"

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
	Email     string `json:"email"`
	Message   string `json:"message"`
}

func SignupHandler(res http.ResponseWriter, req *http.Request) {
	// extract user details from the request
	// check if user already exists in database
	// if yes, then send error, if not create new user
	// hash password
	// store email and hashed password in database
	// return success

	result, err := io.ReadAll(req.Body)
	if err != nil {
		InternalServerError(res)
		return
	}
	var user User
	err = json.Unmarshal(result, &user)
	if err != nil {
		InternalServerError(res)
		return
	}
	// extract user details
	name := user.Name
	email := user.Email
	password := user.Password

	// check if name, email and password are valid
	if len(name) <= 0 || len(email) <= 0 || len(password) <= 0 {
		BadRequestError(res)
		return
	}

	// check if user already exists in database
	userExists, err := IsUserDuplicate(db, email)
	if err != nil {
		InternalServerError(res)
		return
	}
	if userExists {
		ConflictError(res)
		return
	}

	// if user does not exist, then we can continue
	hashedPassword, err := hashPassword(password, secret)
	if err != nil {
		InternalServerError(res)
		return
	}
	err = CreateNewUser(db, name, email, hashedPassword)
	if err != nil {
		InternalServerError(res)
		return
	}

	// if everything worked perfectly, then send back success message
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("User successfully created"))
}

func SigninHandler(res http.ResponseWriter, req *http.Request) {
	// when signing in, user sends email and password
	// check if user exists using email
	// if true, hash password and compare to stored hashed password
	// if equal, return JSON with auth_token and add a reference into sessions database

	result, err := io.ReadAll(req.Body)
	if err != nil {
		InternalServerError(res)
		return
	}
	var userCred Login
	err = json.Unmarshal(result, &userCred)
	if err != nil {
		InternalServerError(res)
		return
	}

	email := userCred.Email
	password := userCred.Password

	if len(email) <= 0 || len(password) <= 0 {
		BadRequestError(res)
		return
	}

	userExists, err := IsUserDuplicate(db, email)
	if err != nil {
		InternalServerError(res)
		return
	}

	if !userExists {
		BadRequestError(res, "User doesn't exist, can't login")
		return
	}

	hashedPassword, err := hashPassword(password, secret)
	if err != nil {
		InternalServerError(res)
		return
	}

	_, err = GetUser(db, email, hashedPassword)
	if err != nil {
		BadRequestError(res, err.Error())
		return
	}

	var response LoginResponse
	payload := make(map[string]string)
	payload["email"] = email
	payload["password"] = password
	response.AuthToken, err = generateToken(payload, secret)
	if err != nil {
		InternalServerError(res)
		return
	}
	response.Email = email
	response.Message = "User successfully signed in"

	result, err = json.Marshal(response)
	if err != nil {
		InternalServerError(res)
		return
	}

	//set this auth token as a session into the database
	err = CreateSession(sessionDB, response.AuthToken)
	if err != nil {
		if err.Error() == "User is already logged in" {
			BadRequestError(res, "User is already logged in")
			return
		}

		BadRequestError(res)
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func SignoutHandler(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	authToken := strings.Split(authHeader, " ")[1]
	err := DeleteSession(sessionDB, authToken)
	if err != nil {
		if err.Error() == "User not logged in" {
			BadRequestError(res, "User not logged in")
			return
		}
		BadRequestError(res)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Successfully logged out"))
}
