package reqhandlers

import (
	"context"
	"encoding/json"
	"firstMS/repository"
	"firstMS/repository/models"
	"io/ioutil"
	"net/http"
)

func GetAddressBook(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	book, err := repository.AddressBookDb.GetAddressBook(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var person models.Person
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(req, &person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = repository.AddressBookQueue.StoreOnePerson(ctx, person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("person added"))
}
