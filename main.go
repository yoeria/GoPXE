package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/yoeria/gopxe/conf"
	"github.com/yoeria/gopxe/routes"
)

// This is the main package
// Output is webserver om port
func main() {
	conf.Setup()
	port := flag.Lookup("port").Value.String()

	routes := routes.New()
	log.Printf("Serving on port: %s", port)
	if err := http.ListenAndServe(":"+port, routes); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
