package codec

import (
	"math/rand/v2"
	"testing"
)

func Test_shuffleCodec(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "shuffleCodec",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//chars := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
			chars := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
			rand.Shuffle(len(chars), func(i, j int) {
				chars[i], chars[j] = chars[j], chars[i]
			})
			t.Logf("%s", chars)
		})
	}
}

func Test_codecEncode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "codecEncode",
			args: args{data: []byte("hello world!")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecEncode(tt.args.data)
			if err != nil {
				t.Errorf("codecEncode() error = %v", err)
				return
			}
			t.Logf("codecEncode() = %v", got)
		})
	}
}

func Test_codecDecode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "codecDecode",
			args: args{str: "V3Z8penTHITFoL7FEhTtbb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := codecDecode(tt.args.str)
			if err != nil {
				t.Errorf("codecDecode() error = %v", err)
				return
			}
			t.Logf("codecDecode() got = %v,%s", got, got)
		})
	}
}
