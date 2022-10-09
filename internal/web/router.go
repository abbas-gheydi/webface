package web

import (
	"log"
	"net/http"
)

const listenAddress = "0.0.0.0:8080"

func Router() {
	http.HandleFunc("/login/", loginPage)
	http.Handle("/", MustAuth(ProxyRequestHandler(proxy)))
	log.Printf("start Listenning on %v ... \n", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
