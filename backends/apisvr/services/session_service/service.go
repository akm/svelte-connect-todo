package sessionservice

import (
	"context"
	"fmt"
	"strings"

	"connectrpc.com/connect"

	v1 "apisvr/gen/session/v1"
	"apisvr/gen/session/v1/sessionv1connect"

	"apisvr/services/auth"
)

type SessionService struct{}

var _ sessionv1connect.SessionServiceHandler = (*SessionService)(nil)

func (s *SessionService) Create(ctx context.Context, req *connect.Request[v1.SessionCreateRequest]) (*connect.Response[v1.Void], error) {
	idToken := strings.ToLower(strings.TrimSpace(req.Msg.IdToken))
	// ID Token validation
	switch idToken {
	case "", "ng", "invalid":
		return nil, fmt.Errorf("Invalid id token")
	}
	res := new(connect.Response[v1.Void])
	res.Header().Add("Set-Cookie", auth.SessionCookieValue(idToken))
	return res, nil
}

func (s *SessionService) Delete(ctx context.Context, req *connect.Request[v1.Void]) (*connect.Response[v1.Void], error) {
	res := new(connect.Response[v1.Void])
	res.Header().Add("Set-Cookie", auth.SessionCookieValue(""))
	return res, nil
}
