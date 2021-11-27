package main

import (
	"github.com/SvyatobatkoVlad/Rest-API/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

	//func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	//	name := params.ByName("name")
	//	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
	//}


	func main() {
		log.Println("create router")
		router := httprouter.New()

		handler := user.NewHandler()
		handler.Register(router)

		run(router)
	}

	func run(router *httprouter.Router) {
		listener, err := net.Listen("tcp", ":1234")
		if err != nil{
			panic(err)
		}

		server := &http.Server{
			Handler: router,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Fatalln(server.Serve(listener))
	}