package backend

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
)

var (
	secretKey, _ = hex.DecodeString("asasdksd349034ldfoiwe234nkj")
	cookieName   = "adminCookie"
	cookieValue  = "Hello Admin!"
)

func handleSetCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Path:     "/",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
	}

	err := handleWriteSigned(w, cookie)
	if err != nil {
		fmt.Println("[ERROR] Failed to set cookie:", err)
		http.Error(w, fmt.Sprintf("Failed to set cookie: %v", err), http.StatusInternalServerError)
		return
	}
}

func handleGetCookie(r *http.Request) error {
	_, err := handleReadSigned(r)
	if err != nil {
		return err
	}
	return nil
}

func handleWriteSigned(w http.ResponseWriter, cookie http.Cookie) error {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)

	cookie.Value = string(signature) + cookie.Value

	return handleWriteBase64(w, cookie)
}

func handleWriteBase64(w http.ResponseWriter, cookie http.Cookie) error {
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	if len(cookie.Value) > 4096 {
		return errors.New("cookie value exceeds maximum length")
	}

	http.SetCookie(w, &cookie)
	return nil
}

func handleReadSigned(r *http.Request) (string, error) {
	signedValue, err := handleReadBase64(r)
	if err != nil {
		return "", err
	}

	if len(signedValue) < sha256.Size {
		return "", errors.New("cookie value is too short")
	}

	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookieName))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", errors.New("invalid cookie signature")
	}

	return value, nil
}

func handleReadBase64(r *http.Request) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", err
	}

	return string(value), nil
}
