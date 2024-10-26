package json

import (
	"testing"
)

func TestDecode(t *testing.T) {
	var encodeData []byte

	//生成编码的测试数据
	err := Encoder(&student{
		Name: "小明",
		Age:  18,
	}, &encodeData)
	if err != nil {
		t.Log(err)
		return
	}

	var s student

	type args struct {
		data []byte
		out  any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "解码json",
			args: args{
				data: encodeData,
				out:  &s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err = Decode(tt.args.data, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Decode() result = %v", s)
		})
	}
}
