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

// BACnetNotificationParameters is the corresponding interface of BACnetNotificationParameters
type BACnetNotificationParameters interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetNotificationParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParameters.
// This is useful for switch cases.
type BACnetNotificationParametersExactly interface {
	BACnetNotificationParameters
	isBACnetNotificationParameters() bool
}

// _BACnetNotificationParameters is the data-structure of this message
type _BACnetNotificationParameters struct {
	_BACnetNotificationParametersChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

type _BACnetNotificationParametersChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetNotificationParametersParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetNotificationParameters, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetNotificationParametersChild interface {
	utils.Serializable
	InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetNotificationParameters

	GetTypeName() string
	BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetNotificationParameters) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetNotificationParameters) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetNotificationParameters) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParameters factory function for _BACnetNotificationParameters
func NewBACnetNotificationParameters(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParameters {
	return &_BACnetNotificationParameters{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParameters(structType interface{}) BACnetNotificationParameters {
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParameters) GetTypeName() string {
	return "BACnetNotificationParameters"
}

func (m *_BACnetNotificationParameters) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParameters) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParameters, error) {
	return BACnetNotificationParametersParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument) // TODO: get endianness from mspec
}

func BACnetNotificationParametersParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetNotificationParameters")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetNotificationParametersChildSerializeRequirement interface {
		BACnetNotificationParameters
		InitializeParent(BACnetNotificationParameters, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetNotificationParameters
	}
	var _childTemp interface{}
	var _child BACnetNotificationParametersChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfBitString
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfBitStringParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfState
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfStateParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(2): // BACnetNotificationParametersChangeOfValue
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfValueParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(3): // BACnetNotificationParametersCommandFailure
		_childTemp, typeSwitchError = BACnetNotificationParametersCommandFailureParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(4): // BACnetNotificationParametersFloatingLimit
		_childTemp, typeSwitchError = BACnetNotificationParametersFloatingLimitParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(5): // BACnetNotificationParametersOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersOutOfRangeParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(6): // BACnetNotificationParametersComplexEventType
		_childTemp, typeSwitchError = BACnetNotificationParametersComplexEventTypeParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(8): // BACnetNotificationParametersChangeOfLifeSafety
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfLifeSafetyParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(9): // BACnetNotificationParametersExtended
		_childTemp, typeSwitchError = BACnetNotificationParametersExtendedParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(10): // BACnetNotificationParametersBufferReady
		_childTemp, typeSwitchError = BACnetNotificationParametersBufferReadyParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(11): // BACnetNotificationParametersUnsignedRange
		_childTemp, typeSwitchError = BACnetNotificationParametersUnsignedRangeParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(13): // BACnetNotificationParametersAccessEvent
		_childTemp, typeSwitchError = BACnetNotificationParametersAccessEventParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(14): // BACnetNotificationParametersDoubleOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersDoubleOutOfRangeParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(15): // BACnetNotificationParametersSignedOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersSignedOutOfRangeParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(16): // BACnetNotificationParametersUnsignedOutOfRange
		_childTemp, typeSwitchError = BACnetNotificationParametersUnsignedOutOfRangeParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(17): // BACnetNotificationParametersChangeOfCharacterString
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfCharacterStringParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(18): // BACnetNotificationParametersChangeOfStatusFlags
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfStatusFlagsParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(19): // BACnetNotificationParametersChangeOfReliability
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfReliabilityParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(21): // BACnetNotificationParametersChangeOfDiscreteValue
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	case peekedTagNumber == uint8(22): // BACnetNotificationParametersChangeOfTimer
		_childTemp, typeSwitchError = BACnetNotificationParametersChangeOfTimerParseWithBuffer(readBuffer, tagNumber, objectTypeArgument, peekedTagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetNotificationParameters")
	}
	_child = _childTemp.(BACnetNotificationParametersChildSerializeRequirement)

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetNotificationParameters")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParameters")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetNotificationParameters) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetNotificationParameters, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNotificationParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParameters")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotificationParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNotificationParameters) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNotificationParameters) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetNotificationParameters) isBACnetNotificationParameters() bool {
	return true
}

func (m *_BACnetNotificationParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}