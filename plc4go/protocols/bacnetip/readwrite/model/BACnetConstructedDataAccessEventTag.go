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

// BACnetConstructedDataAccessEventTag is the data-structure of this message
type BACnetConstructedDataAccessEventTag struct {
	*BACnetConstructedData
	AccessEventTag *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAccessEventTag is the corresponding interface of BACnetConstructedDataAccessEventTag
type IBACnetConstructedDataAccessEventTag interface {
	IBACnetConstructedData
	// GetAccessEventTag returns AccessEventTag (property field)
	GetAccessEventTag() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataAccessEventTag) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAccessEventTag) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_TAG
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccessEventTag) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccessEventTag) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccessEventTag) GetAccessEventTag() *BACnetApplicationTagUnsignedInteger {
	return m.AccessEventTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventTag factory function for BACnetConstructedDataAccessEventTag
func NewBACnetConstructedDataAccessEventTag(accessEventTag *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAccessEventTag {
	_result := &BACnetConstructedDataAccessEventTag{
		AccessEventTag:        accessEventTag,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccessEventTag(structType interface{}) *BACnetConstructedDataAccessEventTag {
	if casted, ok := structType.(BACnetConstructedDataAccessEventTag); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventTag); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccessEventTag(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccessEventTag(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccessEventTag) GetTypeName() string {
	return "BACnetConstructedDataAccessEventTag"
}

func (m *BACnetConstructedDataAccessEventTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccessEventTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessEventTag)
	lengthInBits += m.AccessEventTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAccessEventTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessEventTagParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAccessEventTag, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventTag"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessEventTag)
	if pullErr := readBuffer.PullContext("accessEventTag"); pullErr != nil {
		return nil, pullErr
	}
	_accessEventTag, _accessEventTagErr := BACnetApplicationTagParse(readBuffer)
	if _accessEventTagErr != nil {
		return nil, errors.Wrap(_accessEventTagErr, "Error parsing 'accessEventTag' field")
	}
	accessEventTag := CastBACnetApplicationTagUnsignedInteger(_accessEventTag)
	if closeErr := readBuffer.CloseContext("accessEventTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventTag"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccessEventTag{
		AccessEventTag:        CastBACnetApplicationTagUnsignedInteger(accessEventTag),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccessEventTag) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventTag"); pushErr != nil {
			return pushErr
		}

		// Simple Field (accessEventTag)
		if pushErr := writeBuffer.PushContext("accessEventTag"); pushErr != nil {
			return pushErr
		}
		_accessEventTagErr := m.AccessEventTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("accessEventTag"); popErr != nil {
			return popErr
		}
		if _accessEventTagErr != nil {
			return errors.Wrap(_accessEventTagErr, "Error serializing 'accessEventTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventTag"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccessEventTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
