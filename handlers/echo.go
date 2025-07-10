package handlers

import (
	"encoding/json"
	"net/http"
)

type  EchoRequest struct{
	message string
}

type EchoResponse struct{
	echo string
}

func Echo(w http.ResponseWriter, r *http.Request){
	var echoReq EchoRequest;
	switch r.Method{
		case http.MethodGet:
			message := r.URL.Query().Get("message")
			echoReq = EchoRequest{
				message,
			}
		default:
			return 
	}
	
	response := EchoResponse{echo: echoReq.message}
	w.Header().Set("Content-Type","json")
	json.NewEncoder(w).Encode(response)
	
}
