package keys

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var (
	verifyKey      *rsa.PublicKey
	signKey        *rsa.PrivateKey
	privateKeyPath string = "keys/app.rsa"
	publicKeyPath  string = "keys/app.rsa.pub"
)

// InitPublicKey return *rsa.PublicKey
func InitPublicKey() (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}

// InitPrivateKey return *rsa.PrivateKey
func InitPrivateKey() (*rsa.PrivateKey, error) {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	return signKey, nil
}

// SetKeyPath set Path for private and public key
func SetKeyPath(privatePath, publicPath string) {
	privateKeyPath = privatePath
	publicKeyPath = publicPath
}
