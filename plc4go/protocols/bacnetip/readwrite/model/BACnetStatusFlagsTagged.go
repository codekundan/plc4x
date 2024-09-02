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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetStatusFlagsTagged is the corresponding interface of BACnetStatusFlagsTagged
type BACnetStatusFlagsTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetInAlarm returns InAlarm (virtual field)
	GetInAlarm() bool
	// GetFault returns Fault (virtual field)
	GetFault() bool
	// GetOverridden returns Overridden (virtual field)
	GetOverridden() bool
	// GetOutOfService returns OutOfService (virtual field)
	GetOutOfService() bool
}

// BACnetStatusFlagsTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetStatusFlagsTagged.
// This is useful for switch cases.
type BACnetStatusFlagsTaggedExactly interface {
	BACnetStatusFlagsTagged
	isBACnetStatusFlagsTagged() bool
}

// _BACnetStatusFlagsTagged is the data-structure of this message
type _BACnetStatusFlagsTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetStatusFlagsTagged = (*_BACnetStatusFlagsTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetStatusFlagsTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetStatusFlagsTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetStatusFlagsTagged) GetInAlarm() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetStatusFlagsTagged) GetFault() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetStatusFlagsTagged) GetOverridden() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetStatusFlagsTagged) GetOutOfService() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (3))), func() any { return bool(m.GetPayload().GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetStatusFlagsTagged factory function for _BACnetStatusFlagsTagged
func NewBACnetStatusFlagsTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetStatusFlagsTagged {
	return &_BACnetStatusFlagsTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetStatusFlagsTagged(structType any) BACnetStatusFlagsTagged {
	if casted, ok := structType.(BACnetStatusFlagsTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetStatusFlagsTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetStatusFlagsTagged) GetTypeName() string {
	return "BACnetStatusFlagsTagged"
}

func (m *_BACnetStatusFlagsTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetStatusFlagsTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetStatusFlagsTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetStatusFlagsTagged, error) {
	return BACnetStatusFlagsTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetStatusFlagsTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetStatusFlagsTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetStatusFlagsTagged, error) {
		return BACnetStatusFlagsTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetStatusFlagsTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetStatusFlagsTagged, error) {
	v, err := (&_BACnetStatusFlagsTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetStatusFlagsTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetStatusFlagsTagged BACnetStatusFlagsTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetStatusFlagsTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetStatusFlagsTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	inAlarm, err := ReadVirtualField[bool](ctx, "inAlarm", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'inAlarm' field"))
	}
	_ = inAlarm

	fault, err := ReadVirtualField[bool](ctx, "fault", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fault' field"))
	}
	_ = fault

	overridden, err := ReadVirtualField[bool](ctx, "overridden", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'overridden' field"))
	}
	_ = overridden

	outOfService, err := ReadVirtualField[bool](ctx, "outOfService", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (3))), func() any { return bool(payload.GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'outOfService' field"))
	}
	_ = outOfService

	if closeErr := readBuffer.CloseContext("BACnetStatusFlagsTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetStatusFlagsTagged")
	}

	return m, nil
}

func (m *_BACnetStatusFlagsTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetStatusFlagsTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetStatusFlagsTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetStatusFlagsTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	inAlarm := m.GetInAlarm()
	_ = inAlarm
	if _inAlarmErr := writeBuffer.WriteVirtual(ctx, "inAlarm", m.GetInAlarm()); _inAlarmErr != nil {
		return errors.Wrap(_inAlarmErr, "Error serializing 'inAlarm' field")
	}
	// Virtual field
	fault := m.GetFault()
	_ = fault
	if _faultErr := writeBuffer.WriteVirtual(ctx, "fault", m.GetFault()); _faultErr != nil {
		return errors.Wrap(_faultErr, "Error serializing 'fault' field")
	}
	// Virtual field
	overridden := m.GetOverridden()
	_ = overridden
	if _overriddenErr := writeBuffer.WriteVirtual(ctx, "overridden", m.GetOverridden()); _overriddenErr != nil {
		return errors.Wrap(_overriddenErr, "Error serializing 'overridden' field")
	}
	// Virtual field
	outOfService := m.GetOutOfService()
	_ = outOfService
	if _outOfServiceErr := writeBuffer.WriteVirtual(ctx, "outOfService", m.GetOutOfService()); _outOfServiceErr != nil {
		return errors.Wrap(_outOfServiceErr, "Error serializing 'outOfService' field")
	}

	if popErr := writeBuffer.PopContext("BACnetStatusFlagsTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetStatusFlagsTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetStatusFlagsTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetStatusFlagsTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetStatusFlagsTagged) isBACnetStatusFlagsTagged() bool {
	return true
}

func (m *_BACnetStatusFlagsTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
