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

// MFuncPropStateReadReq is the corresponding interface of MFuncPropStateReadReq
type MFuncPropStateReadReq interface {
	CEMI
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _MFuncPropStateReadReq is the data-structure of this message
type _MFuncPropStateReadReq struct {
	*_CEMI

	// Arguments.
	Size uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropStateReadReq) GetMessageCode() uint8 {
	return 0xF9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropStateReadReq) InitializeParent(parent CEMI) {}

func (m *_MFuncPropStateReadReq) GetParent() CEMI {
	return m._CEMI
}

// NewMFuncPropStateReadReq factory function for _MFuncPropStateReadReq
func NewMFuncPropStateReadReq(size uint16) *_MFuncPropStateReadReq {
	_result := &_MFuncPropStateReadReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMFuncPropStateReadReq(structType interface{}) MFuncPropStateReadReq {
	if casted, ok := structType.(MFuncPropStateReadReq); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropStateReadReq); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropStateReadReq) GetTypeName() string {
	return "MFuncPropStateReadReq"
}

func (m *_MFuncPropStateReadReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MFuncPropStateReadReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_MFuncPropStateReadReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MFuncPropStateReadReqParse(readBuffer utils.ReadBuffer, size uint16) (MFuncPropStateReadReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropStateReadReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropStateReadReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropStateReadReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropStateReadReq")
	}

	// Create a partially initialized instance
	_child := &_MFuncPropStateReadReq{
		_CEMI: &_CEMI{},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MFuncPropStateReadReq) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropStateReadReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropStateReadReq")
		}

		if popErr := writeBuffer.PopContext("MFuncPropStateReadReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropStateReadReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MFuncPropStateReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
