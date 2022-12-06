package app

import (
	"log"
	"net/http"

	"github.com/Tamiquell/banking-API/domain"
	"github.com/Tamiquell/banking-API/service"
	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandler{
		service.NewCustomerService(
			domain.NewCustomerRepositoryStub(),
		),
	}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe("127.0.0.1:8000", router))
}
