package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/pkg/errors"
)

// DeriveKey insecurely turns a password of any length into a 32-byte AES key (AES-256).
// WARNING: This is not a secure Key Derivation Function (KDF).
// If the password is user-supplied, use a strong KDF like scrypt or Argon2 instead.
func DeriveKey(password []byte) []byte {
	sum := sha256.Sum256(password)
	return sum[:]
}

// Encrypt encrypts data using AES-GCM and returns the nonce prepended to the ciphertext.
func Encrypt(plainIn, password []byte) ([]byte, error) {
	key := DeriveKey(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new cipher")
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new GCM")
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.Wrap(err, "failed to read nonce")
	}

	cipherByte := aead.Seal(nil, nonce, plainIn, nil)
	cipherOut := append(nonce, cipherByte...)
	return cipherOut, nil
}

// Decrypt decrypts data (nonce|ciphertext) that was encrypted with Encrypt.
func Decrypt(encodedIn []byte, password []byte) ([]byte, error) {
	/*data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println(err)
		return
	}*/

	key := DeriveKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new cipher")
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new GCM")
	}

	nonceSize := aead.NonceSize()
	if len(encodedIn) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherText := encodedIn[:nonceSize], encodedIn[nonceSize:]
	plainOut, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open cipher text")
	}

	return plainOut, nil
}
