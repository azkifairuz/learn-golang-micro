package main

import (
	"fmt"
	"log"
	"net/http"
)


const webPosrt  = "8002"

type Config struct {}
func main()  {
	app := Config{}

	log.Printf("Starting Broker servive on port %s\n",webPosrt)

	//define the server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPosrt),
		Handler: app.routes(),
	}

	//start server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

	
}