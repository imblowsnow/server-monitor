package utils

import "encoding/json"

func CopyObject(src interface{}, dest interface{}) {
	// 使用json序列化和反序列化实现对象的深拷贝
	// 1. 将src对象序列化为json字符串
	jsonStr, _ := json.Marshal(src)

	// 2. 将json字符串反序列化为dest对象
	json.Unmarshal(jsonStr, dest)
}
