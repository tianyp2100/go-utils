package tsgutils

/*
 time utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"errors"
	"time"
)

const (
	MILLION  = 1000000
)

/*
  Get a UTC time string
  eg: 2018-04-08 16:16:06.327540712 +0000 UTC
*/
func UTC() string {
	return time.Now().UTC().String()
}

/*
  get system current time's nanosecond
*/
func Nanosecond() int64 {
	return time.Now().UnixNano()
}

/*
  get system current time's millisecond
*/
func Millisecond() int64 {
	return Nanosecond() / 1e6
}

/*
  get system current time's second
*/
func Second() int64 {
	return time.Now().Unix()
}

/*
 Get a epoch time
  eg: 1521221737.376
*/
func EpochTime(time time.Time) string {
	return formatEpochTime(time.UnixNano() / MILLION)
}

func EpochTimeNow() string {
	return EpochTime(time.Now())
}

func formatEpochTime(time int64) string {
	str := NewStringInt64(time)
	return str.SubstringEnd(10).AppendString(DOT).Append(str.SubstringBegin(10)).ToString()
}

/*
 Get a iso time
  eg: 2018-03-16T18:02:48.284Z
*/
func IsoTime(time time.Time) string {
	return formatIsoTime(time.UTC().String())
}

func IsoTimeNow() string {
	return IsoTime(time.Now())
}

func formatIsoTime(time string) string {
	str := NewString(time)
	builder := NewStringBuilder()
	builder.Append(str.SubstringEnd(10).ToString()).Append("T")
	builder.Append(str.Substring(11, 23).ToString()).Append("Z")
	return builder.ToString()
}

/*
  iso time change to time.Time
  eg: "2018-11-18T16:51:55.933Z" -> 2018-11-18 16:51:55.000000933 +0000 UTC
*/
func IsoToTime(iso string) (time.Time, error) {
	nilTime := time.Now()
	if iso == "" {
		return nilTime, errors.New("illegal parameter")
	}
	// "2018-03-18T06:51:05.933Z"
	str := NewString(iso)
	year, err := str.Substring(0, 4).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal year")
	}
	month, err := str.Substring(5, 7).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal month")
	}
	day, err := str.Substring(8, 10).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal day")
	}
	hour, err := str.Substring(11, 13).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal hour")
	}
	min, err := str.Substring(14, 16).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal min")
	}
	sec, err := str.Substring(17, 19).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal sec")
	}
	nsec, err := str.Substring(20, str.Len()-1).ToInt()
	if err != nil {
		return nilTime, errors.New("illegal nsec")
	}
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, time.UTC), nil
}

func EpochToTime(epoch string) (time.Time, error) {
	nilTime := time.Now()
	if epoch == "" {
		return nilTime, errors.New("illegal parameter")
	}
	f, err := NewString(epoch).ToFloat()
	if err != nil {
		return nilTime, errors.New("illegal epoch")
	}
	sec := int64(int(f))
	return time.Unix(sec, 0), nil
}

var LAYOUT = map[int]string{
	1: "2006-01-02",
	2: "2006-01-02 15:04",
	3: "2006-01-02 15:04:05",
	4: "2006-01-02 15:04:05.001",
	5: "2006-01-02T15:04:05.001Z",
	6: "1136214245.001",
}

func StringToTime(timeString string, layoutNo int) (time.Time, error) {
	if layoutNo == 5 {
		return IsoToTime(timeString)
	}
	if layoutNo == 6 {
		return EpochToTime(timeString)
	}
	layout := LAYOUT[layoutNo]
	return time.Parse(layout, timeString)
}

func TimeToString(time time.Time, layoutNo int) string {
	if layoutNo == 6 {
		return EpochTime(time)
	}
	layout := LAYOUT[layoutNo]
	return time.Format(layout)
}
