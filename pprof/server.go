package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var data = make([]byte, 0)

func main() {

	http.HandleFunc("/enlarge", func(w http.ResponseWriter, r *http.Request) {
		d := make([]byte, 10*1024*1024) // 10M
		data = append(data, d...)
		fmt.Fprintf(w, "slice len: %d", len(data))
	})

	log.Println(http.ListenAndServe("localhost:6060", nil))
}
