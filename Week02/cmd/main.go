package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"week02/service"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	username := vars["u"]
	res := service.GetUser(username[0])
	json_bytes, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "{\"code\":\"9999\",\"msg\":\"系统处理异常，请稍后再试\"}")
		return
	}
	w.Write(json_bytes)
}

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
