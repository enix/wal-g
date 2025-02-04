// Code generated by Yandex patched mockery v1.1.0. DO NOT EDIT.

package archivemocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/enix/wal-g/pkg/databases/mongo/models"
)

// Purger is an autogenerated mock type for the Purger type
type Purger struct {
	mock.Mock
}

// DeleteBackups provides a mock function with given fields: backups
func (_m *Purger) DeleteBackups(backups []models.Backup) error {
	ret := _m.Called(backups)

	var r0 error
	if rf, ok := ret.Get(0).(func([]models.Backup) error); ok {
		r0 = rf(backups)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGarbage provides a mock function with given fields: garbage
func (_m *Purger) DeleteGarbage(garbage []string) error {
	ret := _m.Called(garbage)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(garbage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOplogArchives provides a mock function with given fields: archives
func (_m *Purger) DeleteOplogArchives(archives []models.Archive) error {
	ret := _m.Called(archives)

	var r0 error
	if rf, ok := ret.Get(0).(func([]models.Archive) error); ok {
		r0 = rf(archives)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
