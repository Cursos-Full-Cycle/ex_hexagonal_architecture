package server

import (
	"go_hexagonal/application"	
	"net/http"
	"time"
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())
	
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}