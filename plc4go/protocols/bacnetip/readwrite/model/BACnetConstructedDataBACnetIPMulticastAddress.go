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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataBACnetIPMulticastAddress is the corresponding interface of BACnetConstructedDataBACnetIPMulticastAddress
type BACnetConstructedDataBACnetIPMulticastAddress interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpMulticastAddress returns IpMulticastAddress (property field)
	GetIpMulticastAddress() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataBACnetIPMulticastAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPMulticastAddress.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPMulticastAddressExactly interface {
	BACnetConstructedDataBACnetIPMulticastAddress
	isBACnetConstructedDataBACnetIPMulticastAddress() bool
}

// _BACnetConstructedDataBACnetIPMulticastAddress is the data-structure of this message
type _BACnetConstructedDataBACnetIPMulticastAddress struct {
	*_BACnetConstructedData
	IpMulticastAddress BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_MULTICAST_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetIpMulticastAddress() BACnetApplicationTagOctetString {
	return m.IpMulticastAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetActualValue() BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetIpMulticastAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPMulticastAddress factory function for _BACnetConstructedDataBACnetIPMulticastAddress
func NewBACnetConstructedDataBACnetIPMulticastAddress(ipMulticastAddress BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPMulticastAddress {
	_result := &_BACnetConstructedDataBACnetIPMulticastAddress{
		IpMulticastAddress:     ipMulticastAddress,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPMulticastAddress(structType interface{}) BACnetConstructedDataBACnetIPMulticastAddress {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPMulticastAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPMulticastAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPMulticastAddress"
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipMulticastAddress)
	lengthInBits += m.IpMulticastAddress.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBACnetIPMulticastAddressParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPMulticastAddress, error) {
	return BACnetConstructedDataBACnetIPMulticastAddressParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataBACnetIPMulticastAddressParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPMulticastAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPMulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPMulticastAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipMulticastAddress)
	if pullErr := readBuffer.PullContext("ipMulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipMulticastAddress")
	}
	_ipMulticastAddress, _ipMulticastAddressErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _ipMulticastAddressErr != nil {
		return nil, errors.Wrap(_ipMulticastAddressErr, "Error parsing 'ipMulticastAddress' field of BACnetConstructedDataBACnetIPMulticastAddress")
	}
	ipMulticastAddress := _ipMulticastAddress.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("ipMulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipMulticastAddress")
	}

	// Virtual field
	_actualValue := ipMulticastAddress
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPMulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPMulticastAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPMulticastAddress{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IpMulticastAddress: ipMulticastAddress,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPMulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPMulticastAddress")
		}

		// Simple Field (ipMulticastAddress)
		if pushErr := writeBuffer.PushContext("ipMulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipMulticastAddress")
		}
		_ipMulticastAddressErr := writeBuffer.WriteSerializable(m.GetIpMulticastAddress())
		if popErr := writeBuffer.PopContext("ipMulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipMulticastAddress")
		}
		if _ipMulticastAddressErr != nil {
			return errors.Wrap(_ipMulticastAddressErr, "Error serializing 'ipMulticastAddress' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPMulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPMulticastAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) isBACnetConstructedDataBACnetIPMulticastAddress() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPMulticastAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}