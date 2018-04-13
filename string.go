package tsgutils

/*
 string utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"strconv"
	"strings"
)

func Int2String(arg int) string {
	return strconv.Itoa(arg)
}

func Int642String(arg int64) string {
	return strconv.FormatInt(int64(arg), 10)
}

/*
  If a string contains a string, return true, and ignore case.
  eg: "strings insert chars"
     chars = "insert" -> true
     chars = "Insert" -> true
     chars = "key" -> false
*/
func StringContainsIgnoreCase(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

/*
  If a string contains a string, return true
  eg: "strings insert chars"
     chars = "insert" -> true
     chars = "Insert" -> false
     chars = "key" -> false
*/
func StringContains(s, chars string) bool {
	return strings.Contains(s, chars)
}
