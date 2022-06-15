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

// NotTransmittedTooLong is the corresponding interface of NotTransmittedTooLong
type NotTransmittedTooLong interface {
	Confirmation
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _NotTransmittedTooLong is the data-structure of this message
type _NotTransmittedTooLong struct {
	*_Confirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NotTransmittedTooLong) GetConfirmationType() byte {
	return 0x27
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NotTransmittedTooLong) InitializeParent(parent Confirmation, alpha Alpha) {
	m.Alpha = alpha
}

func (m *_NotTransmittedTooLong) GetParent() Confirmation {
	return m._Confirmation
}

// NewNotTransmittedTooLong factory function for _NotTransmittedTooLong
func NewNotTransmittedTooLong(alpha Alpha) *_NotTransmittedTooLong {
	_result := &_NotTransmittedTooLong{
		_Confirmation: NewConfirmation(alpha),
	}
	_result._Confirmation._ConfirmationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNotTransmittedTooLong(structType interface{}) NotTransmittedTooLong {
	if casted, ok := structType.(NotTransmittedTooLong); ok {
		return casted
	}
	if casted, ok := structType.(*NotTransmittedTooLong); ok {
		return *casted
	}
	return nil
}

func (m *_NotTransmittedTooLong) GetTypeName() string {
	return "NotTransmittedTooLong"
}

func (m *_NotTransmittedTooLong) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NotTransmittedTooLong) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_NotTransmittedTooLong) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NotTransmittedTooLongParse(readBuffer utils.ReadBuffer) (NotTransmittedTooLong, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NotTransmittedTooLong"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NotTransmittedTooLong")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NotTransmittedTooLong"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NotTransmittedTooLong")
	}

	// Create a partially initialized instance
	_child := &_NotTransmittedTooLong{
		_Confirmation: &_Confirmation{},
	}
	_child._Confirmation._ConfirmationChildRequirements = _child
	return _child, nil
}

func (m *_NotTransmittedTooLong) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotTransmittedTooLong"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NotTransmittedTooLong")
		}

		if popErr := writeBuffer.PopContext("NotTransmittedTooLong"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NotTransmittedTooLong")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NotTransmittedTooLong) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
