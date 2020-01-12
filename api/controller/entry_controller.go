package controller

import (
	"api/db"
	"encoding/json"
	"fmt"
	"net/http"

	"api/model"
)

type EntryController interface {
	GetEntriesHandler(http.ResponseWriter, *http.Request)
	PostEntiresHandler(http.ResponseWriter, *http.Request)
}
type EntryControllerConcrete struct {
	repository db.EntryRepository
}

func NewEntryController(repository db.EntryRepository) EntryControllerConcrete {
	return EntryControllerConcrete{
		repository: repository,
	}
}

func (c *EntryControllerConcrete) GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	all, err := c.repository.GetAll()

	if nil != err {
		WriteError(w, fmt.Errorf("Error getting all", err))
		return
	}

	response, _ := json.Marshal(all)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *EntryControllerConcrete) PostEntriesHandler(w http.ResponseWriter, r *http.Request) {
	visitor, err := decodeAndValidate(r)
	if nil != err {
		fmt.Println("1")
		WriteError(w, err)
		return
	}

	err = c.repository.Create(visitor)
	if nil != err {
		WriteError(w, fmt.Errorf("Error creating", err))
		return
	}

	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusAccepted)
}

func WriteError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("bad request", err)))
}

func decodeAndValidate(r *http.Request) (model.Visitor, error) {
	var visitor model.Visitor
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&visitor)

	if err == nil {
		return visitor, err
	}

	return visitor, visitor.Validate(r)
}
