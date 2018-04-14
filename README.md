# Go utils

##### Some of frequently used utils.

---------------------------------------

### intro
|file|function
|---|---
|code |  GUID, UUID
|json |  StructToJson, JsonToStruct
|log |  CheckAndPrintError
|signer |  HmacSha256Base64Signer, Md5Signer
|string |  Trim, Replace, Contains, ToArray, ToInt, ToInt64, ToFloat, Substring, Append, StringBuilder
|time |  Millisecond, EpochTime, IsoTime, UTC, StringToTime, TimeToString
|other |  IIIOperator, symbol, test

### Installation
```
$ go get -u github.com/timespacegroup/go-utils
```

### Usage

##### More info see: utils_test.go

```
import (
	tsgutils "github.com/timespacegroup/go-utils"
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
