package routes

import (
	"go-echo-server/handlers"
	"net/http"
)


func HandleRoutes(r *http.ServeMux){
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// w.Write([]byte)
	})
	
	r.HandleFunc("/echo",handlers.Echo)
}

