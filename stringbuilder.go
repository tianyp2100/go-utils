package tsgutils

/*
 string builder utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"bytes"
)

/*
  StringBuilder struct.
	  Usage:
		builder := NewStringBuilder()
		builder.Append("a").Append("b").Append("c")
		s := builder.String()
		println(s)
	  print:
		abc
 */
type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	var builder StringBuilder
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}
func (builder *StringBuilder) ToString() string {
	return builder.buffer.String()
}
