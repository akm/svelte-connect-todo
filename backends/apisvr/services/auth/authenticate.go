package auth

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/authn"
)

func Authenticate(ctx context.Context, req authn.Request) (any, error) {
	fbClient, err := NewFirebaseClient(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

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

	token, err := fbClient.VerifySessionCookie(ctx, sessionCookie.Value)
	if err != nil {
		log.Printf("error verifying session cookie: %v\n", err)
		return nil, authn.Errorf("unauthenticated")
	}

	log.Printf("token: %+v\n", *token)

	return token, nil
}
