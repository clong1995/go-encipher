package json

import (
	"encoding/json"
	"io"
	"log"
)

// Encoder 编码数据为json
// in 要编码的数据的指针
// writer 接收结果的指针
// data := writer.Bytes() 可以将流转 []byte
// 将 w http.ResponseWriter 传入writer，可直接写入http响应流
func Encoder(in any, writer io.Writer) (err error) {
	data, err := json.Marshal(in)
	if err != nil {
		log.Println(err)
		return
	}
	if _, err = writer.Write(data); err != nil {
		log.Println(err)
		return
	}
	return
}
