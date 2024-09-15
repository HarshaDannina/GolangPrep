package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("no authorization header found")
	}
	vals := strings.Split(auth, " ")
	if len(vals) < 2 {
		return "", fmt.Errorf("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", fmt.Errorf("malformed first part of auth header")
	}
	return vals[1], nil
}
