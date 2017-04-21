package jwt

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	signKey, verifyKey []byte
	privRSA            *rsa.PrivateKey
	pubRSA             *rsa.PublicKey
)

func init() {
	signKey, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
	privRSA, _ = jwt.ParseRSAPrivateKeyFromPEM(signKey)

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
	pubRSA, _ = jwt.ParseRSAPublicKeyFromPEM(verifyKey)

}

func GetToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "adminsilca",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(privRSA)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return pubRSA, nil
	})
	return token.Valid
}
