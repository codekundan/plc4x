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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadAlarm8 is the corresponding interface of S7PayloadAlarm8
type S7PayloadAlarm8 interface {
	S7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() AlarmMessagePushType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _S7PayloadAlarm8 is the data-structure of this message
type _S7PayloadAlarm8 struct {
	*_S7PayloadUserDataItem
	AlarmMessage AlarmMessagePushType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadAlarm8) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadAlarm8) GetCpuSubfunction() uint8 {
	return 0x05
}

func (m *_S7PayloadAlarm8) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadAlarm8) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadAlarm8) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadAlarm8) GetAlarmMessage() AlarmMessagePushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadAlarm8 factory function for _S7PayloadAlarm8
func NewS7PayloadAlarm8(alarmMessage AlarmMessagePushType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7PayloadAlarm8 {
	_result := &_S7PayloadAlarm8{
		AlarmMessage:           alarmMessage,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadAlarm8(structType interface{}) S7PayloadAlarm8 {
	if casted, ok := structType.(S7PayloadAlarm8); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadAlarm8); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadAlarm8) GetTypeName() string {
	return "S7PayloadAlarm8"
}

func (m *_S7PayloadAlarm8) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7PayloadAlarm8) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits()

	return lengthInBits
}

func (m *_S7PayloadAlarm8) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadAlarm8Parse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadAlarm8, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadAlarm8"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadAlarm8")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmMessage")
	}
	_alarmMessage, _alarmMessageErr := AlarmMessagePushTypeParse(readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field")
	}
	alarmMessage := _alarmMessage.(AlarmMessagePushType)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmMessage")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadAlarm8"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadAlarm8")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadAlarm8{
		AlarmMessage:           alarmMessage,
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadAlarm8) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadAlarm8"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadAlarm8")
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmMessage")
		}
		_alarmMessageErr := writeBuffer.WriteSerializable(m.GetAlarmMessage())
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmMessage")
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadAlarm8"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadAlarm8")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7PayloadAlarm8) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
