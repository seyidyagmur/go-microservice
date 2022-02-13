package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/turbine/config", turbineConfigHandler)
	log.Println("shipping service is up! at host:4003")
	log.Fatal(http.ListenAndServe(":4003", nil))
}

func turbineConfigHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"serviceName":"shipping-service"}`)))
}
