package api

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			username, password, ok := r.BasicAuth()
			if ok && username == "test" && password == "test" {
				next.ServeHTTP(w, r)
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		*/
		next.ServeHTTP(w, r)
	}
}

type Credentials struct {
	Username string
	Password []byte
}

const (
	MB = 1024
)

var (
	expirationTime = 2 * time.Hour
)

func Login(config Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Credentials{}
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(403)
			return
		}

		key := argon2.IDKey(c.Password, []byte(config.Salt), 4, 512*MB, 8, 64)

		userCheck := subtle.ConstantTimeCompare([]byte(c.Username), []byte(config.Username))
		pwCheck := subtle.ConstantTimeCompare(key, []byte(config.Password))

		if userCheck == 0 || pwCheck == 0 {
			fmt.Println("Auth error for user", c.Username)
			w.WriteHeader(403)
			return
		}

		token, err := createJWT(config.HMACSecret)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		// user authenticated
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(expirationTime),
		})
		w.WriteHeader(200)

	}
}

func createJWT(hmacSecret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	return token.SignedString(hmacSecret)
}
