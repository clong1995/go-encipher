package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

// Encrypt 加密
func Encrypt(plainText, key []byte) (cipherText string, err error) {
	key = adjustKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return
	}
	blockSize := block.BlockSize()
	plainText = pkcs7Padding(plainText, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipher_ := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipher_, plainText)

	cipherText = base64.StdEncoding.EncodeToString(cipher_)
	return
}

// Decrypt 解密
func Decrypt(cipherText string, key []byte) (plainText string, err error) {
	key = adjustKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return
	}
	blockSize := block.BlockSize()

	// 解码 Base64
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		log.Println(err)
		return
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	plain := make([]byte, len(data))
	blockMode.CryptBlocks(plain, data)
	plain = pkcs7UnPadding(plain)

	plainText = string(plain)

	return
}

// PKCS7 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 去除 PKCS7 填充
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

func adjustKey(key []byte) []byte {
	if len(key) < 16 {
		return append(key, make([]byte, 16-len(key))...)
	} else if len(key) < 24 {
		return append(key, make([]byte, 24-len(key))...)
	} else if len(key) < 32 {
		return append(key, make([]byte, 32-len(key))...)
	}
	return key[:32]
}

/*func main() {
	// 16 字节的密钥
	key := []byte("1234567890123456")
	plainText := "Hello, Go!"

	fmt.Println("原文:", plainText)

	// 加密
	encrypted, err := aesEncrypt([]byte(plainText), key)
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}
	fmt.Println("加密后的密文:", encrypted)

	// 解密
	decrypted, err := aesDecrypt(encrypted, key)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}
	fmt.Println("解密后的明文:", decrypted)
}*/
