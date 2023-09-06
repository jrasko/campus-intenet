package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"golang.org/x/crypto/argon2"
)

// CARE! These must match the ones specified in backend/model/config.go to produce a valid hash
const (
	argonKeyLength = 64              // 512 bits
	argonThreads   = 8               // recommended: 2 x server cores
	argonMemory    = 2 * 1024 * 1024 // [in KB] - 2 GiB
	argonTime      = 4
)

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
	hash := base64.StdEncoding.EncodeToString(
		argon2.IDKey(
			[]byte(password),
			[]byte(salt),
			argonTime,
			argonMemory,
			argonThreads,
			argonKeyLength,
		),
	)

	println("Hash:")
	println(hash)
}
