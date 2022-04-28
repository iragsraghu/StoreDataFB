package main

import (
	"StoreDataFB/controller/demo_controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "App Running...")
	// })
	// router.Handle("/", http.FileServer(http.Dir("./template")))
	router.HandleFunc("/", demo_controller.Index)
	router.HandleFunc("/register", StoreDeviceData).Methods("POST")
	router.HandleFunc("/listData", ListStoreData).Methods("GET")
	router.HandleFunc("/referrals", demo_controller.ReferralsPage).Methods("GET")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
