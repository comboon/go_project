package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

)

func main() {
	fmt.Println("Start server...")
	http.HandleFunc("/test", healthz)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client request...")
	name := r.URL.Query().Get("name")
	version :=os.Getenv("version")
	//fmt.Println(version)
	io.WriteString(w, fmt.Sprintf("version=%s\n", version))
	if name != "" {
		io.WriteString(w, fmt.Sprintf("hello %s\n", name))
	} else {
		io.WriteString(w, "hello default\n")
	}
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
