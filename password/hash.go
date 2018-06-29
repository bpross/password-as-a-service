package password

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashPassword512(password string) string {
	hasher := sha512.New()
	ba := []byte(password)
	hasher.Write(ba)
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
