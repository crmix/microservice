package handlers

import(
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)
type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return&Hello{l}
}

func(h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err :=ioutil.ReadAll(r.Body)
		if err !=nil{
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", d)

	}