package json

import (
	"encoding/json"
	"io"
	"log"
)

// Decode 解码json
// data 二进制数据
// out 接收结果的指针
// bytes.NewBuffer(encoderData) 可以将encoderData二进制转流
// 将 r *http.Request reader,可直接解http请求流
func Decode(reader io.Reader, out any) (err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(data, out); err != nil {
		log.Println(err)
		return
	}
	return
}
