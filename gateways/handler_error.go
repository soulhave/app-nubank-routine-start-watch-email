package gateways

import (
	"fmt"
	"log"
	"net/http"
)

func HandleError(phase string, err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Fatal(err)
		fmt.Fprint(w, "{error:\"", err, "\"}")
		return true
	}
	return false
}
