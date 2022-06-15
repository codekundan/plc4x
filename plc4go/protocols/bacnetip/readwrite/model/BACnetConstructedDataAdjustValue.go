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

// BACnetConstructedDataAdjustValue is the corresponding interface of BACnetConstructedDataAdjustValue
type BACnetConstructedDataAdjustValue interface {
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataAdjustValue is the data-structure of this message
type _BACnetConstructedDataAdjustValue struct {
	*_BACnetConstructedData
	AdjustValue BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAdjustValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAdjustValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetAdjustValue() BACnetApplicationTagSignedInteger {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetActualValue() BACnetApplicationTagSignedInteger {
	return CastBACnetApplicationTagSignedInteger(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAdjustValue factory function for _BACnetConstructedDataAdjustValue
func NewBACnetConstructedDataAdjustValue(adjustValue BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAdjustValue {
	_result := &_BACnetConstructedDataAdjustValue{
		AdjustValue:            adjustValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAdjustValue(structType interface{}) BACnetConstructedDataAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataAdjustValue"
}

func (m *_BACnetConstructedDataAdjustValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAdjustValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAdjustValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAdjustValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAdjustValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (adjustValue)
	if pullErr := readBuffer.PullContext("adjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for adjustValue")
	}
	_adjustValue, _adjustValueErr := BACnetApplicationTagParse(readBuffer)
	if _adjustValueErr != nil {
		return nil, errors.Wrap(_adjustValueErr, "Error parsing 'adjustValue' field")
	}
	adjustValue := _adjustValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("adjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for adjustValue")
	}

	// Virtual field
	_actualValue := adjustValue
	actualValue := _actualValue.(BACnetApplicationTagSignedInteger)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAdjustValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAdjustValue{
		AdjustValue:            adjustValue,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAdjustValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAdjustValue")
		}

		// Simple Field (adjustValue)
		if pushErr := writeBuffer.PushContext("adjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for adjustValue")
		}
		_adjustValueErr := writeBuffer.WriteSerializable(m.GetAdjustValue())
		if popErr := writeBuffer.PopContext("adjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for adjustValue")
		}
		if _adjustValueErr != nil {
			return errors.Wrap(_adjustValueErr, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAdjustValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
