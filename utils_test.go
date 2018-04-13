package tsgutils

import "testing"

func TestStruct2Json(t *testing.T) {
	var info Info
	info.Field1 = "Welcome to time space chain ..."
	info.Field2 = 1688
	info.Field3 = 1314.521

	jsonString := Struct2Json(info)
	FmtPrintln(jsonString)
}

func TestJson2Struct(t *testing.T) {
	jsonString := "{\"field1\":\"Welcome to time space chain ...\",\"field2\":\"1688\",\"field3\":\"1314.521\"}"

	var info Info
	Json2Struct(jsonString, &info)
	FmtPrintln(info.Field1)
}