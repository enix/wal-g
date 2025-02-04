// Code generated by Yandex patched mockery v1.1.0. DO NOT EDIT.

package stagesmocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/enix/wal-g/pkg/databases/mongo/models"
)

// BetweenFetcher is an autogenerated mock type for the BetweenFetcher type
type BetweenFetcher struct {
	mock.Mock
}

// FetchBetween provides a mock function with given fields: _a0, _a1, _a2
func (_m *BetweenFetcher) FetchBetween(_a0 context.Context, _a1 models.Timestamp, _a2 models.Timestamp) (chan *models.Oplog, chan error, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 1 {
		rf, ok := ret.Get(0).(func(context.Context, models.Timestamp, models.Timestamp) (chan *models.Oplog, chan error, error))
		if ok {
			return rf(_a0, _a1, _a2)
		}
	}

	var r0 chan *models.Oplog
	if rf, ok := ret.Get(0).(func(context.Context, models.Timestamp, models.Timestamp) chan *models.Oplog); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *models.Oplog)
		}
	}

	var r1 chan error
	if rf, ok := ret.Get(1).(func(context.Context, models.Timestamp, models.Timestamp) chan error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, models.Timestamp, models.Timestamp) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
