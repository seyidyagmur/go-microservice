package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/turbine/config", turbineConfigHandler)
	http.HandleFunc("/payments", createPaymentHandler)
	log.Println("payment service is up! at host:4002")
	log.Fatal(http.ListenAndServe(":4002", nil))
}
func createPaymentHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"message":"paid"}`)))

}
func turbineConfigHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf(`{"serviceName":"payment-service"}`)))
}
