package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"golang.org/x/crypto/argon2"
)

const MB = 1024

func main() {
	if len(os.Args) < 2 {
		panic("no password given")
	}
	password := os.Args[1]

	saltBytes := make([]byte, 24)
	_, _ = rand.Read(saltBytes)
	salt := base64.StdEncoding.EncodeToString(saltBytes)
	println("Salt:")
	println(salt)

	// these settings must be synchronized with backend/api/auth.go
	hash := base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(password), []byte(salt), 4, 512*MB, 8, 64))

	println("Hash:")
	println(hash)
}
