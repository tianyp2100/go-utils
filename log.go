package tsgutils

import (
	"log"
	"fmt"
)

/*
  If there is an error print it, pass it
*/
func CheckAndPrintError(flag string, err error) {
	if err != nil {
		log.Println(flag+"\n", err)
	}
}

func Print(arg ...interface{}) {
	log.Print(arg)
}

func Println(arg ...interface{}) {
	log.Println(arg)
}

func Stdout(arg ...interface{}) {
	fmt.Println(arg...)
}
