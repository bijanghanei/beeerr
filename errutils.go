package beeerr

import "log"

func ErrFatal(msg string, err error) {
	log.Fatalf("%v : %v", msg, err)
}

func ErrPrint(msg string, err error) {
	log.Printf("%v : %v", msg, err)
}