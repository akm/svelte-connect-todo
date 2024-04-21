package auth

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/authn"
)

func Authenticate(_ context.Context, req authn.Request) (any, error) {
	var sessionCookie *http.Cookie
	for _, cookie := range req.Cookies() {
		if cookie.Name == cookieKey {
			sessionCookie = cookie
			break
		}
	}
	if sessionCookie == nil {
		log.Printf("sessionCookie is nil\n")
		return nil, nil
	}

	log.Printf("sessionCookie.Value: %q\n", sessionCookie.Value)
	if sessionCookie.Value == "revoked" {
		return nil, authn.Errorf("revoked session")
	}

	return sessionCookie.Value, nil
}
