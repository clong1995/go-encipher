package json

import (
	"encoding/json"
	"io"
	"log"
)

// Encode 编码数据为json
// in 要编码的数据的指针
// writer 接收结果的指针
// data := writer.Bytes() 可以将流转 []byte
// 将 w http.ResponseWriter 传入writer，可直接写入http响应流
func Encode(in any, writer io.Writer) (err error) {
	encoder := json.NewEncoder(writer)
	if err = encoder.Encode(in); err != nil {
		log.Println(err)
		return
	}
	return
}
