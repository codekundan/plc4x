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

// S7Payload is the corresponding interface of S7Payload
type S7Payload interface {
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
	// GetParameterParameterType returns ParameterParameterType (discriminator field)
	GetParameterParameterType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _S7Payload is the data-structure of this message
type _S7Payload struct {
	_S7PayloadChildRequirements

	// Arguments.
	Parameter S7Parameter
}

type _S7PayloadChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetParameterParameterType() uint8
	GetMessageType() uint8
}

type S7PayloadParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child S7Payload, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7PayloadChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent S7Payload)
	GetParent() *S7Payload

	GetTypeName() string
	S7Payload
}

// NewS7Payload factory function for _S7Payload
func NewS7Payload(parameter S7Parameter) *_S7Payload {
	return &_S7Payload{Parameter: parameter}
}

// Deprecated: use the interface for direct cast
func CastS7Payload(structType interface{}) S7Payload {
	if casted, ok := structType.(S7Payload); ok {
		return casted
	}
	if casted, ok := structType.(*S7Payload); ok {
		return *casted
	}
	return nil
}

func (m *_S7Payload) GetTypeName() string {
	return "S7Payload"
}

func (m *_S7Payload) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_S7Payload) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadParse(readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (S7Payload, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Payload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7PayloadChildSerializeRequirement interface {
		S7Payload
		InitializeParent(S7Payload)
		GetParent() S7Payload
	}
	var _childTemp interface{}
	var _child S7PayloadChildSerializeRequirement
	var typeSwitchError error
	switch {
	case CastS7Parameter(parameter).GetParameterType() == 0x04 && messageType == 0x03: // S7PayloadReadVarResponse
		_childTemp, typeSwitchError = S7PayloadReadVarResponseParse(readBuffer, messageType, parameter)
		_child = _childTemp.(S7PayloadChildSerializeRequirement)
	case CastS7Parameter(parameter).GetParameterType() == 0x05 && messageType == 0x01: // S7PayloadWriteVarRequest
		_childTemp, typeSwitchError = S7PayloadWriteVarRequestParse(readBuffer, messageType, parameter)
		_child = _childTemp.(S7PayloadChildSerializeRequirement)
	case CastS7Parameter(parameter).GetParameterType() == 0x05 && messageType == 0x03: // S7PayloadWriteVarResponse
		_childTemp, typeSwitchError = S7PayloadWriteVarResponseParse(readBuffer, messageType, parameter)
		_child = _childTemp.(S7PayloadChildSerializeRequirement)
	case CastS7Parameter(parameter).GetParameterType() == 0x00 && messageType == 0x07: // S7PayloadUserData
		_childTemp, typeSwitchError = S7PayloadUserDataParse(readBuffer, messageType, parameter)
		_child = _childTemp.(S7PayloadChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7Payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Payload")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (m *_S7Payload) Serialize(writeBuffer utils.WriteBuffer) error {
	panic("Required method Serialize not implemented")
}

func (pm *_S7Payload) SerializeParent(writeBuffer utils.WriteBuffer, child S7Payload, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7Payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Payload")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Payload")
	}
	return nil
}

func (m *_S7Payload) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
