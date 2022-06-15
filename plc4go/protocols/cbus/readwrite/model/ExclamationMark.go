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

// ExclamationMark is the corresponding interface of ExclamationMark
type ExclamationMark interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _ExclamationMark is the data-structure of this message
type _ExclamationMark struct {
}

// NewExclamationMark factory function for _ExclamationMark
func NewExclamationMark() *_ExclamationMark {
	return &_ExclamationMark{}
}

// Deprecated: use the interface for direct cast
func CastExclamationMark(structType interface{}) ExclamationMark {
	if casted, ok := structType.(ExclamationMark); ok {
		return casted
	}
	if casted, ok := structType.(*ExclamationMark); ok {
		return *casted
	}
	return nil
}

func (m *_ExclamationMark) GetTypeName() string {
	return "ExclamationMark"
}

func (m *_ExclamationMark) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ExclamationMark) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ExclamationMark) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ExclamationMarkParse(readBuffer utils.ReadBuffer) (ExclamationMark, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExclamationMark"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExclamationMark")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ExclamationMark"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExclamationMark")
	}

	// Create the instance
	return NewExclamationMark(), nil
}

func (m *_ExclamationMark) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ExclamationMark"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ExclamationMark")
	}

	if popErr := writeBuffer.PopContext("ExclamationMark"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ExclamationMark")
	}
	return nil
}

func (m *_ExclamationMark) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
