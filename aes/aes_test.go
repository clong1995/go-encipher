package aes

import (
	"testing"
)

func TestDecrypt(t *testing.T) {
	/*b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(160862117003087871))
	*/

	type args struct {
		cipherText string
		key        []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "解密测试",
			args: args{
				cipherText: "zhFguxJBpAIkvT1Aac8j/OfFBtmUD+sw5OY1QGdkBo9KHWimY9AQ4PaLCH4OXkNsIK384128zo2ML173TdpARv5pi81h",
				key:        []byte("123456789"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPlainText, err := Decrypt(tt.args.cipherText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotPlainText: %s", gotPlainText)
		})
	}
}

func TestEncrypt(t *testing.T) {
	/*b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(160862117003087872))*/

	type args struct {
		plainText []byte
		key       []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "加密测试",
			args: args{
				plainText: []byte("AFAlWJ1nRwIAoO6V2-aCBA2025-11-08 22:18:39"),
				key:       []byte("123456789"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCipherText, err := Encrypt(tt.args.plainText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotCipherText = %v", gotCipherText)
		})
	}
}
