package main

import (
	"fmt"
	"log"
	"net/http"
)


const webPosrt  = "80"

type Config struct {}
func main()  {
	app := Config{}

	log.Printf("Starting Broker servive on port %s\n",webPosrt)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPosrt),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

	
}