package presentation

import (
	"encoding/json"
	"fmt"
	"github.com/fdaines/arch-go-sample-project/src/businesslogic"
	"net/http"
)

func CreateHandler() {
	http.HandleFunc("/", handlerFunction)
}

func handlerFunction(w http.ResponseWriter, r *http.Request){
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