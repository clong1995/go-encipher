package gob

import (
	"bytes"
	"io"
	"testing"
)

func TestDecode(t *testing.T) {

	//先编码一份数据
	var buf bytes.Buffer
	err := Encoder(student{Name: "小明", Age: 18}, &buf)
	if err != nil {
		t.Errorf("Encoder() error = %v", err)
		return
	}
	encoderData := buf.Bytes()

	//测试
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
			name: "gob解码",
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

var s student
