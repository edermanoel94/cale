package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Content(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.Header().Add(contentType, applicationJson)
	return response(w, body, code)
}

func Marshalled(w http.ResponseWriter, v interface{}, code int) (int, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return Error(w, err, http.StatusBadRequest)
	}
	return Content(w, bytes, code)
}

func Error(w http.ResponseWriter, err error, code int) (int, error) {
	if err == nil {
		return 0, fmt.Errorf("err cannot be null")
	}
	bytes := formatMessageError(err.Error())
	return Content(w, bytes, code)
}

func Redirect(w http.ResponseWriter, body []byte, url string, code int) (int, error) {
	w.Header().Add(location, url)
	return Content(w, body, code)
}

func formatMessageError(message string) []byte {
	return []byte(fmt.Sprintf("{\"message\": \"%s\"}", message))
}

func response(w http.ResponseWriter, body []byte, code int) (int, error) {
	w.WriteHeader(code)
	return w.Write(body)
}

type CustomErrorJson struct {
	Message string
	Path    string
	Code    int
}

func (e CustomErrorJson) Error() string {
	bytes, _ := json.Marshal(&e)
	return string(bytes)
}
