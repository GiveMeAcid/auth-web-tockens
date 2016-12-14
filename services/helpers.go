package services

import (
	"regexp"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func IsEmailValid(email string) bool {

	const EMAIL_REGEXP = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

	exp, err := regexp.Compile(EMAIL_REGEXP)

	if err != nil {
		panic(err.Error())
	}

	return exp.MatchString(email)
}

// generates hashed password
func ToSha1(plainText string) string {

	hashMaker := sha1.New()

	hashMaker.Write([]byte(plainText))

	hashedPassword := hex.EncodeToString(hashMaker.Sum(nil))

	return hashedPassword
}

// get back data in json format
func ToJSON(w http.ResponseWriter, data interface{}, statusCode int) {

	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(data) // function to convert to json

	w.Write([]byte(res))
}

type responseDataError struct {
	Success      bool `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}

type responseDataSuccess struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func MakeErrorResponse(message string) responseDataError {
	return responseDataError{Success:false, ErrorMessage:message}
}

func MakeSuccessResponse(message string) responseDataSuccess {
	return responseDataSuccess{Success:true, Message:message}
}