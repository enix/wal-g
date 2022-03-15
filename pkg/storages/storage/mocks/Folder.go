// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
	storage "github.com/enix/wal-g/pkg/storages/storage"
)

// Folder is an autogenerated mock type for the Folder type
type Folder struct {
	mock.Mock
}

// DeleteObjects provides a mock function with given fields: objectRelativePaths
func (_m *Folder) DeleteObjects(objectRelativePaths []string) error {
	ret := _m.Called(objectRelativePaths)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(objectRelativePaths)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: objectRelativePath
func (_m *Folder) Exists(objectRelativePath string) (bool, error) {
	ret := _m.Called(objectRelativePath)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(objectRelativePath)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(objectRelativePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPath provides a mock function with given fields:
func (_m *Folder) GetPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSubFolder provides a mock function with given fields: subFolderRelativePath
func (_m *Folder) GetSubFolder(subFolderRelativePath string) storage.Folder {
	ret := _m.Called(subFolderRelativePath)

	var r0 storage.Folder
	if rf, ok := ret.Get(0).(func(string) storage.Folder); ok {
		r0 = rf(subFolderRelativePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Folder)
		}
	}

	return r0
}

// ListFolder provides a mock function with given fields:
func (_m *Folder) ListFolder() ([]storage.Object, []storage.Folder, error) {
	ret := _m.Called()

	var r0 []storage.Object
	if rf, ok := ret.Get(0).(func() []storage.Object); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.Object)
		}
	}

	var r1 []storage.Folder
	if rf, ok := ret.Get(1).(func() []storage.Folder); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]storage.Folder)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PutObject provides a mock function with given fields: name, content
func (_m *Folder) PutObject(name string, content io.Reader) error {
	ret := _m.Called(name, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(name, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadObject provides a mock function with given fields: objectRelativePath
func (_m *Folder) ReadObject(objectRelativePath string) (io.ReadCloser, error) {
	ret := _m.Called(objectRelativePath)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(objectRelativePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(objectRelativePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
