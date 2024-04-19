package auth

import "fmt"

const cookieKey = "session"

func SessionCookieValue(idToken string) string {
	// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
	return fmt.Sprintf("%s=%s; Secure; HttpOnly", cookieKey, idToken)
}
