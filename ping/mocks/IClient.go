// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// Delete provides a mock function with given fields: path
func (_m *IClient) Delete(path string) (map[string]interface{}, error) {
	ret := _m.Called(path)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: path
func (_m *IClient) Get(path string) (map[string]interface{}, error) {
	ret := _m.Called(path)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: path, body
func (_m *IClient) Post(path string, body map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(path, body)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(path, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(path, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: path, body
func (_m *IClient) Put(path string, body map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(path, body)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(path, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(path, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
