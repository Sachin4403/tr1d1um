// Code generated by mockery v1.0.0. DO NOT EDIT.
package stat

import (
	"github.com/Comcast/tr1d1um/src/tr1d1um/common"
	"github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// RequestStat provides a mock function with given fields: authHeaderValue, deviceID
func (_m *MockService) RequestStat(authHeaderValue string, deviceID string) (*common.XmidtResponse, error) {
	ret := _m.Called(authHeaderValue, deviceID)

	var r0 *common.XmidtResponse
	if rf, ok := ret.Get(0).(func(string, string) *common.XmidtResponse); ok {
		r0 = rf(authHeaderValue, deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.XmidtResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(authHeaderValue, deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
