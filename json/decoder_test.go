package json

import (
	"bytes"
	"io"
	"testing"
)

func TestDecode(t *testing.T) {
	var encoderBuf bytes.Buffer

	//生成编码的测试数据
	err := Encoder(
		&student{Name: "小明", Age: 18},
		&encoderBuf,
	)
	if err != nil {
		t.Errorf("Encoder() error = %v", err)
		return
	}

	encoderData := encoderBuf.Bytes()

	var s student

	type args struct {
		reader io.Reader
		out    any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "解码json",
			args: args{
				reader: bytes.NewBuffer(encoderData),
				out:    &s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err = Decode(tt.args.reader, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Decode() result = %v", s)
		})
	}
}
