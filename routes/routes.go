package routes

import "net/http"


func HandleRoutes(r *http.ServeMux){
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// w.Write([]byte)
	})
}

