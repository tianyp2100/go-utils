# Go utils

##### Some of frequently used utils.

---------------------------------------

### intro
|file|function
|---|---
|id_generator |  GUID, UUID
|json |  StructToJson, JsonToStruct
|log |  CheckAndPrintError
|signer |  HmacSha256Base64Signer, Md5Signer
|string |  Trim, Replace, Contains, ToArray, ToInt, ToInt64, ToFloat, Substring, Append, StringBuilder, Clear, Remove,RemoveLast
|string_builder| StringBuilder,Append,Replace,RemoveLast
|interface_builder|InterfaceBuilder, Append, Clear
|time |  Millisecond, EpochTime, IsoTime, UTC, StringToTime, TimeToString
|other,utils_test |  IIIOperator, symbol, test

### Installation
```
$ go get -u github.com/typa01/go-utils
```

### Usage

##### More info see: utils_test.go

```
import (
	tsgutils "github.com/typa01/go-utils"
)
```
```
	str1 := tsgutils.NewString("13990521234")
	str2 := tsgutils.NewString("14")
	str3 := str1.Substring(0, 2).Append(str2).AppendString("520")
	println(str1.ToString())
	println(str2.ToString())
	println(str3.ToString())
```
```
	builder := NewStringBuilder()
	builder.Append("a").Append("b").Append("c")
	str := builder.ToString()
```
```
func (user *User) Rows2struct(rows *mysql.Rows) {
	var users []User
	builder := tsgutils.NewInterfaceBuilder()
	for rows.Next() {
		builder.Clear()
		builder.Append(&user.Host).Append(&user.User)
		builder.Append(&user.AuthenticationString)
		err := rows.Scan(builder.ToInterfaces()...)
		tsgutils.CheckAndPrintError("MySQL query rows scan error", err)
		users = append(users, *user)
	}
	if rows != nil {
		defer rows.Close()
	}
	user.Users = users
}
```
### The future will continue to be updated.
##### https://blog.csdn.net/typa01_kk/article/category/7629914
