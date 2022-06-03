/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetAuthenticationPolicyListEntry is the data-structure of this message
type BACnetAuthenticationPolicyListEntry struct {
	CredentialDataInput *BACnetDeviceObjectReferenceEnclosed
	Index               *BACnetContextTagUnsignedInteger
}

// IBACnetAuthenticationPolicyListEntry is the corresponding interface of BACnetAuthenticationPolicyListEntry
type IBACnetAuthenticationPolicyListEntry interface {
	// GetCredentialDataInput returns CredentialDataInput (property field)
	GetCredentialDataInput() *BACnetDeviceObjectReferenceEnclosed
	// GetIndex returns Index (property field)
	GetIndex() *BACnetContextTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetAuthenticationPolicyListEntry) GetCredentialDataInput() *BACnetDeviceObjectReferenceEnclosed {
	return m.CredentialDataInput
}

func (m *BACnetAuthenticationPolicyListEntry) GetIndex() *BACnetContextTagUnsignedInteger {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationPolicyListEntry factory function for BACnetAuthenticationPolicyListEntry
func NewBACnetAuthenticationPolicyListEntry(credentialDataInput *BACnetDeviceObjectReferenceEnclosed, index *BACnetContextTagUnsignedInteger) *BACnetAuthenticationPolicyListEntry {
	return &BACnetAuthenticationPolicyListEntry{CredentialDataInput: credentialDataInput, Index: index}
}

func CastBACnetAuthenticationPolicyListEntry(structType interface{}) *BACnetAuthenticationPolicyListEntry {
	if casted, ok := structType.(BACnetAuthenticationPolicyListEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicyListEntry); ok {
		return casted
	}
	return nil
}

func (m *BACnetAuthenticationPolicyListEntry) GetTypeName() string {
	return "BACnetAuthenticationPolicyListEntry"
}

func (m *BACnetAuthenticationPolicyListEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetAuthenticationPolicyListEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (credentialDataInput)
	lengthInBits += m.CredentialDataInput.GetLengthInBits()

	// Simple field (index)
	lengthInBits += m.Index.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetAuthenticationPolicyListEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAuthenticationPolicyListEntryParse(readBuffer utils.ReadBuffer) (*BACnetAuthenticationPolicyListEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicyListEntry"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (credentialDataInput)
	if pullErr := readBuffer.PullContext("credentialDataInput"); pullErr != nil {
		return nil, pullErr
	}
	_credentialDataInput, _credentialDataInputErr := BACnetDeviceObjectReferenceEnclosedParse(readBuffer, uint8(uint8(0)))
	if _credentialDataInputErr != nil {
		return nil, errors.Wrap(_credentialDataInputErr, "Error parsing 'credentialDataInput' field")
	}
	credentialDataInput := CastBACnetDeviceObjectReferenceEnclosed(_credentialDataInput)
	if closeErr := readBuffer.CloseContext("credentialDataInput"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (index)
	if pullErr := readBuffer.PullContext("index"); pullErr != nil {
		return nil, pullErr
	}
	_index, _indexErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field")
	}
	index := CastBACnetContextTagUnsignedInteger(_index)
	if closeErr := readBuffer.CloseContext("index"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicyListEntry"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetAuthenticationPolicyListEntry(credentialDataInput, index), nil
}

func (m *BACnetAuthenticationPolicyListEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicyListEntry"); pushErr != nil {
		return pushErr
	}

	// Simple Field (credentialDataInput)
	if pushErr := writeBuffer.PushContext("credentialDataInput"); pushErr != nil {
		return pushErr
	}
	_credentialDataInputErr := m.CredentialDataInput.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("credentialDataInput"); popErr != nil {
		return popErr
	}
	if _credentialDataInputErr != nil {
		return errors.Wrap(_credentialDataInputErr, "Error serializing 'credentialDataInput' field")
	}

	// Simple Field (index)
	if pushErr := writeBuffer.PushContext("index"); pushErr != nil {
		return pushErr
	}
	_indexErr := m.Index.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("index"); popErr != nil {
		return popErr
	}
	if _indexErr != nil {
		return errors.Wrap(_indexErr, "Error serializing 'index' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicyListEntry"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetAuthenticationPolicyListEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
