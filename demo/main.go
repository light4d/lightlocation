package main

import (
	"fmt"
	"github.com/light4d/lightlocation/lightlocation"
	"net/http"
)

func main() {
	http.HandleFunc("/testLocationAPI", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Fprintf(writer, "hello, %q", html.EscapeString(request.URL.Path))
		longitude, latitude, _ := lightlocation.GetLocation(request)
		fmt.Fprintf(writer, "Current GPS is "+longitude+","+latitude)
	})

	http.ListenAndServe(":9999", nil)
}
