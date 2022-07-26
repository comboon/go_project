package main
import (
	"fmt"
	"net"
	"net/http"
	"os"   
	"strings"
	"log"
	"io"
)
func index(w http.ResponseWriter, r *http.Request){
	os.Setenv("VERSION","v1.1")
	version :=os.Getenv("VERSION")
	w.Header().Set("VERSION",version)
	fmt.Printf("os version: %s \n", version)
	name := r.URL.Query().Get("name")
	if name != "" {
		io.WriteString(w, fmt.Sprintf("hello %s\n", name))
	} else {
		io.WriteString(w, "hello default\n")
	}
	for k ,v := range r.Header {
		for _ , vv:= range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)        
			w.Header().Set(k, vv)	
		}
	}

	clientip := ClientIP(r)  
	log.Printf("Success! Response code: %d", 200)   
	log.Printf("Success! clientip: %s", clientip)
}
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}
func ClientIP(r *http.Request) string {   
	xForwardedFor := r.Header.Get("X-Forwarded-For")   
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])   
	if ip != "" {      
		return ip   
	}   
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))  
	if ip != "" {      
		return ip   
	}   
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {     
		return ip  
	}   
	return ""
}
func main() {
	mux := http.NewServeMux()
	fmt.Println("Start server...")
	mux.HandleFunc("/test/yl",index)
	mux.HandleFunc("/healthz",healthz)
	if err := http.ListenAndServe(":8080",mux); err !=nil {
		log.Fatalf("start http server failed, error: %s\n",err.Error())
	}
}
