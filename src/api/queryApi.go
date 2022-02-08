package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Maryam-Yumna/FCL-Website-Backend/src/controller"
	"github.com/Maryam-Yumna/FCL-Website-Backend/src/types"
)

type Query struct {
	QueryString string `json:"QueryString"`
}

func ExecuteQuery(w http.ResponseWriter, r *http.Request) {

	var newQuery types.Query
	//update content type
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter query")
	}

	json.Unmarshal(reqBody, &newQuery)

	result := controller.ExecuteQuery(newQuery)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}

}
