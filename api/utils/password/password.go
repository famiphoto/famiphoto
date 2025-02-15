package password

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/famiphoto/famiphoto/api/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, secretKey string) (string, error) {
	dst, err := bcrypt.GenerateFromPassword(hmacHash(password, secretKey), 10)
	if err != nil {
		return "", errors.New(errors.HashPasswordFatal, err)
	}
	return base64.StdEncoding.EncodeToString(dst), nil
}

func MatchPassword(password, correctPassword, secretKey string) (bool, error) {
	decodedCorrect, err := base64.StdEncoding.DecodeString(correctPassword)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(decodedCorrect, hmacHash(password, secretKey))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, errors.New(errors.NoMatchPasswordFatal, err)
}

func hmacHash(src, secretKey string) []byte {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(src))
	return h.Sum(nil)
}
