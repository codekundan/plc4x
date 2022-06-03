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

// BACnetConstructedDataOccupancyUpperLimitEnforced is the data-structure of this message
type BACnetConstructedDataOccupancyUpperLimitEnforced struct {
	*BACnetConstructedData
	OccupancyUpperLimitEnforced *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataOccupancyUpperLimitEnforced is the corresponding interface of BACnetConstructedDataOccupancyUpperLimitEnforced
type IBACnetConstructedDataOccupancyUpperLimitEnforced interface {
	IBACnetConstructedData
	// GetOccupancyUpperLimitEnforced returns OccupancyUpperLimitEnforced (property field)
	GetOccupancyUpperLimitEnforced() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_UPPER_LIMIT_ENFORCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetOccupancyUpperLimitEnforced() *BACnetApplicationTagBoolean {
	return m.OccupancyUpperLimitEnforced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyUpperLimitEnforced factory function for BACnetConstructedDataOccupancyUpperLimitEnforced
func NewBACnetConstructedDataOccupancyUpperLimitEnforced(occupancyUpperLimitEnforced *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataOccupancyUpperLimitEnforced {
	_result := &BACnetConstructedDataOccupancyUpperLimitEnforced{
		OccupancyUpperLimitEnforced: occupancyUpperLimitEnforced,
		BACnetConstructedData:       NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataOccupancyUpperLimitEnforced(structType interface{}) *BACnetConstructedDataOccupancyUpperLimitEnforced {
	if casted, ok := structType.(BACnetConstructedDataOccupancyUpperLimitEnforced); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyUpperLimitEnforced); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyUpperLimitEnforced(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyUpperLimitEnforced(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetTypeName() string {
	return "BACnetConstructedDataOccupancyUpperLimitEnforced"
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (occupancyUpperLimitEnforced)
	lengthInBits += m.OccupancyUpperLimitEnforced.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOccupancyUpperLimitEnforcedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataOccupancyUpperLimitEnforced, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyUpperLimitEnforced)
	if pullErr := readBuffer.PullContext("occupancyUpperLimitEnforced"); pullErr != nil {
		return nil, pullErr
	}
	_occupancyUpperLimitEnforced, _occupancyUpperLimitEnforcedErr := BACnetApplicationTagParse(readBuffer)
	if _occupancyUpperLimitEnforcedErr != nil {
		return nil, errors.Wrap(_occupancyUpperLimitEnforcedErr, "Error parsing 'occupancyUpperLimitEnforced' field")
	}
	occupancyUpperLimitEnforced := CastBACnetApplicationTagBoolean(_occupancyUpperLimitEnforced)
	if closeErr := readBuffer.CloseContext("occupancyUpperLimitEnforced"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataOccupancyUpperLimitEnforced{
		OccupancyUpperLimitEnforced: CastBACnetApplicationTagBoolean(occupancyUpperLimitEnforced),
		BACnetConstructedData:       &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); pushErr != nil {
			return pushErr
		}

		// Simple Field (occupancyUpperLimitEnforced)
		if pushErr := writeBuffer.PushContext("occupancyUpperLimitEnforced"); pushErr != nil {
			return pushErr
		}
		_occupancyUpperLimitEnforcedErr := m.OccupancyUpperLimitEnforced.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("occupancyUpperLimitEnforced"); popErr != nil {
			return popErr
		}
		if _occupancyUpperLimitEnforcedErr != nil {
			return errors.Wrap(_occupancyUpperLimitEnforcedErr, "Error serializing 'occupancyUpperLimitEnforced' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataOccupancyUpperLimitEnforced) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
