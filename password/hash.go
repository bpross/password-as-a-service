package password

import (
	"crypto/sha512"
	"encoding/base64"
)

type Password struct {
	hashedPassword []byte
}

func CreateAndHash(password string) *Password {
	hasher := sha512.New()
	ba := []byte(password)
	hasher.Write(ba)
	sha := hasher.Sum(nil)
	p := &Password{
		hashedPassword: sha,
	}
	return p
}

func (p *Password) UrlEncode() string {
	urlEncoded := base64.StdEncoding.EncodeToString(p.hashedPassword)
	return urlEncoded
}
