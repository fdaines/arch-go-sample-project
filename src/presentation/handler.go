package presentation

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/arch-go-sample-project/src/businesslogic"
	"github.com/fdaines/arch-go-sample-project/src/persistence"
	"net/http"
)

func CreateHandler() {
	http.HandleFunc("/", handlerFunction)
}

func handlerFunction(w http.ResponseWriter, r *http.Request){
	dummyValue := persistence.RetrieveDataFromDatabase()
	fmt.Printf("%s\n", dummyValue)

	response := &Response{
		Value: businesslogic.GetProcessedData(),
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(resp))
}

type Response struct {
	Value string `json"value"`
}