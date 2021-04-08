package main

import (
	"github.com/fdaines/arch-go-sample-project/src/presentation"
	"log"
	"net/http"
)

func main() {
	presentation.CreateHandler()
	log.Fatal(http.ListenAndServe(":10000", nil))
}