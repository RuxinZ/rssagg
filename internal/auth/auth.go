package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
GetAPIKey extracts an API key from the header of an HTTP request
Example:
Authorization: ApiKey {insert apikey here}
*/
func GetAPIKey(headers http.Header)(string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header included")
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}
	return splitAuth[1], nil
}