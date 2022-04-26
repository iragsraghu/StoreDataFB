package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App Running...")
	})
	router.HandleFunc("/generateData", StoreDeviceData).Methods("POST")
	router.HandleFunc("/listData", ListStoreData).Methods("GET")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
