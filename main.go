package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	//"github.com/nicholasjackson/env"
	"server/handlers"

	//"golang.org/x/tools/go/callgraph/cha"
)

//var bindAddress = env.String("BIND_ADDRESS", false, "Bind address for the server")

func main(){
	// env.Parse()
	    l :=log.New(os.Stdout, "product-api", log.LstdFlags)
         //hh:= handlers.NewHello(l)
         gh:=handlers.NewGoodBye(l)
		 gp:=handlers.NewProducts(l)

		 sm:=http.NewServeMux()
		 sm.Handle("/", gp)
		 sm.Handle("/goodbye", gh)
	  
	
	s := http.Server{
		Addr: ":9090",
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
		IdleTimeout: 120*time.Second,
	}

	go func ()  {
		l.Println("Starting server on port 9090")

		err:= s.ListenAndServe()
		if err!=nil{
			l.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	
    sigChan :=make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

     tc,_:= context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}