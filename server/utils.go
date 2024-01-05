package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

// function to check is db is connected, otherwise shut down server
func checkDBConnection() error {
	_, redisErr := db.Ping(ctx).Result()
	if redisErr != nil {
		return redisErr
	}
	_, redisErr = sessionDB.Ping(ctx).Result()
	if redisErr != nil {
		return redisErr
	}
	return nil
}

// generate auth token / custom jwt token
// jwt has three part, 'header', 'payload' and 'signature'
func generateToken(payload map[string]string, secret string) (string, error) {
	var headerMap = make(map[string]string)
	headerMap["type"] = "jwt"
	headerMap["alg"] = "HMAC256"
	header, err := json.Marshal(headerMap)
	if err != nil {
		log.Fatal("Error generating jwt token")
		return "", err
	}
	h := hmac.New(sha256.New, []byte(secret)) // create a new hash of type 256

	// convert header to base64 encoded string
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	// Marshal payload into json string
	payloadString, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error generating token")
		return string(payloadString), err
	}
	// convert payload to base64
	payload64 := base64.StdEncoding.EncodeToString(payloadString)

	unsignedString := string(header) + string(payloadString)
	// write this unsigned string to the hash created earlier
	h.Write([]byte(unsignedString))

	// now convert this signed string into base64 string

	signedString := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// generate the final token by combining all of base64 strings with dot(.) operator

	token := header64 + "." + payload64 + "." + signedString

	return token, nil
}

// function to validate jwt
func validateToken(token string, secret string) (bool, error) {
	// jwt has 3 parts, seperated by dot(.)
	splitTokens := strings.Split(token, ".")
	// if length is not 3, then token is corrupt
	if len(splitTokens) != 3 {
		return false, nil
	}

	// decode header and payload back to normal strings
	header, err := base64.StdEncoding.DecodeString(splitTokens[0])
	if err != nil {
		return false, err
	}

	payload, err := base64.StdEncoding.DecodeString(splitTokens[1])
	if err != nil {
		return false, err
	}

	// create the signed string again

	unsignedString := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedString))

	signedString := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signedString != splitTokens[2] {
		return false, nil
	}

	// if all checks work, that means it is a valid token
	return true, nil
}

func hashPassword(password string, secret string) (string, error) {
	unhashedPassword := []byte(password)
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write(unhashedPassword)
	if err != nil {
		return "", err
	}
	hashedPassword := hex.EncodeToString(h.Sum(nil))
	return hashedPassword, nil
}

// custom error functions
func InternalServerError(res http.ResponseWriter, message ...string) {
	res.WriteHeader(http.StatusInternalServerError)
	if len(message) == 0 {
		res.Write([]byte("Something went wrong, please try again"))
	} else {
		res.Write([]byte(message[0]))
	}
}

func ConflictError(res http.ResponseWriter, message ...string) {
	res.WriteHeader(http.StatusConflict)
	if len(message) == 0 {
		res.Write([]byte("User already exists!"))
	} else {
		res.Write([]byte(message[0]))
	}
}

func BadRequestError(res http.ResponseWriter, message ...string) {
	res.WriteHeader(http.StatusBadRequest)
	if len(message) == 0 {
		res.Write([]byte("Bad request, please check the data"))
	} else {
		res.Write([]byte(message[0]))
	}
}

// database functions
func CreateNewUser(db *redis.Client, name string, email string, hashedpassword string) error {
	key := fmt.Sprintf("user:%s", email)
	_, err := db.HSet(ctx, key, "name", name, "password", hashedpassword).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *redis.Client, email string, hashedPassword string) (map[string]string, error) {
	key := fmt.Sprintf("user:%s", email)
	result, err := db.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if result["password"] != hashedPassword {
		return nil, fmt.Errorf("invalid credentials")
	}

	return result, nil
}

func IsUserDuplicate(db *redis.Client, email string) (bool, error) {
	key := fmt.Sprintf("user:%s", email)
	result, err := db.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func CreateSession(sessionDB *redis.Client, authToken string) error {
	result, err := sessionDB.Exists(ctx, authToken).Result()
	if err != nil {
		return err
	}
	if result == 1 {
		return fmt.Errorf("User is already logged in")
	}

	_, err = sessionDB.Set(ctx, authToken, true, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func DeleteSession(sessionDB *redis.Client, authToken string) error {
	result, err := sessionDB.Exists(ctx, authToken).Result()
	if err != nil {
		return err
	}
	if result == 0 {
		return fmt.Errorf("User not logged in")
	}

	_, err = sessionDB.Del(ctx, authToken).Result()
	if err != nil {
		return err
	}
	return nil
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
	})
}