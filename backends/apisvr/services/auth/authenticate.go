package auth

import (
	"applib/log/slog"
	"context"
	"net/http"

	"connectrpc.com/authn"
	"firebase.google.com/go/v4/auth"
	"github.com/akm/slogwrap"
)

func Authenticate(logger *slog.Logger) func(ctx context.Context, req authn.Request) (any, error) {
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

		// この結果は auth.SetInfo されて ctx に保存される
		return token, nil
	}
}

func init() {
	slogwrap.Register(
		func(orig slogwrap.HandleFunc) slogwrap.HandleFunc {
			return func(ctx context.Context, rec slog.Record) error {
				// Authenticate の戻り値の関数の戻り値の token を取得
				token, ok := authn.GetInfo(ctx).(*auth.Token)
				if ok {
					rec.Add("auth.UID", token.UID)
				}
				return orig(ctx, rec)
			}
		},
	)
}
