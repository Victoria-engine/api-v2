package jwtauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var minSecretLen = 128

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwtauth token is valid.
	ttl time.Duration

	// JWT signing algorithm
	alg jwt.SigningMethod
}


// New : Generates new JWT service necessary for authmiddleware middleware
func New(alg, secret string, ttlMinutes time.Duration, minSecretLength int) (Service, error) {
	if minSecretLength > 0 {
		minSecretLen = minSecretLength
	}

	if len(secret) < minSecretLen {
		return Service{}, fmt.Errorf("jwtauth secret length is %v, which is less than required %v", len(secret), minSecretLen)
	}

	signingMethod := jwt.GetSigningMethod(alg)
	if signingMethod == nil {
		return Service{}, fmt.Errorf("invalid jwtauth signing method: %s", alg)
	}

	return Service{
		key:  []byte(secret),
		alg: signingMethod,
		ttl:  ttlMinutes * time.Minute,
	}, nil
}


// GenerateToken : Creates a jwtauth token for a session
func (s Service) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(s.ttl).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

// IsTokenValid : Checks if the token if valid
func (s Service) IsTokenValid(r *http.Request) error {
	tokenString := s.ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		pretty(claims)
	}

	return nil
}

// ExtractToken : Extracts the token from the request headers or the url
func (s Service) ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("access_token")

	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

// ExtractTokenID : Extracts the "user_id" from the token
func (s Service) ExtractTokenID(r *http.Request) (uint, error) {

	tokenString := s.ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}

// Pretty : Displays the claims nicely in the terminal
func pretty(data interface{}) {
	_, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
}
