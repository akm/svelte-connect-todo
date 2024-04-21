package auth

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/authn"
)

func Authenticate(_ context.Context, req authn.Request) (any, error) {
	var sessionCookie *http.Cookie
	log.Printf("req.Cookies(): %+v\n", req.Cookies())
	for _, cookie := range req.Cookies() {
		log.Printf("cookie: %+v\n", *cookie)
		if cookie.Name == cookieKey {
			sessionCookie = cookie
			break
		}
	}
	if sessionCookie == nil {
		log.Printf("sessionCookie is nil\n")
		return nil, nil
	}

	log.Printf("sessionCookie: %+v\n", *sessionCookie)
	log.Printf("sessionCookie.Value: %q\n", sessionCookie.Value)
	if sessionCookie.Value == "revoked" {
		return nil, authn.Errorf("revoked session")
	}

	return sessionCookie.Value, nil
}
