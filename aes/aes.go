package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/pkg/errors"
)

// Encipher 持有用于加密和解密的 AEAD 实例。
type Encipher struct {
	aead cipher.AEAD
}

// NewEncipher 使用给定的密码创建一个新的 Encipher 实例。
func NewEncipher(password []byte) (*Encipher, error) {
	key := deriveKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Encipher{aead: aead}, nil
}

// Encrypt 使用预先计算的 AEAD 实例加密数据。
func (e *Encipher) Encrypt(plainIn []byte) ([]byte, error) {
	nonce := make([]byte, e.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.WithStack(err)
	}
	cipherByte := e.aead.Seal(nil, nonce, plainIn, nil)
	return append(nonce, cipherByte...), nil
}

// Decrypt 使用预先计算的 AEAD 实例解密数据。
func (e *Encipher) Decrypt(encodedIn []byte) ([]byte, error) {
	nonceSize := e.aead.NonceSize()
	if len(encodedIn) < nonceSize {
		return nil, errors.New("密文太短")
	}
	nonce, cipherText := encodedIn[:nonceSize], encodedIn[nonceSize:]
	plainOut, err := e.aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return plainOut, nil
}

// deriveKey 将任意长度的密码不安全地转换为 32 字节的 AES 密钥 (AES-256)。
// 警告：这不是一个安全的密钥派生函数 (KDF)。
// 如果密码由用户提供，请改用像 scrypt 或 Argon2 这样的强 KDF。
func deriveKey(password []byte) []byte {
	sum := sha256.Sum256(password)
	return sum[:]
}
