package main

import (
	"fmt"
	"github.com/light4d/lightlocation"
	"net/http"
	_ "net/http/pprof"
	"log"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:10000", nil))
	}()
	http.HandleFunc("/testLocationAPI", func(writer http.ResponseWriter, request *http.Request) {
		longitude, latitude, _ := lightlocation.GetLocation(request)
		fmt.Fprintf(writer, "Current GPS is "+longitude+","+latitude)
	})

	http.ListenAndServe(":9999", nil)
}
