package auth

import (
	"context"
	"fmt"

	"connectrpc.com/authn"
)

func Authenticate(_ context.Context, req authn.Request) (any, error) {
	cookie, err := req.Cookie(cookieKey)
	if err != nil {
		return nil, err
	}
	fmt.Printf("cookie: %+v\n", *cookie)
	fmt.Printf("cookie.Value: %q\n", cookie.Value)

	if cookie.Value == "revoked" {
		return nil, authn.Errorf("revoked session")
	}

	return cookie.Value, nil
}
