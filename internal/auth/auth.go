package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extract an API Key from the header of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	headerAuth := headers.Get("Authorization")
	if headerAuth == "" {
		return "", errors.New("no authentication info found")
	}

	headerAuths := strings.Split(headerAuth, " ")
	if len(headerAuths) != 2 { //need to have ApiKey and the actualy apikey
		return "", errors.New("malformed auth header")
	}
	if headerAuths[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return headerAuths[1], nil
}
