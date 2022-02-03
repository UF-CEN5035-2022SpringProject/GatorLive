package test

import (
	"fmt"
	"net/http"
)

func EchoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GatorStore Backend is alive")
}
