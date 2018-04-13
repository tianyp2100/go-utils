package tsgutils

import (
	"bytes"
	"log"
)

/*
  If there is an error print it, pass it
*/
func CheckAndPrintError(flag string, err error) {
	if err != nil {
		log.Fatal(flag+"\n", err)
	}
}

func Print(arg ...interface{}) {
	log.Fatal(arg)
}

func PrintSlowConn(driverName, host, dbName string, consume int64) {
	var buffer bytes.Buffer
	buffer.WriteString(driverName)
	buffer.WriteString("Conn Timeout: ")
	buffer.WriteString("Host: ")
	buffer.WriteString(host)
	buffer.WriteString(", DBName: ")
	buffer.WriteString(dbName)
	buffer.WriteString("Consume time: ")
	buffer.WriteString(Int642String(consume))
	buffer.WriteString("ms")
	log.Fatal(buffer.String())
}

func PrintErrorSql(driverName string, err error, sql string, args ...interface{}) {
	if err != nil {
		log.Fatal(driverName+"Error Sql: ", sql, args)
	}
}

func PrintSlowSql(driverName, host, dbName string, consume int64, sql string, args ...interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString(driverName)
	buffer.WriteString("Slow Sql: ")
	buffer.WriteString("Host: ")
	buffer.WriteString(host)
	buffer.WriteString(", DBName: ")
	buffer.WriteString(dbName)
	buffer.WriteString("Consume time: ")
	buffer.WriteString(Int642String(consume))
	buffer.WriteString("ms, Sql: ")
	buffer.WriteString(sql)
	buffer.WriteString(", Args: ")
	log.Fatal(buffer.String(), args)
}
