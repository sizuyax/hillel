// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"project-auction/internal/domain/entity"
)

// PGUserRepository is an autogenerated mock type for the PGUserRepository type
type PGUserRepository struct {
	mock.Mock
}

// InsertUser provides a mock function with given fields: _a0
func (_m *PGUserRepository) InsertUser(_a0 *entity.User) (*entity.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for InsertUser")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.User) (*entity.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPGUserRepository creates a new instance of PGUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPGUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PGUserRepository {
	mock := &PGUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}