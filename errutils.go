package beeerr

import "log"

func errFatal(msg string, err error) {
	log.Fatalf("%v : %v", msg, err)
}