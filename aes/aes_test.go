package aes

import (
	"bytes"
	"testing"
)

func TestEncipher_Success(t *testing.T) {
	password := []byte("123456789")
	encipher, err := NewEncipher(password)
	if err != nil {
		t.Fatalf("Failed to create encipher: %v", err)
	}

	plainText := []byte("AFAlWJ1nRwIAoO6V2-aCBA2025-11-08 22:18:39")

	// Encrypt
	cipherText, err := encipher.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decrypt
	decryptedText, err := encipher.Decrypt(cipherText)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Verify
	if !bytes.Equal(plainText, decryptedText) {
		t.Errorf("Decrypted text does not match original text. got %q, want %q", decryptedText, plainText)
	}
}

func TestEncipher_DecryptFailsWithWrongPassword(t *testing.T) {
	correctPassword := []byte("correct-password")
	wrongPassword := []byte("wrong-password")

	// Create encipher with the correct password
	encipher, err := NewEncipher(correctPassword)
	if err != nil {
		t.Fatalf("Failed to create encipher: %v", err)
	}

	plainText := []byte("this is a secret message")

	// Encrypt with the correct password
	cipherText, err := encipher.Encrypt(plainText)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Create a new encipher with the wrong password
	wrongEncipher, err := NewEncipher(wrongPassword)
	if err != nil {
		t.Fatalf("Failed to create encipher with wrong password: %v", err)
	}

	// Try to decrypt with the wrong password, expecting an error
	_, err = wrongEncipher.Decrypt(cipherText)
	if err == nil {
		t.Error("Expected decryption to fail with wrong password, but it succeeded.")
	}
}
