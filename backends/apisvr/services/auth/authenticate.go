package auth

import (
	"applib/log/slog"
	"context"
	"net/http"

	"connectrpc.com/authn"
)

func Authenticate(logger slog.Logger) func(ctx context.Context, req authn.Request) (any, error) {
	return func(ctx context.Context, req authn.Request) (any, error) {
		fbClient, err := NewFirebaseClient(ctx)
		if err != nil {
			logger.Error("initializing firebase client", "cause", err)
		}

		var sessionCookie *http.Cookie
		for _, cookie := range req.Cookies() {
			if cookie.Name == cookieKey {
				sessionCookie = cookie
				break
			}
		}
		if sessionCookie == nil {
			logger.Debug("sessionCookie is nil")
			return nil, nil
		}

		logger.Debug("sessionCookie", "value", sessionCookie.Value)

		token, err := fbClient.VerifySessionCookie(ctx, sessionCookie.Value)
		if err != nil {
			logger.Error("verifying session cookie", "cause", err)
			return nil, authn.Errorf("unauthenticated")
		}

		logger.Debug("verified token", "token", token)

		return token, nil
	}
}
