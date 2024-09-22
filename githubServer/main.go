package main

import (
	"log"
	"gitserver/server"
	"net/http"
	"fmt"

)

func main(){
	route := server.GitModel{}
	fmt.Println("[*] Runing server in 0.0.0.0:8000")
	http.HandleFunc("/",route.Get)
	if err := http.ListenAndServe("0.0.0.0:8000",nil); err != nil {
		log.Fatalf("Error server in %s",err)
	}

}