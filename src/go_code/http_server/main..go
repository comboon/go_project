package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

)

func main() {
	fmt.Println("Starting http server...")
	http.HandleFunc("/test", rootHandler)
	err := http.ListenAndServe(":80", nil)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", rootHandler)
	// mux.HandleFunc("/healthz", healthz)
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

// func healthz(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "ok\n")
// }

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	version :=os.Getenv("Version")
	//fmt.Println(version)
	io.WriteString(w, fmt.Sprintf("version=%s\n", version))
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}