package errors

import (
	"log"
)

func Handler(err error) {
	if err != nil {
		log.Println("[WARN] ", err.Error())
	}
}