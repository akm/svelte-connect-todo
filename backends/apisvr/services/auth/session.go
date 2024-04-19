package auth

import "fmt"

func SessionCookieValue(idToken string) string {
	// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
	return fmt.Sprintf("session=%s; Secure; HttpOnly", idToken)
}
