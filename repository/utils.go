package repository

import "log"

func ExitIfError(e error, msg string) {
	if e != nil {
		log.Fatal(msg)
	}
}
