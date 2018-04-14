package tsgutils

import (
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
	builder := NewStringBuilder()
	builder.Append(driverName)
	builder.Append("Conn Timeout: ")
	builder.Append("Host: ")
	builder.Append(host)
	builder.Append(", DBName: ")
	builder.Append(dbName)
	builder.Append("Consume time: ")
	builder.Append(NewStringInt64(consume).ToString())
	builder.Append("ms")
	log.Fatal(builder.ToString())
}

func PrintErrorSql(driverName string, err error, sql string, args ...interface{}) {
	if err != nil {
		log.Fatal(driverName+"Error Sql: ", sql, args)
	}
}

func PrintSlowSql(driverName, host, dbName string, consume int64, sql string, args ...interface{}) {
	builder := NewStringBuilder()
	builder.Append(driverName)
	builder.Append("Slow Sql: ")
	builder.Append("Host: ")
	builder.Append(host)
	builder.Append(", DBName: ")
	builder.Append(dbName)
	builder.Append("Consume time: ")
	builder.Append(NewStringInt64(consume).ToString())
	builder.Append("ms, Sql: ")
	builder.Append(sql)
	builder.Append(", Args: ")
	log.Fatal(builder.ToString(), args)
}
