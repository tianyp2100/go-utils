package tsgutils

/*
 json utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"encoding/json"
)

/**
  Convert struct to json string
  eg:
	type Info struct {
		Field1 string  `json:"field1"`
		Field2 int64   `json:"field2,string"`
		Field3 float64 `json:"field3,string"`
	}

	var info Info
	info.Field1 = "Welcome to time space chain ..."
	info.Field2 = 1688
	info.Field3 = 1314.521

  return:
    {"field1":"Welcome to time space chain ...","field2":"1688","field3":"1314.521"}
*/
func StructToJson(object interface{}) string {
	data, err := json.Marshal(object)
	CheckAndPrintError("Struct convert json string failed", err)
	return string(data)
}

/*
  Convert json string to struct
  eg:
	jsonString := "{\"field1\":\"Welcome to time space chain ...\",\"field2\":\"1688\",\"field3\":\"1314.521\"}"

	var info Info
	Json2Struct(jsonString, &info)
	println(info.Field1)
  print:
    Welcome to time space chain ...
*/
func JsonToStruct(jsonString string, object interface{}) {
	err := json.Unmarshal([]byte(jsonString), object)
	CheckAndPrintError("Json string convert struct failed", err)
}
