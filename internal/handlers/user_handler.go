package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Rai-Sahil/go-setup/internal/database"
	"github.com/Rai-Sahil/go-setup/internal/models"
	"io/ioutil"
	"net/http"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	err := database.CreateTable()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var newUser models.User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.InsertUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintf(w, "User added successfully")
	if err != nil {
		return
	}
}
