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

// BACnetConstructedDataLogInterval is the corresponding interface of BACnetConstructedDataLogInterval
type BACnetConstructedDataLogInterval interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLogInterval returns LogInterval (property field)
	GetLogInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataLogIntervalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLogInterval.
// This is useful for switch cases.
type BACnetConstructedDataLogIntervalExactly interface {
	BACnetConstructedDataLogInterval
	isBACnetConstructedDataLogInterval() bool
}

// _BACnetConstructedDataLogInterval is the data-structure of this message
type _BACnetConstructedDataLogInterval struct {
	*_BACnetConstructedData
	LogInterval BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLogInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLogInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLogInterval) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLogInterval) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLogInterval) GetLogInterval() BACnetApplicationTagUnsignedInteger {
	return m.LogInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLogInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetLogInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLogInterval factory function for _BACnetConstructedDataLogInterval
func NewBACnetConstructedDataLogInterval(logInterval BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLogInterval {
	_result := &_BACnetConstructedDataLogInterval{
		LogInterval:            logInterval,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLogInterval(structType interface{}) BACnetConstructedDataLogInterval {
	if casted, ok := structType.(BACnetConstructedDataLogInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLogInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLogInterval) GetTypeName() string {
	return "BACnetConstructedDataLogInterval"
}

func (m *_BACnetConstructedDataLogInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLogInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (logInterval)
	lengthInBits += m.LogInterval.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLogInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLogIntervalParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLogInterval, error) {
	return BACnetConstructedDataLogIntervalParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataLogIntervalParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLogInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLogInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLogInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logInterval)
	if pullErr := readBuffer.PullContext("logInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logInterval")
	}
	_logInterval, _logIntervalErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _logIntervalErr != nil {
		return nil, errors.Wrap(_logIntervalErr, "Error parsing 'logInterval' field of BACnetConstructedDataLogInterval")
	}
	logInterval := _logInterval.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("logInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logInterval")
	}

	// Virtual field
	_actualValue := logInterval
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLogInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLogInterval")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLogInterval{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LogInterval: logInterval,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLogInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLogInterval) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLogInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLogInterval")
		}

		// Simple Field (logInterval)
		if pushErr := writeBuffer.PushContext("logInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logInterval")
		}
		_logIntervalErr := writeBuffer.WriteSerializable(m.GetLogInterval())
		if popErr := writeBuffer.PopContext("logInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logInterval")
		}
		if _logIntervalErr != nil {
			return errors.Wrap(_logIntervalErr, "Error serializing 'logInterval' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLogInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLogInterval")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLogInterval) isBACnetConstructedDataLogInterval() bool {
	return true
}

func (m *_BACnetConstructedDataLogInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}