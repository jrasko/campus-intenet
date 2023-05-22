package api

import (
	"backend/model"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

func AuthMiddleware(next http.HandlerFunc, config model.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read http authorization header
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") || len(header) <= 7 {
			log.Println("Unauthicated access")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		header = header[7:]

		// parse and check jwt in header
		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
			}
			return []byte(config.HMACSecret), nil
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if !token.Valid {
			log.Println("Invalid token")
			w.WriteHeader(http.StatusForbidden)
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
		// decode creadentials
		c := Credentials{}
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// hash given password
		key := base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(c.Password), []byte(config.Salt), 4, 512*MB, 8, 64))

		// check if username and password are equal
		userCheck := subtle.ConstantTimeCompare([]byte(c.Username), []byte(config.Username))
		pwCheck := subtle.ConstantTimeCompare([]byte(key), []byte(config.Password))

		if userCheck == 0 || pwCheck == 0 {
			log.Println("Auth error for user", c.Username)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// create a jwt
		token, err := createJWT(config.HMACSecret)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		// send token back to user
		_, _ = w.Write([]byte(fmt.Sprintf(`{ "token": "%s" }`, token)))
		if err != nil {
			sendHttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)

	}
}

func createJWT(hmacSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
	})
	return token.SignedString([]byte(hmacSecret))
}
