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

// BACnetServiceAck is the corresponding interface of BACnetServiceAck
type BACnetServiceAck interface {
	BACnetServiceAckContract
	BACnetServiceAckRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetServiceAckContract provides a set of functions which can be overwritten by a sub struct
type BACnetServiceAckContract interface {
	// GetServiceAckPayloadLength returns ServiceAckPayloadLength (virtual field)
	GetServiceAckPayloadLength() uint32
	// GetServiceAckLength() returns a parser argument
	GetServiceAckLength() uint32
}

// BACnetServiceAckRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetServiceAckRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetConfirmedServiceChoice
}

// BACnetServiceAckExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAck.
// This is useful for switch cases.
type BACnetServiceAckExactly interface {
	BACnetServiceAck
	isBACnetServiceAck() bool
}

// _BACnetServiceAck is the data-structure of this message
type _BACnetServiceAck struct {
	_SubType BACnetServiceAck

	// Arguments.
	ServiceAckLength uint32
}

var _ BACnetServiceAckContract = (*_BACnetServiceAck)(nil)

type BACnetServiceAckChild interface {
	utils.Serializable

	GetParent() *BACnetServiceAck

	GetTypeName() string
	BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetServiceAck) GetServiceAckPayloadLength() uint32 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint32(utils.InlineIf((bool((m.GetServiceAckLength()) > (0))), func() any { return uint32((uint32(m.GetServiceAckLength()) - uint32(uint32(1)))) }, func() any { return uint32(uint32(0)) }).(uint32))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAck factory function for _BACnetServiceAck
func NewBACnetServiceAck(serviceAckLength uint32) *_BACnetServiceAck {
	return &_BACnetServiceAck{ServiceAckLength: serviceAckLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAck(structType any) BACnetServiceAck {
	if casted, ok := structType.(BACnetServiceAck); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAck) GetTypeName() string {
	return "BACnetServiceAck"
}

func (m *_BACnetServiceAck) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetServiceAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckParse[T BACnetServiceAck](ctx context.Context, theBytes []byte, serviceAckLength uint32) (T, error) {
	return BACnetServiceAckParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckParseWithBufferProducer[T BACnetServiceAck](serviceAckLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetServiceAckParseWithBuffer[T](ctx, readBuffer, serviceAckLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetServiceAckParseWithBuffer[T BACnetServiceAck](ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (T, error) {
	v, err := (&_BACnetServiceAck{ServiceAckLength: serviceAckLength}).parse(ctx, readBuffer, serviceAckLength)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetServiceAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (__bACnetServiceAck BACnetServiceAck, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceChoice, err := ReadDiscriminatorEnumField[BACnetConfirmedServiceChoice](ctx, "serviceChoice", "BACnetConfirmedServiceChoice", ReadEnum(BACnetConfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceChoice' field"))
	}

	serviceAckPayloadLength, err := ReadVirtualField[uint32](ctx, "serviceAckPayloadLength", (*uint32)(nil), utils.InlineIf((bool((serviceAckLength) > (0))), func() any { return uint32((uint32(serviceAckLength) - uint32(uint32(1)))) }, func() any { return uint32(uint32(0)) }).(uint32))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceAckPayloadLength' field"))
	}
	_ = serviceAckPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetServiceAck
	switch {
	case serviceChoice == BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY: // BACnetServiceAckGetAlarmSummary
		if _child, err = (&_BACnetServiceAckGetAlarmSummary{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckGetAlarmSummary for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY: // BACnetServiceAckGetEnrollmentSummary
		if _child, err = (&_BACnetServiceAckGetEnrollmentSummary{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckGetEnrollmentSummary for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION: // BACnetServiceAckGetEventInformation
		if _child, err = (&_BACnetServiceAckGetEventInformation{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckGetEventInformation for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_READ_FILE: // BACnetServiceAckAtomicReadFile
		if _child, err = (&_BACnetServiceAckAtomicReadFile{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAtomicReadFile for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE: // BACnetServiceAckAtomicWriteFile
		if _child, err = (&_BACnetServiceAckAtomicWriteFile{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAtomicWriteFile for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT: // BACnetServiceAckCreateObject
		if _child, err = (&_BACnetServiceAckCreateObject{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckCreateObject for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY: // BACnetServiceAckReadProperty
		if _child, err = (&_BACnetServiceAckReadProperty{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckReadProperty for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE: // BACnetServiceAckReadPropertyMultiple
		if _child, err = (&_BACnetServiceAckReadPropertyMultiple{}).parse(ctx, readBuffer, m, serviceAckPayloadLength, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckReadPropertyMultiple for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_RANGE: // BACnetServiceAckReadRange
		if _child, err = (&_BACnetServiceAckReadRange{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckReadRange for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER: // BACnetServiceAckConfirmedPrivateTransfer
		if _child, err = (&_BACnetServiceAckConfirmedPrivateTransfer{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckConfirmedPrivateTransfer for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_VT_OPEN: // BACnetServiceAckVTOpen
		if _child, err = (&_BACnetServiceAckVTOpen{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckVTOpen for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_VT_DATA: // BACnetServiceAckVTData
		if _child, err = (&_BACnetServiceAckVTData{}).parse(ctx, readBuffer, m, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckVTData for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_AUTHENTICATE: // BACnetServiceAckAuthenticate
		if _child, err = (&_BACnetServiceAckAuthenticate{}).parse(ctx, readBuffer, m, serviceAckPayloadLength, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckAuthenticate for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_REQUEST_KEY: // BACnetServiceAckRequestKey
		if _child, err = (&_BACnetServiceAckRequestKey{}).parse(ctx, readBuffer, m, serviceAckPayloadLength, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckRequestKey for type-switch of BACnetServiceAck")
		}
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL: // BACnetServiceAckReadPropertyConditional
		if _child, err = (&_BACnetServiceAckReadPropertyConditional{}).parse(ctx, readBuffer, m, serviceAckPayloadLength, serviceAckLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetServiceAckReadPropertyConditional for type-switch of BACnetServiceAck")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [serviceChoice=%v]", serviceChoice)
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAck")
	}

	return _child, nil
}

func (pm *_BACnetServiceAck) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetServiceAck, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetServiceAck"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetServiceAck")
	}

	if err := WriteDiscriminatorEnumField(ctx, "serviceChoice", "BACnetConfirmedServiceChoice", m.GetServiceChoice(), WriteEnum[BACnetConfirmedServiceChoice, uint8](BACnetConfirmedServiceChoice.GetValue, BACnetConfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'serviceChoice' field")
	}
	// Virtual field
	serviceAckPayloadLength := m.GetServiceAckPayloadLength()
	_ = serviceAckPayloadLength
	if _serviceAckPayloadLengthErr := writeBuffer.WriteVirtual(ctx, "serviceAckPayloadLength", m.GetServiceAckPayloadLength()); _serviceAckPayloadLengthErr != nil {
		return errors.Wrap(_serviceAckPayloadLengthErr, "Error serializing 'serviceAckPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetServiceAck"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetServiceAck")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetServiceAck) GetServiceAckLength() uint32 {
	return m.ServiceAckLength
}

//
////

func (m *_BACnetServiceAck) isBACnetServiceAck() bool {
	return true
}
