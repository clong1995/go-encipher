package json

import (
	"encoding/json"
	"log"
)

// Decode 解码josn
// data 二进制数据
// out 接收结果的指针
func Decode(data []byte, out any) (err error) {
	if err = json.Unmarshal(data, out); err != nil {
		log.Println(err)
		return
	}
	return
}
