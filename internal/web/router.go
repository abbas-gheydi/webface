package web

import (
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/wf_login/", loginPage)
	http.Handle("/wf_login/assets/", http.StripPrefix("/wf_login/assets/", http.FileServer(http.Dir("./web/assets"))))
	http.Handle("/", MustAuth(ProxyRequestHandler(proxy)))
	log.Printf("start Listenning on %v ... \n", LISTEN_ADDR)
	log.Fatal(http.ListenAndServe(LISTEN_ADDR, nil))
}
