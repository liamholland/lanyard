package logging

import "log"

func LogControllerRequest(controller string, method string, from string) {
	log.Printf("received %s request to '%s' from %s", method, controller, from)
}
