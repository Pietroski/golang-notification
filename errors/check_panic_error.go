package errors

import (
	"log"
)

func (e *Errors) CheckPanicError(err error, message ...string) {
	if err != nil {
		if message != nil {
			var messageToLog string = message[0]
			if messageToLog != "" {
				log.Fatalln(message)
			}
		}
		log.Fatalln("Fatal Error!! ->", err)
	}
}
