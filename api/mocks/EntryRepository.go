// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "api/model"
import primitive "go.mongodb.org/mongo-driver/bson/primitive"

// EntryRepository is an autogenerated mock type for the EntryRepository type
type EntryRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *EntryRepository) Create(_a0 model.Visitor) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Visitor) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *EntryRepository) GetAll() ([]model.Visitor, error) {
	ret := _m.Called()

	var r0 []model.Visitor
	if rf, ok := ret.Get(0).(func() []model.Visitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Visitor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignOut provides a mock function with given fields: id
func (_m *EntryRepository) SignOut(id primitive.ObjectID) (model.Visitor, error) {
	ret := _m.Called(id)

	var r0 model.Visitor
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) model.Visitor); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Visitor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
