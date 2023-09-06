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

type AuthHandler struct {
	config model.Configuration
}

// Middleware wraps a http.HandlerFunc and checks if the caller is authenticated with a valid token
// if the caller has a valid token, the request will be forwarded to the wrapped http.HandlerFunc
func (a AuthHandler) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read http authorization header
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			log.Println("[INFO] Unauthenticated access")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		header = header[7:]

		// parse and check jwt in header
		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
			}
			return []byte(a.config.HMACSecret), nil
		})
		if err != nil {
			log.Printf("[WARNING] when parsing jwt: %v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if !token.Valid {
			log.Println("[DEBUG] got invalid token")
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

func (a AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode creadentials
		c := Credentials{}
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			log.Printf("[WARNING] when decoding login credentials: %v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// hash given password
		key := base64.StdEncoding.EncodeToString(argon2.IDKey([]byte(c.Password), []byte(a.config.Salt), 4, 512*MB, 8, 64))

		// check if username and password are equal
		userCheck := subtle.ConstantTimeCompare([]byte(c.Username), []byte(a.config.Username))
		pwCheck := subtle.ConstantTimeCompare([]byte(key), []byte(a.config.Password))

		if userCheck == 0 || pwCheck == 0 {
			log.Printf("[DEBUG] auth error for user %s", c.Username)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// create a json web token
		token, err := createJWT(a.config.HMACSecret)
		if err != nil {
			log.Printf("[ERROR] when creating jwt: %v", err)
			sendHttpError(w, err)
			return
		}

		// send token back to user
		_, err = w.Write([]byte(fmt.Sprintf(`{ "token": "%s" }`, token)))
		if err != nil {
			log.Printf("[ERROR] when writing response: %v", err)
			sendHttpError(w, err)
			return
		}
		log.Printf("[DEBUG] new login for user %s", c.Username)
	}
}

func createJWT(hmacSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
	})
	return token.SignedString([]byte(hmacSecret))
}
