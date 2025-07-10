package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	errorshandler "github.com/tacheraSasi/tripwire/errorsHandler"
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

	response := map[string]string{"echo": echoReq.message}
	fmt.Println(response)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)

}

func echoPost(w http.ResponseWriter,r *http.Request){
	body, err := io.ReadAll(r.Body)
	errorshandler.Check(err, "failed to read the body--->Post")

	defer r.Body.Close()
	var data map[string]string

	if err := json.Unmarshal(body, &data); err != nil{
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	return data

}
