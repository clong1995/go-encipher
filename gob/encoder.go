package gob

import (
	"encoding/gob"
	"io"
	"log"
)

// Encoder 编码gob
// in 要编码的数据的指针
// writer 接收结果的指针
// data := writer.Bytes() 可以将流转 []byte
// 将 w http.ResponseWriter 传入writer，可直接写入http响应流
func Encoder(in any, writer io.Writer) (err error) {
	encoder := gob.NewEncoder(writer)
	if err = encoder.Encode(in); err != nil {
		log.Println(err)
		return
	}
	return
}
