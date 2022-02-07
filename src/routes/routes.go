package routes

import (
	"fmt"
	"net/http"

	"github.com/Maryam-Yumna/FCL-Website-Backend/src/api"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	apiRoutes := router.PathPrefix("/api").Subrouter().StrictSlash(true)

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Home Route")

	})

	//execute query
	apiRoutes.HandleFunc("/query/execute", api.ExecuteQuery).Methods("POST")
	return router
}
