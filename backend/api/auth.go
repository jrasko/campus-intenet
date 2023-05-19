package api

import (
	"backend/model"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

func AuthMiddleware(next http.HandlerFunc, config model.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") || len(header) <= 7 {
			fmt.Println("Unauthicated access")
			w.WriteHeader(403)
			return
		}
		header = header[7:]

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
			}
			return []byte(config.HMACSecret), nil
		})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(403)
			return
		}
		if !token.Valid {
			fmt.Println("Invalid token")
			w.WriteHeader(403)
			return
		}
		next.ServeHTTP(w, r)
	}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	MB = 1024
)

var (
	expirationTime = 2 * time.Hour
)

func Login(config model.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Credentials{}
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(403)
			return
		}

		key := base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(c.Password), []byte(config.Salt), 4, 512*MB, 8, 64))

		userCheck := subtle.ConstantTimeCompare([]byte(c.Username), []byte(config.Username))
		pwCheck := subtle.ConstantTimeCompare([]byte(key), []byte(config.Password))

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
		_, _ = w.Write([]byte(fmt.Sprintf(`{ "token": "%s" }`, token)))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)

	}
}

func createJWT(hmacSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
	})
	return token.SignedString([]byte(hmacSecret))
}
