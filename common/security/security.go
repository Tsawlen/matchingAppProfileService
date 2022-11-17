package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

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
