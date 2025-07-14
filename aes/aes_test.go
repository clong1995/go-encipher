package aes

import "testing"

func TestDecrypt(t *testing.T) {
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
				cipherText: "u3OHOgpzWdaZOiy+yriuNw==",
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
				plainText: []byte("hello world"),
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
