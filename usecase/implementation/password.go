package implementation

import (
	"bytes"
	"golang.org/x/crypto/argon2"
)

func HashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func compPass(passHash []byte, plainPassword string) bool {
	salt := make([]byte, 8)
	copy(salt, passHash[0:8])

	userPassHash := HashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}
