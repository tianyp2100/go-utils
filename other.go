package tsgutils

import "fmt"

/*
 other utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
  ternary operator replace language: a == b ? c : d
*/
func T3O(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func FmtPrintln(info interface{}) {
	if info != nil {
		switch t := info.(type) {
		case struct{}:
			fmt.Println(Struct2Json(t))
		default:
			fmt.Println(t)
		}
	} else {
		fmt.Println(nil)
	}
}
