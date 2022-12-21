package security

import (
	"app/matchingAppProfileService/common/dataStructures"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func GenerateJWT(requestedUser *dataStructures.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":  requestedUser.ID,
		"exp":  time.Now().Add(3 * time.Hour).Unix(),
		"user": requestedUser.ID,
	})

	// Get RSA private key

	key, errGetKey := GetPrivateToken()

	if errGetKey != nil {
		return "", errGetKey
	}

	// Sign the token
	tokenString, errSign := token.SignedString(key)

	if errSign != nil {
		return "", errSign
	}

	return tokenString, nil
}

func GetPrivateToken() (*rsa.PrivateKey, error) {

	if privateKey == nil {
		fmt.Println("Getting token...")
		key, err := getPrivateTokenFromEnvironment()
		if err != nil {
			return nil, err
		}
		return key, nil
	}

	return privateKey, nil

}

func GetPublicToken() (*rsa.PublicKey, error) {
	if publicKey == nil {
		key, err := getPublicTokenFromEnvironment()
		if err != nil {
			return nil, err
		}
		return key, nil
	}
	return publicKey, nil
}

func getPrivateTokenFromEnvironment() (*rsa.PrivateKey, error) {
	envToken := os.Getenv("PRIVATE_SECRET")

	block, _ := pem.Decode([]byte(envToken))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Could not extract private key!")
	}
	return key, nil
}

func getPublicTokenFromEnvironment() (*rsa.PublicKey, error) {
	envToken := os.Getenv("PUBLIC_SECRET")

	block, _ := pem.Decode([]byte(envToken))
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Could not extract public key!")
	}
	var keyReturn = key.(*rsa.PublicKey)
	return keyReturn, nil
}
