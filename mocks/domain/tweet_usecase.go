// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "atwell/domain"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// TweetUsecase is an autogenerated mock type for the TweetUsecase type
type TweetUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: email, comment
func (_m *TweetUsecase) Create(email string, comment string) (domain.Tweet, error) {
	ret := _m.Called(email, comment)

	var r0 domain.Tweet
	if rf, ok := ret.Get(0).(func(string, string) domain.Tweet); ok {
		r0 = rf(email, comment)
	} else {
		r0 = ret.Get(0).(domain.Tweet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: email, id
func (_m *TweetUsecase) Delete(email string, id uint) error {
	ret := _m.Called(email, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(email, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: email, from, to
func (_m *TweetUsecase) Get(email string, from time.Time, to time.Time) ([]domain.Tweet, error) {
	ret := _m.Called(email, from, to)

	var r0 []domain.Tweet
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time) []domain.Tweet); ok {
		r0 = rf(email, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Tweet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Time, time.Time) error); ok {
		r1 = rf(email, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
