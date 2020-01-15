package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"api/mocks"
	"api/model"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/gorilla/mux"
)

type EntryControllerTestSuite struct {
	suite.Suite
	entryRepository *mocks.EntryRepository
	ctrl            EntryController
	router          *mux.Router
}

func (suite *EntryControllerTestSuite) SetupSuite() {
	fmt.Fprintf(os.Stdout, "\n*********** SetupSuite [EntryControllerTestSuite] ***********\n")

	suite.entryRepository = new(mocks.EntryRepository)
	suite.router = mux.NewRouter()
	suite.ctrl = NewEntryController(suite.entryRepository)
}

func (suite *EntryControllerTestSuite) TestEntryControllerGetAll() {
	visitorsInDb := []model.Visitor{
		model.Visitor{
			Name:  gofakeit.Name(),
			Notes: gofakeit.Company(),
		},
		model.Visitor{
			Name:  gofakeit.Name(),
			Notes: gofakeit.Company(),
		},
	}

	suite.entryRepository.On("GetAll").Return(visitorsInDb, nil)
	suite.router.HandleFunc("/entries", suite.ctrl.GetEntriesHandler).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "/entries", nil)
	respRec := httptest.NewRecorder()
	suite.router.ServeHTTP(respRec, req)

	var visitors []model.Visitor
	body, _ := ioutil.ReadAll(respRec.Body)
	json.Unmarshal(body, &visitors)

	assert.Equal(suite.T(), respRec.Code, http.StatusOK)
	assert.Equal(suite.T(), len(visitors), 2)
	assert.Equal(suite.T(), visitors[0].Name, visitorsInDb[0].Name)
	assert.Equal(suite.T(), visitors[1].Notes, visitorsInDb[1].Notes)
}

func (suite *EntryControllerTestSuite) TestEntryControllerCreate() {
	johnSmith := model.Visitor{
		Name:  gofakeit.Name(),
		Notes: gofakeit.Company(),
	}

	suite.entryRepository.On("Create", johnSmith).Return(nil)
	suite.router.HandleFunc("/entries", suite.ctrl.PostEntriesHandler).Methods(http.MethodPost)

	requestByte, _ := json.Marshal(johnSmith)
	requestReader := bytes.NewReader(requestByte)
	req := httptest.NewRequest(http.MethodPost, "/entries", requestReader)
	respRec := httptest.NewRecorder()
	suite.router.ServeHTTP(respRec, req)

	assert.Equal(suite.T(), respRec.Code, http.StatusAccepted)
}

func (suite *EntryControllerTestSuite) TestEntryControllerUpdate() {
	johnSmith := model.Visitor{
		Name:  gofakeit.Name(),
		Notes: gofakeit.Company(),
		ID:    primitive.NewObjectID(),
	}

	suite.entryRepository.On("SignOut", johnSmith.ID).Return(johnSmith, nil)
	suite.router.HandleFunc("/entries/{id}", suite.ctrl.SignOutEntryHandler).Methods(http.MethodPatch)

	req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/entries/%v", johnSmith.ID.Hex()), nil)
	respRec := httptest.NewRecorder()
	suite.router.ServeHTTP(respRec, req)

	assert.Equal(suite.T(), respRec.Code, http.StatusOK)
}

func TestEntryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(EntryControllerTestSuite))
}
