package tsgutils

import (
	"testing"
	"time"
	"fmt"
)

func TestGUID(t *testing.T) {
	FmtPrintln("GUID: " + GUID())
}

func TestUUID(t *testing.T) {
	FmtPrintln("UUID: " + UUID())
}

func TestStructToJson(t *testing.T) {
	var info Info
	info.Field1 = "Welcome to time space chain ..."
	info.Field2 = 1688
	info.Field3 = 1314.521

	jsonString := StructToJson(info)
	FmtPrintln(jsonString)
}

func TestJsonToStruct(t *testing.T) {
	jsonString := "{\"field1\":\"Welcome to time space chain ...\",\"field2\":\"1688\",\"field3\":\"1314.521\"}"

	var info Info
	JsonToStruct(jsonString, &info)
	FmtPrintln("Field1: " + info.Field1)
}

func TestIIIOperator(t *testing.T) {
	result1 := IIIInterfaceOperator(1 > 2, 3, 4)
	result2 := IIIInterfaceOperator(1 > 2, "1", "2")
	result3 := IIIInterfaceOperator(1 > 2, "(?,", "(")
	FmtPrintln(result1)
	FmtPrintln(result2)
	FmtPrintln(result3)
}

func TestHmacSha256Base64Signer(t *testing.T) {
	timestamp := "2018-03-08T10:59:25.789Z"
	method := "POST"
	requestPath := "/orders?before=2&limit=30"
	body := "{\"product_id\":\"BTC-USD-0309\",\"order_id\":\"377454671037440\"}"
	preHashString := PreHashString(timestamp, method, requestPath, body)
	FmtPrintln("preHashString: " + preHashString)
	secretKey := "E65791902180E9EF4510DB6A77F6EBAE"
	sign, err := HmacSha256Base64Signer(preHashString, secretKey)
	CheckAndPrintError("Using hmac sha256 base64 signing a message failed", err)
	FmtPrintln("TestHmacSha256Base64Signer: " + sign)
}

func TestMd5Signer(t *testing.T) {
	s := Md5Signer("123")
	FmtPrintln(s)
}

func TestToStringArray(t *testing.T) {
	s := "460364431014955c2483ec91230e5435"
	FmtPrintln(NewString(s).ToArray())
}

func TestString(t *testing.T) {
	str := NewString("abcde")
	FmtPrintln(str.Substring(0, 2).ToString())
	FmtPrintln(str.SubstringBegin(2).ToString())
	FmtPrintln(str.SubstringEnd(3).ToString())
	FmtPrintln(str.Append(str).ToString())
	FmtPrintln(str.AppendString("123").ToString())
	FmtPrintln(str.Len())
	FmtPrintln(NewString(" 1 23 ").Trim().ToString())
	FmtPrintln(NewString("%111%abc%987%").Replace("%", "$").ToString())
	FmtPrintln(NewString("123xxxbbb5990").StartsWith("123x"))
	FmtPrintln(NewString("123xxxbbb5990").EndsWith("5990"))
	FmtPrintln(NewString("aaa").ToUpper().ToString())
	FmtPrintln(NewString("BBB").ToLower().ToString())
	FmtPrintln(NewString("abcdef").Index("b"))
	FmtPrintln(NewString("abcdef").LastIndex("e"))
}

func TestNewStringFloat64(t *testing.T) {
	str := NewStringFloat64(3.000123456789)
	FmtPrintln(str.ToString()) // 3.000123456789E+00
	f, err := str.ToFloat()
	CheckAndPrintError("String to float failed", err)
	FmtPrintln(f) //3.000123456789
}

func TestStringBuilder(t *testing.T) {
	builder := NewStringBuilder()
	builder.Append("a").Append("b").Append("c")
	s := builder.ToString()
	FmtPrintln("StringBuilder Append: " + s)
}

func TestIsoTime(t *testing.T) {
	iso := IsoTimeNow()
	FmtPrintln("Iso time: " + iso)
}

func TestIsoToTime(t *testing.T) {
	s := "2018-11-18T16:51:55.933Z"
	time, err := IsoToTime(s)
	CheckAndPrintError("Convert iso string to time failed", err)
	FmtPrintln(s + " -> " + time.String())
}

func TestEpochTime(t *testing.T) {
	time := EpochTimeNow()
	FmtPrintln("Epoch time:" + time)
}

func TestEpochToTime(t *testing.T) {
	epoch := "1523702966.293"
	time, err := EpochToTime(epoch)
	CheckAndPrintError("Convert epoch string to time failed", err)
	FmtPrintln(time)
}

func TestTimeToString(t *testing.T) {
	time := time.Now()
	timeStr := TimeToString(time, 5)
	FmtPrintln(timeStr)
}

func TestStringToTime(t *testing.T) {
	timeStr := "2018-03-16T18:02:48.284Z"
	time, err := StringToTime(timeStr, 5)
	CheckAndPrintError("Convert string to time failed", err)
	FmtPrintln(time)
}

func TestInterfaceBuilder(t *testing.T) {
	var s1 string = "Abc"
	var i1 int = 123
	var f1 float64 = 123.40
	builder1 := NewInterfaceBuilder()
	builder1.Append(s1).Append(i1).Append(f1)
	fmt.Println(builder1.ToInterfaces())
	builder2 := builder1.Clear()
	builder2.Appends(s1, i1, f1)
	fmt.Println(builder2.ToInterfaces())
}

func TestStringRemove(t *testing.T) {
	str := "abc"
	FmtPrintln(NewString(str).Remove(1).ToString())
	FmtPrintln(NewString(str).RemoveLast().ToString())
}

func TestStringBuilderReplace(t *testing.T) {
	builder := NewStringBuilder()
	builder.Append("%111%abc%987%")
	FmtPrintln(builder.Replace("%", "$").ToString())
	FmtPrintln(builder.RemoveLast().ToString())
}