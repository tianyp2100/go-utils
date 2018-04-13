package tsgutils

/*
 code utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"github.com/satori/go.uuid"
	"log"
)

/*
  Get a UUID string
  eg: f423fd72-74ba-4277-a7d3-008f2bcb739b
*/
func UUID() string {
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Fatal("UUID Generator: ", err)
	}
	return UUID.String()
}
