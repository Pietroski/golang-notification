package errors

import "log"

func (e *Errors) CheckFatalError(err error, message ...string) {
	if err != nil {
		if message != nil {
			var messageToLog string = message[0]
			if messageToLog != "" {
				log.Fatalln(message)
			}
		}
		log.Fatalln("Panic Error!! ->", err)
	}
}
