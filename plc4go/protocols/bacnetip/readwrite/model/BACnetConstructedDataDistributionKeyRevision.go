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

// BACnetConstructedDataDistributionKeyRevision is the corresponding interface of BACnetConstructedDataDistributionKeyRevision
type BACnetConstructedDataDistributionKeyRevision interface {
	BACnetConstructedData
	// GetDistributionKeyRevision returns DistributionKeyRevision (property field)
	GetDistributionKeyRevision() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataDistributionKeyRevision is the data-structure of this message
type _BACnetConstructedDataDistributionKeyRevision struct {
	*_BACnetConstructedData
	DistributionKeyRevision BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DISTRIBUTION_KEY_REVISION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetDistributionKeyRevision() BACnetApplicationTagUnsignedInteger {
	return m.DistributionKeyRevision
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDistributionKeyRevision) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetDistributionKeyRevision())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDistributionKeyRevision factory function for _BACnetConstructedDataDistributionKeyRevision
func NewBACnetConstructedDataDistributionKeyRevision(distributionKeyRevision BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDistributionKeyRevision {
	_result := &_BACnetConstructedDataDistributionKeyRevision{
		DistributionKeyRevision: distributionKeyRevision,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDistributionKeyRevision(structType interface{}) BACnetConstructedDataDistributionKeyRevision {
	if casted, ok := structType.(BACnetConstructedDataDistributionKeyRevision); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDistributionKeyRevision); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetTypeName() string {
	return "BACnetConstructedDataDistributionKeyRevision"
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (distributionKeyRevision)
	lengthInBits += m.DistributionKeyRevision.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDistributionKeyRevision) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDistributionKeyRevisionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDistributionKeyRevision, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDistributionKeyRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDistributionKeyRevision")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (distributionKeyRevision)
	if pullErr := readBuffer.PullContext("distributionKeyRevision"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for distributionKeyRevision")
	}
	_distributionKeyRevision, _distributionKeyRevisionErr := BACnetApplicationTagParse(readBuffer)
	if _distributionKeyRevisionErr != nil {
		return nil, errors.Wrap(_distributionKeyRevisionErr, "Error parsing 'distributionKeyRevision' field")
	}
	distributionKeyRevision := _distributionKeyRevision.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("distributionKeyRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for distributionKeyRevision")
	}

	// Virtual field
	_actualValue := distributionKeyRevision
	actualValue := _actualValue.(BACnetApplicationTagUnsignedInteger)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDistributionKeyRevision"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDistributionKeyRevision")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDistributionKeyRevision{
		DistributionKeyRevision: distributionKeyRevision,
		_BACnetConstructedData:  &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDistributionKeyRevision) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDistributionKeyRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDistributionKeyRevision")
		}

		// Simple Field (distributionKeyRevision)
		if pushErr := writeBuffer.PushContext("distributionKeyRevision"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for distributionKeyRevision")
		}
		_distributionKeyRevisionErr := writeBuffer.WriteSerializable(m.GetDistributionKeyRevision())
		if popErr := writeBuffer.PopContext("distributionKeyRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for distributionKeyRevision")
		}
		if _distributionKeyRevisionErr != nil {
			return errors.Wrap(_distributionKeyRevisionErr, "Error serializing 'distributionKeyRevision' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDistributionKeyRevision"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDistributionKeyRevision")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDistributionKeyRevision) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
