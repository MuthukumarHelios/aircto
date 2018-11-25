package main

import (
	"fmt"
   "net/http"
   "os"
   "log"
   "github.com/joho/godotenv"
   "aircto/routes"
)

func main(){

	fmt.Println("-----")
    err :=	godotenv.Load("config.env")
   if err != nil{
     	panic(err)
   }

     port := os.Getenv("port")
    log.Println("--prt", port)

 	http.ListenAndServe(port, routes.MainRoutes())
}