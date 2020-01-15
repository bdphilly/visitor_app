package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"api/model"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type EntryRepositoryTestSuite struct {
	suite.Suite
	db         *mongo.Database
	repository EntryRepository
	visitors   []model.Visitor
}

func (suite *EntryRepositoryTestSuite) SetupTest() {
	fmt.Fprintf(os.Stdout, "\n*********** SetupSuite [EntryRepositoryTestSuite] ***********\n")

	ctx := context.Background()
	suite.db, _ = ConfigureDB(ctx)

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

	suite.visitors = visitorsInDb
	seed := make([]interface{}, len(visitorsInDb))
	for i := range visitorsInDb {
		seed[i] = visitorsInDb[i]
	}
	suite.db.Collection("entry").InsertMany(ctx, seed)

	suite.repository = NewEntryRepository(ctx, suite.db)
}

func (suite *EntryRepositoryTestSuite) TearDownTest() {
	fmt.Fprintf(os.Stdout, "\n*********** TeardownSuite [EntryRepositoryTestSuite] ***********\n")

	ctx := context.Background()
	suite.db.Collection("entry").Drop(ctx)
}

func (suite *EntryRepositoryTestSuite) TestEntryRepositoryGetAll() {
	res, err := suite.repository.GetAll()

	assert.Len(suite.T(), res, 2)
	assert.NotNil(suite.T(), res[0].ID)
	assert.NotNil(suite.T(), res[0].CreatedAt)
	assert.Nil(suite.T(), res[0].SignOutTime)
	assert.Nil(suite.T(), err)
}

func (suite *EntryRepositoryTestSuite) TestEntryRepositoryCreate() {
	existing, _ := suite.repository.GetAll()
	assert.Len(suite.T(), existing, 2)

	newVisitor := model.Visitor{
		Name:  gofakeit.Name(),
		Notes: gofakeit.Company(),
	}
	err := suite.repository.Create(newVisitor)
	assert.Nil(suite.T(), err)

	existing, _ = suite.repository.GetAll()
	assert.Len(suite.T(), existing, 3)
}

func (suite *EntryRepositoryTestSuite) TestEntryRepositorySignOut() {
	existing, _ := suite.repository.GetAll()
	assert.Nil(suite.T(), existing[0].SignOutTime)

	res, err := suite.repository.SignOut(existing[0].ID)

	assert.NotNil(suite.T(), res.SignOutTime)
	assert.Nil(suite.T(), err)
}

func TestEntryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(EntryRepositoryTestSuite))
}
