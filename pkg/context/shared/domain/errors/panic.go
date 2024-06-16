package errors

import (
	"log"
)

func PanicOnError(where, what string) {
	log.Panicf("(%v): [%v]", where, what)
}
