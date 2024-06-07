package app

import (
	"github.com/Rai-Sahil/go-setup/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello World"))
		if err != nil {
			return
		}
	})

	r.HandleFunc("/addUser", handlers.AddUserHandler).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
