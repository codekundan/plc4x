/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package bacnetip

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	mock "github.com/stretchr/testify/mock"

	spi "github.com/apache/plc4x/plc4go/spi"

	utils "github.com/apache/plc4x/plc4go/spi/utils"
)

// MockPDU is an autogenerated mock type for the PDU type
type MockPDU struct {
	mock.Mock
}

type MockPDU_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPDU) EXPECT() *MockPDU_Expecter {
	return &MockPDU_Expecter{mock: &_m.Mock}
}

// DeepCopy provides a mock function with given fields:
func (_m *MockPDU) DeepCopy() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeepCopy")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockPDU_DeepCopy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeepCopy'
type MockPDU_DeepCopy_Call struct {
	*mock.Call
}

// DeepCopy is a helper method to define mock.On call
func (_e *MockPDU_Expecter) DeepCopy() *MockPDU_DeepCopy_Call {
	return &MockPDU_DeepCopy_Call{Call: _e.mock.On("DeepCopy")}
}

func (_c *MockPDU_DeepCopy_Call) Run(run func()) *MockPDU_DeepCopy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_DeepCopy_Call) Return(_a0 interface{}) *MockPDU_DeepCopy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_DeepCopy_Call) RunAndReturn(run func() interface{}) *MockPDU_DeepCopy_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields:
func (_m *MockPDU) Get() (byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 byte
	var r1 error
	if rf, ok := ret.Get(0).(func() (byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() byte); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(byte)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPDU_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockPDU_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockPDU_Expecter) Get() *MockPDU_Get_Call {
	return &MockPDU_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockPDU_Get_Call) Run(run func()) *MockPDU_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_Get_Call) Return(_a0 byte, _a1 error) *MockPDU_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPDU_Get_Call) RunAndReturn(run func() (byte, error)) *MockPDU_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetData provides a mock function with given fields: dlen
func (_m *MockPDU) GetData(dlen int) ([]byte, error) {
	ret := _m.Called(dlen)

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]byte, error)); ok {
		return rf(dlen)
	}
	if rf, ok := ret.Get(0).(func(int) []byte); ok {
		r0 = rf(dlen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(dlen)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPDU_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type MockPDU_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
//   - dlen int
func (_e *MockPDU_Expecter) GetData(dlen interface{}) *MockPDU_GetData_Call {
	return &MockPDU_GetData_Call{Call: _e.mock.On("GetData", dlen)}
}

func (_c *MockPDU_GetData_Call) Run(run func(dlen int)) *MockPDU_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockPDU_GetData_Call) Return(_a0 []byte, _a1 error) *MockPDU_GetData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPDU_GetData_Call) RunAndReturn(run func(int) ([]byte, error)) *MockPDU_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// GetExpectingReply provides a mock function with given fields:
func (_m *MockPDU) GetExpectingReply() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetExpectingReply")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPDU_GetExpectingReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExpectingReply'
type MockPDU_GetExpectingReply_Call struct {
	*mock.Call
}

// GetExpectingReply is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetExpectingReply() *MockPDU_GetExpectingReply_Call {
	return &MockPDU_GetExpectingReply_Call{Call: _e.mock.On("GetExpectingReply")}
}

func (_c *MockPDU_GetExpectingReply_Call) Run(run func()) *MockPDU_GetExpectingReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetExpectingReply_Call) Return(_a0 bool) *MockPDU_GetExpectingReply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetExpectingReply_Call) RunAndReturn(run func() bool) *MockPDU_GetExpectingReply_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBits provides a mock function with given fields: ctx
func (_m *MockPDU) GetLengthInBits(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLengthInBits")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockPDU_GetLengthInBits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBits'
type MockPDU_GetLengthInBits_Call struct {
	*mock.Call
}

// GetLengthInBits is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPDU_Expecter) GetLengthInBits(ctx interface{}) *MockPDU_GetLengthInBits_Call {
	return &MockPDU_GetLengthInBits_Call{Call: _e.mock.On("GetLengthInBits", ctx)}
}

func (_c *MockPDU_GetLengthInBits_Call) Run(run func(ctx context.Context)) *MockPDU_GetLengthInBits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPDU_GetLengthInBits_Call) Return(_a0 uint16) *MockPDU_GetLengthInBits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetLengthInBits_Call) RunAndReturn(run func(context.Context) uint16) *MockPDU_GetLengthInBits_Call {
	_c.Call.Return(run)
	return _c
}

// GetLengthInBytes provides a mock function with given fields: ctx
func (_m *MockPDU) GetLengthInBytes(ctx context.Context) uint16 {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLengthInBytes")
	}

	var r0 uint16
	if rf, ok := ret.Get(0).(func(context.Context) uint16); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// MockPDU_GetLengthInBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLengthInBytes'
type MockPDU_GetLengthInBytes_Call struct {
	*mock.Call
}

// GetLengthInBytes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPDU_Expecter) GetLengthInBytes(ctx interface{}) *MockPDU_GetLengthInBytes_Call {
	return &MockPDU_GetLengthInBytes_Call{Call: _e.mock.On("GetLengthInBytes", ctx)}
}

func (_c *MockPDU_GetLengthInBytes_Call) Run(run func(ctx context.Context)) *MockPDU_GetLengthInBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPDU_GetLengthInBytes_Call) Return(_a0 uint16) *MockPDU_GetLengthInBytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetLengthInBytes_Call) RunAndReturn(run func(context.Context) uint16) *MockPDU_GetLengthInBytes_Call {
	_c.Call.Return(run)
	return _c
}

// GetLong provides a mock function with given fields:
func (_m *MockPDU) GetLong() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLong")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPDU_GetLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLong'
type MockPDU_GetLong_Call struct {
	*mock.Call
}

// GetLong is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetLong() *MockPDU_GetLong_Call {
	return &MockPDU_GetLong_Call{Call: _e.mock.On("GetLong")}
}

func (_c *MockPDU_GetLong_Call) Run(run func()) *MockPDU_GetLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetLong_Call) Return(_a0 int64, _a1 error) *MockPDU_GetLong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPDU_GetLong_Call) RunAndReturn(run func() (int64, error)) *MockPDU_GetLong_Call {
	_c.Call.Return(run)
	return _c
}

// GetNetworkPriority provides a mock function with given fields:
func (_m *MockPDU) GetNetworkPriority() model.NPDUNetworkPriority {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetworkPriority")
	}

	var r0 model.NPDUNetworkPriority
	if rf, ok := ret.Get(0).(func() model.NPDUNetworkPriority); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.NPDUNetworkPriority)
	}

	return r0
}

// MockPDU_GetNetworkPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNetworkPriority'
type MockPDU_GetNetworkPriority_Call struct {
	*mock.Call
}

// GetNetworkPriority is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetNetworkPriority() *MockPDU_GetNetworkPriority_Call {
	return &MockPDU_GetNetworkPriority_Call{Call: _e.mock.On("GetNetworkPriority")}
}

func (_c *MockPDU_GetNetworkPriority_Call) Run(run func()) *MockPDU_GetNetworkPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetNetworkPriority_Call) Return(_a0 model.NPDUNetworkPriority) *MockPDU_GetNetworkPriority_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetNetworkPriority_Call) RunAndReturn(run func() model.NPDUNetworkPriority) *MockPDU_GetNetworkPriority_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUDestination provides a mock function with given fields:
func (_m *MockPDU) GetPDUDestination() *Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUDestination")
	}

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// MockPDU_GetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUDestination'
type MockPDU_GetPDUDestination_Call struct {
	*mock.Call
}

// GetPDUDestination is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetPDUDestination() *MockPDU_GetPDUDestination_Call {
	return &MockPDU_GetPDUDestination_Call{Call: _e.mock.On("GetPDUDestination")}
}

func (_c *MockPDU_GetPDUDestination_Call) Run(run func()) *MockPDU_GetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetPDUDestination_Call) Return(_a0 *Address) *MockPDU_GetPDUDestination_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetPDUDestination_Call) RunAndReturn(run func() *Address) *MockPDU_GetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUSource provides a mock function with given fields:
func (_m *MockPDU) GetPDUSource() *Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUSource")
	}

	var r0 *Address
	if rf, ok := ret.Get(0).(func() *Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Address)
		}
	}

	return r0
}

// MockPDU_GetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUSource'
type MockPDU_GetPDUSource_Call struct {
	*mock.Call
}

// GetPDUSource is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetPDUSource() *MockPDU_GetPDUSource_Call {
	return &MockPDU_GetPDUSource_Call{Call: _e.mock.On("GetPDUSource")}
}

func (_c *MockPDU_GetPDUSource_Call) Run(run func()) *MockPDU_GetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetPDUSource_Call) Return(_a0 *Address) *MockPDU_GetPDUSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetPDUSource_Call) RunAndReturn(run func() *Address) *MockPDU_GetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// GetPDUUserData provides a mock function with given fields:
func (_m *MockPDU) GetPDUUserData() spi.Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPDUUserData")
	}

	var r0 spi.Message
	if rf, ok := ret.Get(0).(func() spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.Message)
		}
	}

	return r0
}

// MockPDU_GetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPDUUserData'
type MockPDU_GetPDUUserData_Call struct {
	*mock.Call
}

// GetPDUUserData is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetPDUUserData() *MockPDU_GetPDUUserData_Call {
	return &MockPDU_GetPDUUserData_Call{Call: _e.mock.On("GetPDUUserData")}
}

func (_c *MockPDU_GetPDUUserData_Call) Run(run func()) *MockPDU_GetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetPDUUserData_Call) Return(_a0 spi.Message) *MockPDU_GetPDUUserData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetPDUUserData_Call) RunAndReturn(run func() spi.Message) *MockPDU_GetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// GetPduData provides a mock function with given fields:
func (_m *MockPDU) GetPduData() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPduData")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// MockPDU_GetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPduData'
type MockPDU_GetPduData_Call struct {
	*mock.Call
}

// GetPduData is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetPduData() *MockPDU_GetPduData_Call {
	return &MockPDU_GetPduData_Call{Call: _e.mock.On("GetPduData")}
}

func (_c *MockPDU_GetPduData_Call) Run(run func()) *MockPDU_GetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetPduData_Call) Return(_a0 []byte) *MockPDU_GetPduData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetPduData_Call) RunAndReturn(run func() []byte) *MockPDU_GetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// GetRootMessage provides a mock function with given fields:
func (_m *MockPDU) GetRootMessage() spi.Message {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRootMessage")
	}

	var r0 spi.Message
	if rf, ok := ret.Get(0).(func() spi.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.Message)
		}
	}

	return r0
}

// MockPDU_GetRootMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRootMessage'
type MockPDU_GetRootMessage_Call struct {
	*mock.Call
}

// GetRootMessage is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetRootMessage() *MockPDU_GetRootMessage_Call {
	return &MockPDU_GetRootMessage_Call{Call: _e.mock.On("GetRootMessage")}
}

func (_c *MockPDU_GetRootMessage_Call) Run(run func()) *MockPDU_GetRootMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetRootMessage_Call) Return(_a0 spi.Message) *MockPDU_GetRootMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_GetRootMessage_Call) RunAndReturn(run func() spi.Message) *MockPDU_GetRootMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetShort provides a mock function with given fields:
func (_m *MockPDU) GetShort() (int16, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetShort")
	}

	var r0 int16
	var r1 error
	if rf, ok := ret.Get(0).(func() (int16, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int16)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPDU_GetShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShort'
type MockPDU_GetShort_Call struct {
	*mock.Call
}

// GetShort is a helper method to define mock.On call
func (_e *MockPDU_Expecter) GetShort() *MockPDU_GetShort_Call {
	return &MockPDU_GetShort_Call{Call: _e.mock.On("GetShort")}
}

func (_c *MockPDU_GetShort_Call) Run(run func()) *MockPDU_GetShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_GetShort_Call) Return(_a0 int16, _a1 error) *MockPDU_GetShort_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPDU_GetShort_Call) RunAndReturn(run func() (int16, error)) *MockPDU_GetShort_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: _a0
func (_m *MockPDU) Put(_a0 byte) {
	_m.Called(_a0)
}

// MockPDU_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockPDU_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - _a0 byte
func (_e *MockPDU_Expecter) Put(_a0 interface{}) *MockPDU_Put_Call {
	return &MockPDU_Put_Call{Call: _e.mock.On("Put", _a0)}
}

func (_c *MockPDU_Put_Call) Run(run func(_a0 byte)) *MockPDU_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(byte))
	})
	return _c
}

func (_c *MockPDU_Put_Call) Return() *MockPDU_Put_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_Put_Call) RunAndReturn(run func(byte)) *MockPDU_Put_Call {
	_c.Call.Return(run)
	return _c
}

// PutData provides a mock function with given fields: _a0
func (_m *MockPDU) PutData(_a0 ...byte) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockPDU_PutData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutData'
type MockPDU_PutData_Call struct {
	*mock.Call
}

// PutData is a helper method to define mock.On call
//   - _a0 ...byte
func (_e *MockPDU_Expecter) PutData(_a0 ...interface{}) *MockPDU_PutData_Call {
	return &MockPDU_PutData_Call{Call: _e.mock.On("PutData",
		append([]interface{}{}, _a0...)...)}
}

func (_c *MockPDU_PutData_Call) Run(run func(_a0 ...byte)) *MockPDU_PutData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]byte, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(byte)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockPDU_PutData_Call) Return() *MockPDU_PutData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_PutData_Call) RunAndReturn(run func(...byte)) *MockPDU_PutData_Call {
	_c.Call.Return(run)
	return _c
}

// PutLong provides a mock function with given fields: _a0
func (_m *MockPDU) PutLong(_a0 uint32) {
	_m.Called(_a0)
}

// MockPDU_PutLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutLong'
type MockPDU_PutLong_Call struct {
	*mock.Call
}

// PutLong is a helper method to define mock.On call
//   - _a0 uint32
func (_e *MockPDU_Expecter) PutLong(_a0 interface{}) *MockPDU_PutLong_Call {
	return &MockPDU_PutLong_Call{Call: _e.mock.On("PutLong", _a0)}
}

func (_c *MockPDU_PutLong_Call) Run(run func(_a0 uint32)) *MockPDU_PutLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockPDU_PutLong_Call) Return() *MockPDU_PutLong_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_PutLong_Call) RunAndReturn(run func(uint32)) *MockPDU_PutLong_Call {
	_c.Call.Return(run)
	return _c
}

// PutShort provides a mock function with given fields: _a0
func (_m *MockPDU) PutShort(_a0 uint16) {
	_m.Called(_a0)
}

// MockPDU_PutShort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutShort'
type MockPDU_PutShort_Call struct {
	*mock.Call
}

// PutShort is a helper method to define mock.On call
//   - _a0 uint16
func (_e *MockPDU_Expecter) PutShort(_a0 interface{}) *MockPDU_PutShort_Call {
	return &MockPDU_PutShort_Call{Call: _e.mock.On("PutShort", _a0)}
}

func (_c *MockPDU_PutShort_Call) Run(run func(_a0 uint16)) *MockPDU_PutShort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint16))
	})
	return _c
}

func (_c *MockPDU_PutShort_Call) Return() *MockPDU_PutShort_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_PutShort_Call) RunAndReturn(run func(uint16)) *MockPDU_PutShort_Call {
	_c.Call.Return(run)
	return _c
}

// Serialize provides a mock function with given fields:
func (_m *MockPDU) Serialize() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Serialize")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPDU_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type MockPDU_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *MockPDU_Expecter) Serialize() *MockPDU_Serialize_Call {
	return &MockPDU_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *MockPDU_Serialize_Call) Run(run func()) *MockPDU_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_Serialize_Call) Return(_a0 []byte, _a1 error) *MockPDU_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPDU_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *MockPDU_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeWithWriteBuffer provides a mock function with given fields: ctx, writeBuffer
func (_m *MockPDU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	ret := _m.Called(ctx, writeBuffer)

	if len(ret) == 0 {
		panic("no return value specified for SerializeWithWriteBuffer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.WriteBuffer) error); ok {
		r0 = rf(ctx, writeBuffer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPDU_SerializeWithWriteBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeWithWriteBuffer'
type MockPDU_SerializeWithWriteBuffer_Call struct {
	*mock.Call
}

// SerializeWithWriteBuffer is a helper method to define mock.On call
//   - ctx context.Context
//   - writeBuffer utils.WriteBuffer
func (_e *MockPDU_Expecter) SerializeWithWriteBuffer(ctx interface{}, writeBuffer interface{}) *MockPDU_SerializeWithWriteBuffer_Call {
	return &MockPDU_SerializeWithWriteBuffer_Call{Call: _e.mock.On("SerializeWithWriteBuffer", ctx, writeBuffer)}
}

func (_c *MockPDU_SerializeWithWriteBuffer_Call) Run(run func(ctx context.Context, writeBuffer utils.WriteBuffer)) *MockPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.WriteBuffer))
	})
	return _c
}

func (_c *MockPDU_SerializeWithWriteBuffer_Call) Return(_a0 error) *MockPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_SerializeWithWriteBuffer_Call) RunAndReturn(run func(context.Context, utils.WriteBuffer) error) *MockPDU_SerializeWithWriteBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// SetExpectingReply provides a mock function with given fields: _a0
func (_m *MockPDU) SetExpectingReply(_a0 bool) {
	_m.Called(_a0)
}

// MockPDU_SetExpectingReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetExpectingReply'
type MockPDU_SetExpectingReply_Call struct {
	*mock.Call
}

// SetExpectingReply is a helper method to define mock.On call
//   - _a0 bool
func (_e *MockPDU_Expecter) SetExpectingReply(_a0 interface{}) *MockPDU_SetExpectingReply_Call {
	return &MockPDU_SetExpectingReply_Call{Call: _e.mock.On("SetExpectingReply", _a0)}
}

func (_c *MockPDU_SetExpectingReply_Call) Run(run func(_a0 bool)) *MockPDU_SetExpectingReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockPDU_SetExpectingReply_Call) Return() *MockPDU_SetExpectingReply_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_SetExpectingReply_Call) RunAndReturn(run func(bool)) *MockPDU_SetExpectingReply_Call {
	_c.Call.Return(run)
	return _c
}

// SetNetworkPriority provides a mock function with given fields: _a0
func (_m *MockPDU) SetNetworkPriority(_a0 model.NPDUNetworkPriority) {
	_m.Called(_a0)
}

// MockPDU_SetNetworkPriority_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNetworkPriority'
type MockPDU_SetNetworkPriority_Call struct {
	*mock.Call
}

// SetNetworkPriority is a helper method to define mock.On call
//   - _a0 model.NPDUNetworkPriority
func (_e *MockPDU_Expecter) SetNetworkPriority(_a0 interface{}) *MockPDU_SetNetworkPriority_Call {
	return &MockPDU_SetNetworkPriority_Call{Call: _e.mock.On("SetNetworkPriority", _a0)}
}

func (_c *MockPDU_SetNetworkPriority_Call) Run(run func(_a0 model.NPDUNetworkPriority)) *MockPDU_SetNetworkPriority_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.NPDUNetworkPriority))
	})
	return _c
}

func (_c *MockPDU_SetNetworkPriority_Call) Return() *MockPDU_SetNetworkPriority_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_SetNetworkPriority_Call) RunAndReturn(run func(model.NPDUNetworkPriority)) *MockPDU_SetNetworkPriority_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUDestination provides a mock function with given fields: _a0
func (_m *MockPDU) SetPDUDestination(_a0 *Address) {
	_m.Called(_a0)
}

// MockPDU_SetPDUDestination_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUDestination'
type MockPDU_SetPDUDestination_Call struct {
	*mock.Call
}

// SetPDUDestination is a helper method to define mock.On call
//   - _a0 *Address
func (_e *MockPDU_Expecter) SetPDUDestination(_a0 interface{}) *MockPDU_SetPDUDestination_Call {
	return &MockPDU_SetPDUDestination_Call{Call: _e.mock.On("SetPDUDestination", _a0)}
}

func (_c *MockPDU_SetPDUDestination_Call) Run(run func(_a0 *Address)) *MockPDU_SetPDUDestination_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockPDU_SetPDUDestination_Call) Return() *MockPDU_SetPDUDestination_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_SetPDUDestination_Call) RunAndReturn(run func(*Address)) *MockPDU_SetPDUDestination_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUSource provides a mock function with given fields: source
func (_m *MockPDU) SetPDUSource(source *Address) {
	_m.Called(source)
}

// MockPDU_SetPDUSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUSource'
type MockPDU_SetPDUSource_Call struct {
	*mock.Call
}

// SetPDUSource is a helper method to define mock.On call
//   - source *Address
func (_e *MockPDU_Expecter) SetPDUSource(source interface{}) *MockPDU_SetPDUSource_Call {
	return &MockPDU_SetPDUSource_Call{Call: _e.mock.On("SetPDUSource", source)}
}

func (_c *MockPDU_SetPDUSource_Call) Run(run func(source *Address)) *MockPDU_SetPDUSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Address))
	})
	return _c
}

func (_c *MockPDU_SetPDUSource_Call) Return() *MockPDU_SetPDUSource_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_SetPDUSource_Call) RunAndReturn(run func(*Address)) *MockPDU_SetPDUSource_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDUUserData provides a mock function with given fields: _a0
func (_m *MockPDU) SetPDUUserData(_a0 spi.Message) {
	_m.Called(_a0)
}

// MockPDU_SetPDUUserData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDUUserData'
type MockPDU_SetPDUUserData_Call struct {
	*mock.Call
}

// SetPDUUserData is a helper method to define mock.On call
//   - _a0 spi.Message
func (_e *MockPDU_Expecter) SetPDUUserData(_a0 interface{}) *MockPDU_SetPDUUserData_Call {
	return &MockPDU_SetPDUUserData_Call{Call: _e.mock.On("SetPDUUserData", _a0)}
}

func (_c *MockPDU_SetPDUUserData_Call) Run(run func(_a0 spi.Message)) *MockPDU_SetPDUUserData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spi.Message))
	})
	return _c
}

func (_c *MockPDU_SetPDUUserData_Call) Return() *MockPDU_SetPDUUserData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_SetPDUUserData_Call) RunAndReturn(run func(spi.Message)) *MockPDU_SetPDUUserData_Call {
	_c.Call.Return(run)
	return _c
}

// SetPduData provides a mock function with given fields: _a0
func (_m *MockPDU) SetPduData(_a0 []byte) {
	_m.Called(_a0)
}

// MockPDU_SetPduData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPduData'
type MockPDU_SetPduData_Call struct {
	*mock.Call
}

// SetPduData is a helper method to define mock.On call
//   - _a0 []byte
func (_e *MockPDU_Expecter) SetPduData(_a0 interface{}) *MockPDU_SetPduData_Call {
	return &MockPDU_SetPduData_Call{Call: _e.mock.On("SetPduData", _a0)}
}

func (_c *MockPDU_SetPduData_Call) Run(run func(_a0 []byte)) *MockPDU_SetPduData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockPDU_SetPduData_Call) Return() *MockPDU_SetPduData_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPDU_SetPduData_Call) RunAndReturn(run func([]byte)) *MockPDU_SetPduData_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPDU) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPDU_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPDU_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPDU_Expecter) String() *MockPDU_String_Call {
	return &MockPDU_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPDU_String_Call) Run(run func()) *MockPDU_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPDU_String_Call) Return(_a0 string) *MockPDU_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_String_Call) RunAndReturn(run func() string) *MockPDU_String_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: pci
func (_m *MockPDU) Update(pci Arg) error {
	ret := _m.Called(pci)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Arg) error); ok {
		r0 = rf(pci)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPDU_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockPDU_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - pci Arg
func (_e *MockPDU_Expecter) Update(pci interface{}) *MockPDU_Update_Call {
	return &MockPDU_Update_Call{Call: _e.mock.On("Update", pci)}
}

func (_c *MockPDU_Update_Call) Run(run func(pci Arg)) *MockPDU_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Arg))
	})
	return _c
}

func (_c *MockPDU_Update_Call) Return(_a0 error) *MockPDU_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPDU_Update_Call) RunAndReturn(run func(Arg) error) *MockPDU_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPDU creates a new instance of MockPDU. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPDU(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPDU {
	mock := &MockPDU{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
