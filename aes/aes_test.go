package aes

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestDecrypt(t *testing.T) {
	/*b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(160862117003087871))
	*/

	encodedIn, err := base64.StdEncoding.DecodeString("+dydQ7XR5CpHvT2K1tv/aQCi4Ver0vq/4OHmuhHObGWrmFmv2gfaokKq2ZQHhoqy5nsLVukolZWFONrsOttRhYb90xtW")
	if err != nil {
		log.Println(err)
		return
	}

	type args struct {
		encodedIn []byte
		password  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "解密测试",
			args: args{
				encodedIn: encodedIn,
				password:  []byte("123456789"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var plainOut []byte
			if plainOut, err = Decrypt(tt.args.encodedIn, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotPlainText := string(plainOut)
			t.Logf("gotPlainText: %s", gotPlainText)
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		plainIn  []byte
		password []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "加密测试",
			args: args{
				plainIn:  []byte("AFAlWJ1nRwIAoO6V2-aCBA2025-11-08 22:18:39"),
				password: []byte("123456789"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipherOut, err := Encrypt(tt.args.plainIn, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotCipherText := base64.StdEncoding.EncodeToString(cipherOut)
			t.Logf("gotCipherText = %v", gotCipherText)
		})
	}
}
