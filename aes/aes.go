package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

// DeriveKey 将任意长度密码转为 32 字节 AES 密钥（AES-256）
// 若密码是用户输入的口令，请在上层结合 scrypt/argon2 做强 KDF。
func DeriveKey(password []byte) []byte {
	sum := sha256.Sum256(password)
	return sum[:]
}

// Encrypt 使用 AES-GCM 加密，返回 base64(nonce|ciphertext)
func Encrypt(plain, password []byte) (cipherText string, err error) {
	key := DeriveKey(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
		return
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err)
		return
	}

	cipherByte := aead.Seal(nil, nonce, plain, nil)
	out := append(nonce, cipherByte...)
	cipherText = base64.StdEncoding.EncodeToString(out)
	return
}

// Decrypt 解密 base64(nonce|ciphertext)，返回明文
func Decrypt(encoded string, password []byte) (plainText string, err error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println(err)
		return
	}

	key := DeriveKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err)
		return
	}

	nonceSize := aead.NonceSize()
	if len(data) < nonceSize {
		err = errors.New("ciphertext too short")
		log.Println(err)
		return
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plain, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Println(err)
		return
	}
	plainText = string(plain)

	return
}
