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

// BACnetConstructedDataProtocolServicesSupported is the corresponding interface of BACnetConstructedDataProtocolServicesSupported
type BACnetConstructedDataProtocolServicesSupported interface {
	BACnetConstructedData
	// GetProtocolServicesSupported returns ProtocolServicesSupported (property field)
	GetProtocolServicesSupported() BACnetServicesSupportedTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetServicesSupportedTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataProtocolServicesSupported is the data-structure of this message
type _BACnetConstructedDataProtocolServicesSupported struct {
	*_BACnetConstructedData
	ProtocolServicesSupported BACnetServicesSupportedTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProtocolServicesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProtocolServicesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_SERVICES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProtocolServicesSupported) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProtocolServicesSupported) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolServicesSupported) GetProtocolServicesSupported() BACnetServicesSupportedTagged {
	return m.ProtocolServicesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProtocolServicesSupported) GetActualValue() BACnetServicesSupportedTagged {
	return CastBACnetServicesSupportedTagged(m.GetProtocolServicesSupported())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolServicesSupported factory function for _BACnetConstructedDataProtocolServicesSupported
func NewBACnetConstructedDataProtocolServicesSupported(protocolServicesSupported BACnetServicesSupportedTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProtocolServicesSupported {
	_result := &_BACnetConstructedDataProtocolServicesSupported{
		ProtocolServicesSupported: protocolServicesSupported,
		_BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProtocolServicesSupported(structType interface{}) BACnetConstructedDataProtocolServicesSupported {
	if casted, ok := structType.(BACnetConstructedDataProtocolServicesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolServicesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProtocolServicesSupported) GetTypeName() string {
	return "BACnetConstructedDataProtocolServicesSupported"
}

func (m *_BACnetConstructedDataProtocolServicesSupported) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProtocolServicesSupported) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (protocolServicesSupported)
	lengthInBits += m.ProtocolServicesSupported.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProtocolServicesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProtocolServicesSupportedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProtocolServicesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolServicesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProtocolServicesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolServicesSupported)
	if pullErr := readBuffer.PullContext("protocolServicesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for protocolServicesSupported")
	}
	_protocolServicesSupported, _protocolServicesSupportedErr := BACnetServicesSupportedTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _protocolServicesSupportedErr != nil {
		return nil, errors.Wrap(_protocolServicesSupportedErr, "Error parsing 'protocolServicesSupported' field")
	}
	protocolServicesSupported := _protocolServicesSupported.(BACnetServicesSupportedTagged)
	if closeErr := readBuffer.CloseContext("protocolServicesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for protocolServicesSupported")
	}

	// Virtual field
	_actualValue := protocolServicesSupported
	actualValue := _actualValue.(BACnetServicesSupportedTagged)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolServicesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProtocolServicesSupported")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProtocolServicesSupported{
		ProtocolServicesSupported: protocolServicesSupported,
		_BACnetConstructedData:    &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProtocolServicesSupported) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolServicesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProtocolServicesSupported")
		}

		// Simple Field (protocolServicesSupported)
		if pushErr := writeBuffer.PushContext("protocolServicesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for protocolServicesSupported")
		}
		_protocolServicesSupportedErr := writeBuffer.WriteSerializable(m.GetProtocolServicesSupported())
		if popErr := writeBuffer.PopContext("protocolServicesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for protocolServicesSupported")
		}
		if _protocolServicesSupportedErr != nil {
			return errors.Wrap(_protocolServicesSupportedErr, "Error serializing 'protocolServicesSupported' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolServicesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProtocolServicesSupported")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProtocolServicesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
