package gob

import (
	"bytes"
	"io"
	"testing"
)

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
			name: "gob编码",
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
			err := Encoder(tt.args.in, tt.args.writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Encoder() result = %s", buf.Bytes())
		})
	}
}
