package diffie // import "diffie"

import (
	"crypto"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/aead/ecdh"
)

var cipher = ecdh.X25519()

// DiffieHellman : represents a diffie hellman cipher
type DiffieHellman struct {
	PrivateKey crypto.PrivateKey
}

// GenPublicKey : generates a public key
func (d *DiffieHellman) GenPublicKey() (string, error) {
	var reader io.Reader

	privateKey, publicKey, err := cipher.GenerateKey(reader)
	publicKeyJSON, _ := json.Marshal(publicKey)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKeyJSON)

	d.PrivateKey = privateKey

	return publicKeyBase64, err
}

// GenSharedKey : generates a shared key
func (d *DiffieHellman) GenSharedKey(publicKeyBase64 string) string {
	var publicKey []byte

	publicKeyJSON, err := base64.StdEncoding.DecodeString(publicKeyBase64)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(publicKeyJSON, &publicKey)

	sharedKey := cipher.ComputeSecret(d.PrivateKey, crypto.PublicKey(publicKey))

	return fmt.Sprintf("%x", md5.Sum(sharedKey))
}
