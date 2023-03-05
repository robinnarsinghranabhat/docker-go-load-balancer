package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList = []*httputil.ReverseProxy{
		httputil.NewSingleHostReverseProxy(createUrl("http://s1:8000")),
		httputil.NewSingleHostReverseProxy(createUrl("http://s2:8000")),
	}
	lastServedIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	// go startHealthCheck()
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	// server, err := getHealthyServer()
	// if err != nil {
	// 	http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
	// 	return
	// }

	server := getServer()
	server.ServeHTTP(res, req)
}


func getServer() *httputil.ReverseProxy {
	nextIndex := (lastServedIndex + 1) % len(serverList)
	server := serverList[nextIndex]
	lastServedIndex = nextIndex
	return server
}

func createUrl(urlname string) *url.URL {
	u, _ := url.Parse(urlname)
	return u
}

