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

// BACnetChannelValue is the corresponding interface of BACnetChannelValue
type BACnetChannelValue interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetChannelValue is the data-structure of this message
type _BACnetChannelValue struct {
	_BACnetChannelValueChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetChannelValueChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetChannelValueParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetChannelValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetChannelValueChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetChannelValue

	GetTypeName() string
	BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetChannelValue) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetChannelValue) GetPeekedIsContextTag() bool {
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValue factory function for _BACnetChannelValue
func NewBACnetChannelValue(peekedTagHeader BACnetTagHeader) *_BACnetChannelValue {
	return &_BACnetChannelValue{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValue(structType interface{}) BACnetChannelValue {
	if casted, ok := structType.(BACnetChannelValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValue) GetTypeName() string {
	return "BACnetChannelValue"
}

func (m *_BACnetChannelValue) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetChannelValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueParse(readBuffer utils.ReadBuffer) (BACnetChannelValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Virtual field
	_peekedIsContextTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))
	peekedIsContextTag := bool(_peekedIsContextTag)
	_ = peekedIsContextTag

	// Validation
	if !(bool(bool(!(peekedIsContextTag))) || bool(bool(bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetChannelValueChildSerializeRequirement interface {
		BACnetChannelValue
		InitializeParent(BACnetChannelValue, BACnetTagHeader)
		GetParent() BACnetChannelValue
	}
	var _childTemp interface{}
	var _child BACnetChannelValueChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetChannelValueNull
		_childTemp, typeSwitchError = BACnetChannelValueNullParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetChannelValueReal
		_childTemp, typeSwitchError = BACnetChannelValueRealParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetChannelValueEnumerated
		_childTemp, typeSwitchError = BACnetChannelValueEnumeratedParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetChannelValueUnsigned
		_childTemp, typeSwitchError = BACnetChannelValueUnsignedParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetChannelValueBoolean
		_childTemp, typeSwitchError = BACnetChannelValueBooleanParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetChannelValueInteger
		_childTemp, typeSwitchError = BACnetChannelValueIntegerParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetChannelValueDouble
		_childTemp, typeSwitchError = BACnetChannelValueDoubleParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetChannelValueTime
		_childTemp, typeSwitchError = BACnetChannelValueTimeParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetChannelValueCharacterString
		_childTemp, typeSwitchError = BACnetChannelValueCharacterStringParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetChannelValueOctetString
		_childTemp, typeSwitchError = BACnetChannelValueOctetStringParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetChannelValueBitString
		_childTemp, typeSwitchError = BACnetChannelValueBitStringParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetChannelValueDate
		_childTemp, typeSwitchError = BACnetChannelValueDateParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetChannelValueObjectidentifier
		_childTemp, typeSwitchError = BACnetChannelValueObjectidentifierParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetChannelValueLightingCommand
		_childTemp, typeSwitchError = BACnetChannelValueLightingCommandParse(readBuffer)
		_child = _childTemp.(BACnetChannelValueChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValue")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (m *_BACnetChannelValue) Serialize(writeBuffer utils.WriteBuffer) error {
	panic("Required method Serialize not implemented")
}

func (pm *_BACnetChannelValue) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetChannelValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetChannelValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetChannelValue")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual("peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetChannelValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetChannelValue")
	}
	return nil
}

func (m *_BACnetChannelValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
