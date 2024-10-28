package json

import (
	"bytes"
	"io"
	"log"
	"testing"
)

func init() {
	log.SetFlags(8)
}

func TestEncoder(t *testing.T) {

	var buf bytes.Buffer

	type args struct {
		in     any
		writer io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "编码数据为json",
			args: args{
				in: &student{
					Name: "小明",
					Age:  18,
				},
				writer: &buf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Encode(tt.args.in, tt.args.writer); (err != nil) != tt.wantErr {
				t.Errorf("Encoder() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Encoder() result = %s", buf.Bytes())
		})
	}
}
