package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe("127.0.0.1:8000", router))
}
