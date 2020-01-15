package controller

import (
	"api/db"
	"encoding/json"
	"fmt"
	"net/http"

	"api/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EntryController struct {
	repository db.EntryRepository
}

func NewEntryController(repository db.EntryRepository) EntryController {
	return EntryController{
		repository: repository,
	}
}

func (c *EntryController) GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	all, err := c.repository.GetAll()

	if nil != err {
		WriteError(w, fmt.Errorf("Error getting all: %v", err))
		return
	}

	response, _ := json.Marshal(all)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *EntryController) PostEntriesHandler(w http.ResponseWriter, r *http.Request) {
	visitor, err := decodeAndValidate(r)
	if nil != err {
		WriteError(w, err)
		return
	}

	err = c.repository.Create(visitor)
	if nil != err {
		WriteError(w, fmt.Errorf("Error creating: %v", err))
		return
	}

	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusAccepted)
}

func (c *EntryController) SignOutEntryHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	visitorId, err := primitive.ObjectIDFromHex(id)
	if nil != err {
		WriteError(w, fmt.Errorf("invalid object id: %v", id))
		return
	}

	update, err := c.repository.SignOut(visitorId)
	if nil != err {
		WriteError(w, fmt.Errorf("Error signing out: %v", err))
		return
	}

	response, _ := json.Marshal(update)

	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func WriteError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func decodeAndValidate(r *http.Request) (model.Visitor, error) {
	var visitor model.Visitor
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&visitor)

	if err != nil {
		return visitor, err
	}

	return visitor, visitor.Validate()
}
