package tsgutils

/*
 time utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"errors"
	"strconv"
	"time"
)

/*
  Get a UTC time string
  eg: 2018-04-08 16:16:06.327540712 +0000 UTC
*/
func UTC() string {
	UTC := time.Now().UTC()
	return UTC.String()
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
func EpochTime() string {
	millisecond := time.Now().UnixNano() / 1000000
	epoch := strconv.Itoa(int(millisecond))
	epochBytes := []byte(epoch)
	epoch = string(epochBytes[:10]) + "." + string(epochBytes[10:])
	return epoch
}

/*
 Get a iso time
  eg: 2018-03-16T18:02:48.284Z
*/
func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
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
	isoBytes := []byte(iso)
	year, err := strconv.Atoi(string(isoBytes[0:4]))
	if err != nil {
		return nilTime, errors.New("illegal year")
	}
	month, err := strconv.Atoi(string(isoBytes[5:7]))
	if err != nil {
		return nilTime, errors.New("illegal month")
	}
	day, err := strconv.Atoi(string(isoBytes[8:10]))
	if err != nil {
		return nilTime, errors.New("illegal day")
	}
	hour, err := strconv.Atoi(string(isoBytes[11:13]))
	if err != nil {
		return nilTime, errors.New("illegal hour")
	}
	min, err := strconv.Atoi(string(isoBytes[14:16]))
	if err != nil {
		return nilTime, errors.New("illegal min")
	}
	sec, err := strconv.Atoi(string(isoBytes[17:19]))
	if err != nil {
		return nilTime, errors.New("illegal sec")
	}
	nsec, err := strconv.Atoi(string(isoBytes[20 : len(isoBytes)-1]))
	if err != nil {
		return nilTime, errors.New("illegal nsec")
	}
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, time.UTC), nil
}
