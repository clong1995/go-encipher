package gob

import (
	"encoding/gob"
	"io"
	"log"
)

// Decode 解码gob
// reader 要解码的数据流,
// out 接收结果的指针
// bytes.NewBuffer(encoderData) 可以将encoderData二进制转流
// 将 r *http.Request reader,可直接解http请求流
func Decode(reader io.Reader, out any) (err error) {
	decoder := gob.NewDecoder(reader)
	if err = decoder.Decode(out); err != nil {
		log.Println(err)
		return
	}
	return
}
