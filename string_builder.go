package tsgutils

/*
 string builder utils
 @author Tony Tian
 @date 2018-04-14
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

func NewStringBuilderString(str *String) *StringBuilder {
	var builder StringBuilder
	builder.buffer.WriteString(str.ToString())
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) Replace(old, new string) *StringBuilder {
	builder.Clear()
	str := NewString(builder.ToString()).Replace(old, new)
	builder.buffer.WriteString(str)
	return builder
}

func (builder *StringBuilder) RemoveLast() *StringBuilder {
	builder.Clear()
	str1 := NewString(builder.ToString())
	str2 := str1.Substring(0, str1.Len()-1)
	builder.buffer.WriteString(str2)
	return builder
}

func (builder *StringBuilder) Clear() *StringBuilder {
	var buffer bytes.Buffer
	builder.buffer = buffer
	return builder
}

func (builder *StringBuilder) ToString() string {
	return builder.buffer.String()
}
