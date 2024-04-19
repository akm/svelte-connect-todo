package auth

import "fmt"

const cookieKey = "session"

func SessionCookieValue(idToken string) string {
	// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
	return fmt.Sprintf("%s=%s; HttpOnly", cookieKey, idToken) // Secure flag allows the cookie on HTTPS only
}
