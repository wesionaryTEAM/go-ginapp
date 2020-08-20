package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is ...
type JWTService interface {
	// don't send password... creating token
	CreateToken(Username string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

//CustomClaims is ...
type CustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//JWTAuthService is ...
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Dinesh",
	}
}

func getSecretKey() string {
	secret := os.Getenv("secret_key")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

// token creation:sign claims

// CreateToken is ...
func (s *jwtServices) CreateToken(Username string, isUser bool) string {
	claims := &CustomClaims{
		Username,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 15).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	// pass payloader and header here ... claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// encoded string is ...signed token using secretkey
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// validate token ... use signing method
func (s *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token%v", token.Header["alg"])

		}
		return []byte(s.secretKey), nil
	})

}
