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

// BACnetContextTag is the corresponding interface of BACnetContextTag
type BACnetContextTag interface {
	// GetDataType returns DataType (discriminator field)
	GetDataType() BACnetDataType
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetTagNumber returns TagNumber (virtual field)
	GetTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetContextTag is the data-structure of this message
type _BACnetContextTag struct {
	_BACnetContextTagChildRequirements
	Header BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

type _BACnetContextTagChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetDataType() BACnetDataType
}

type BACnetContextTagParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetContextTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetContextTagChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent BACnetContextTag, header BACnetTagHeader)
	GetParent() *BACnetContextTag

	GetTypeName() string
	BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTag) GetTagNumber() uint8 {
	return uint8(m.GetHeader().GetTagNumber())
}

func (m *_BACnetContextTag) GetActualLength() uint32 {
	return uint32(m.GetHeader().GetActualLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTag factory function for _BACnetContextTag
func NewBACnetContextTag(header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTag {
	return &_BACnetContextTag{Header: header, TagNumberArgument: tagNumberArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTag(structType interface{}) BACnetContextTag {
	if casted, ok := structType.(BACnetContextTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTag) GetTypeName() string {
	return "BACnetContextTag"
}

func (m *_BACnetContextTag) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTag, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{"should be a context tag"})
	}

	// Virtual field
	_tagNumber := header.GetTagNumber()
	tagNumber := uint8(_tagNumber)
	_ = tagNumber

	// Virtual field
	_actualLength := header.GetActualLength()
	actualLength := uint32(_actualLength)
	_ = actualLength

	// Validation
	if !(bool(bool((header.GetLengthValueType()) != (6))) && bool(bool((header.GetLengthValueType()) != (7)))) {
		return nil, errors.WithStack(utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetContextTagChildSerializeRequirement interface {
		BACnetContextTag
		InitializeParent(BACnetContextTag, BACnetTagHeader)
		GetParent() BACnetContextTag
	}
	var _childTemp interface{}
	var _child BACnetContextTagChildSerializeRequirement
	var typeSwitchError error
	switch {
	case dataType == BACnetDataType_NULL: // BACnetContextTagNull
		_childTemp, typeSwitchError = BACnetContextTagNullParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_BOOLEAN: // BACnetContextTagBoolean
		_childTemp, typeSwitchError = BACnetContextTagBooleanParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_UNSIGNED_INTEGER: // BACnetContextTagUnsignedInteger
		_childTemp, typeSwitchError = BACnetContextTagUnsignedIntegerParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_SIGNED_INTEGER: // BACnetContextTagSignedInteger
		_childTemp, typeSwitchError = BACnetContextTagSignedIntegerParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_REAL: // BACnetContextTagReal
		_childTemp, typeSwitchError = BACnetContextTagRealParse(readBuffer, tagNumberArgument, dataType)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_DOUBLE: // BACnetContextTagDouble
		_childTemp, typeSwitchError = BACnetContextTagDoubleParse(readBuffer, tagNumberArgument, dataType)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_OCTET_STRING: // BACnetContextTagOctetString
		_childTemp, typeSwitchError = BACnetContextTagOctetStringParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_CHARACTER_STRING: // BACnetContextTagCharacterString
		_childTemp, typeSwitchError = BACnetContextTagCharacterStringParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_BIT_STRING: // BACnetContextTagBitString
		_childTemp, typeSwitchError = BACnetContextTagBitStringParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_ENUMERATED: // BACnetContextTagEnumerated
		_childTemp, typeSwitchError = BACnetContextTagEnumeratedParse(readBuffer, tagNumberArgument, dataType, header)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_DATE: // BACnetContextTagDate
		_childTemp, typeSwitchError = BACnetContextTagDateParse(readBuffer, tagNumberArgument, dataType)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_TIME: // BACnetContextTagTime
		_childTemp, typeSwitchError = BACnetContextTagTimeParse(readBuffer, tagNumberArgument, dataType)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_BACNET_OBJECT_IDENTIFIER: // BACnetContextTagObjectIdentifier
		_childTemp, typeSwitchError = BACnetContextTagObjectIdentifierParse(readBuffer, tagNumberArgument, dataType)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	case dataType == BACnetDataType_UNKNOWN: // BACnetContextTagUnknown
		_childTemp, typeSwitchError = BACnetContextTagUnknownParse(readBuffer, tagNumberArgument, dataType, actualLength)
		_child = _childTemp.(BACnetContextTagChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTag")
	}

	// Finish initializing
	_child.InitializeParent(_child, header)
	return _child, nil
}

func (m *_BACnetContextTag) Serialize(writeBuffer utils.WriteBuffer) error {
	panic("Required method Serialize not implemented")
}

func (pm *_BACnetContextTag) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetContextTag, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetContextTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetContextTag")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}
	// Virtual field
	if _tagNumberErr := writeBuffer.WriteVirtual("tagNumber", m.GetTagNumber()); _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual("actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetContextTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetContextTag")
	}
	return nil
}

func (m *_BACnetContextTag) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
