package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
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
func generateToken(header string, payload map[string]string, secret string) (string, error) {
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

	unsignedString := header + string(payloadString)
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
