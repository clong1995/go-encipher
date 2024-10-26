package json

import "testing"

func TestEncoder(t *testing.T) {

	var jsonData []byte

	type args struct {
		in   any
		data *[]byte
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
				data: &jsonData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Encoder(tt.args.in, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Encoder() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("Encoder() result = %s", jsonData)
		})
	}
}
