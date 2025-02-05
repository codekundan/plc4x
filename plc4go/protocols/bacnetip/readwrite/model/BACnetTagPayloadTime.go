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

// BACnetTagPayloadTime is the corresponding interface of BACnetTagPayloadTime
type BACnetTagPayloadTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHour returns Hour (property field)
	GetHour() uint8
	// GetMinute returns Minute (property field)
	GetMinute() uint8
	// GetSecond returns Second (property field)
	GetSecond() uint8
	// GetFractional returns Fractional (property field)
	GetFractional() uint8
	// GetWildcard returns Wildcard (virtual field)
	GetWildcard() uint8
	// GetHourIsWildcard returns HourIsWildcard (virtual field)
	GetHourIsWildcard() bool
	// GetMinuteIsWildcard returns MinuteIsWildcard (virtual field)
	GetMinuteIsWildcard() bool
	// GetSecondIsWildcard returns SecondIsWildcard (virtual field)
	GetSecondIsWildcard() bool
	// GetFractionalIsWildcard returns FractionalIsWildcard (virtual field)
	GetFractionalIsWildcard() bool
	// IsBACnetTagPayloadTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadTime()
	// CreateBuilder creates a BACnetTagPayloadTimeBuilder
	CreateBACnetTagPayloadTimeBuilder() BACnetTagPayloadTimeBuilder
}

// _BACnetTagPayloadTime is the data-structure of this message
type _BACnetTagPayloadTime struct {
	Hour       uint8
	Minute     uint8
	Second     uint8
	Fractional uint8
}

var _ BACnetTagPayloadTime = (*_BACnetTagPayloadTime)(nil)

// NewBACnetTagPayloadTime factory function for _BACnetTagPayloadTime
func NewBACnetTagPayloadTime(hour uint8, minute uint8, second uint8, fractional uint8) *_BACnetTagPayloadTime {
	return &_BACnetTagPayloadTime{Hour: hour, Minute: minute, Second: second, Fractional: fractional}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagPayloadTimeBuilder is a builder for BACnetTagPayloadTime
type BACnetTagPayloadTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hour uint8, minute uint8, second uint8, fractional uint8) BACnetTagPayloadTimeBuilder
	// WithHour adds Hour (property field)
	WithHour(uint8) BACnetTagPayloadTimeBuilder
	// WithMinute adds Minute (property field)
	WithMinute(uint8) BACnetTagPayloadTimeBuilder
	// WithSecond adds Second (property field)
	WithSecond(uint8) BACnetTagPayloadTimeBuilder
	// WithFractional adds Fractional (property field)
	WithFractional(uint8) BACnetTagPayloadTimeBuilder
	// Build builds the BACnetTagPayloadTime or returns an error if something is wrong
	Build() (BACnetTagPayloadTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagPayloadTime
}

// NewBACnetTagPayloadTimeBuilder() creates a BACnetTagPayloadTimeBuilder
func NewBACnetTagPayloadTimeBuilder() BACnetTagPayloadTimeBuilder {
	return &_BACnetTagPayloadTimeBuilder{_BACnetTagPayloadTime: new(_BACnetTagPayloadTime)}
}

type _BACnetTagPayloadTimeBuilder struct {
	*_BACnetTagPayloadTime

	err *utils.MultiError
}

var _ (BACnetTagPayloadTimeBuilder) = (*_BACnetTagPayloadTimeBuilder)(nil)

func (b *_BACnetTagPayloadTimeBuilder) WithMandatoryFields(hour uint8, minute uint8, second uint8, fractional uint8) BACnetTagPayloadTimeBuilder {
	return b.WithHour(hour).WithMinute(minute).WithSecond(second).WithFractional(fractional)
}

func (b *_BACnetTagPayloadTimeBuilder) WithHour(hour uint8) BACnetTagPayloadTimeBuilder {
	b.Hour = hour
	return b
}

func (b *_BACnetTagPayloadTimeBuilder) WithMinute(minute uint8) BACnetTagPayloadTimeBuilder {
	b.Minute = minute
	return b
}

func (b *_BACnetTagPayloadTimeBuilder) WithSecond(second uint8) BACnetTagPayloadTimeBuilder {
	b.Second = second
	return b
}

func (b *_BACnetTagPayloadTimeBuilder) WithFractional(fractional uint8) BACnetTagPayloadTimeBuilder {
	b.Fractional = fractional
	return b
}

func (b *_BACnetTagPayloadTimeBuilder) Build() (BACnetTagPayloadTime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTagPayloadTime.deepCopy(), nil
}

func (b *_BACnetTagPayloadTimeBuilder) MustBuild() BACnetTagPayloadTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTagPayloadTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTagPayloadTimeBuilder().(*_BACnetTagPayloadTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTagPayloadTimeBuilder creates a BACnetTagPayloadTimeBuilder
func (b *_BACnetTagPayloadTime) CreateBACnetTagPayloadTimeBuilder() BACnetTagPayloadTimeBuilder {
	if b == nil {
		return NewBACnetTagPayloadTimeBuilder()
	}
	return &_BACnetTagPayloadTimeBuilder{_BACnetTagPayloadTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadTime) GetHour() uint8 {
	return m.Hour
}

func (m *_BACnetTagPayloadTime) GetMinute() uint8 {
	return m.Minute
}

func (m *_BACnetTagPayloadTime) GetSecond() uint8 {
	return m.Second
}

func (m *_BACnetTagPayloadTime) GetFractional() uint8 {
	return m.Fractional
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadTime) GetWildcard() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(0xFF)
}

func (m *_BACnetTagPayloadTime) GetHourIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetHour()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadTime) GetMinuteIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMinute()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadTime) GetSecondIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetSecond()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadTime) GetFractionalIsWildcard() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetFractional()) == (m.GetWildcard())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadTime(structType any) BACnetTagPayloadTime {
	if casted, ok := structType.(BACnetTagPayloadTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadTime) GetTypeName() string {
	return "BACnetTagPayloadTime"
}

func (m *_BACnetTagPayloadTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (hour)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (minute)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (second)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (fractional)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadTimeParse(ctx context.Context, theBytes []byte) (BACnetTagPayloadTime, error) {
	return BACnetTagPayloadTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadTime, error) {
		return BACnetTagPayloadTimeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTagPayloadTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadTime, error) {
	v, err := (&_BACnetTagPayloadTime{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadTime) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTagPayloadTime BACnetTagPayloadTime, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	wildcard, err := ReadVirtualField[uint8](ctx, "wildcard", (*uint8)(nil), 0xFF)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'wildcard' field"))
	}
	_ = wildcard

	hour, err := ReadSimpleField(ctx, "hour", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hour' field"))
	}
	m.Hour = hour

	hourIsWildcard, err := ReadVirtualField[bool](ctx, "hourIsWildcard", (*bool)(nil), bool((hour) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hourIsWildcard' field"))
	}
	_ = hourIsWildcard

	minute, err := ReadSimpleField(ctx, "minute", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minute' field"))
	}
	m.Minute = minute

	minuteIsWildcard, err := ReadVirtualField[bool](ctx, "minuteIsWildcard", (*bool)(nil), bool((minute) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minuteIsWildcard' field"))
	}
	_ = minuteIsWildcard

	second, err := ReadSimpleField(ctx, "second", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'second' field"))
	}
	m.Second = second

	secondIsWildcard, err := ReadVirtualField[bool](ctx, "secondIsWildcard", (*bool)(nil), bool((second) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secondIsWildcard' field"))
	}
	_ = secondIsWildcard

	fractional, err := ReadSimpleField(ctx, "fractional", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fractional' field"))
	}
	m.Fractional = fractional

	fractionalIsWildcard, err := ReadVirtualField[bool](ctx, "fractionalIsWildcard", (*bool)(nil), bool((fractional) == (wildcard)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fractionalIsWildcard' field"))
	}
	_ = fractionalIsWildcard

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadTime")
	}

	return m, nil
}

func (m *_BACnetTagPayloadTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadTime")
	}
	// Virtual field
	wildcard := m.GetWildcard()
	_ = wildcard
	if _wildcardErr := writeBuffer.WriteVirtual(ctx, "wildcard", m.GetWildcard()); _wildcardErr != nil {
		return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "hour", m.GetHour(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'hour' field")
	}
	// Virtual field
	hourIsWildcard := m.GetHourIsWildcard()
	_ = hourIsWildcard
	if _hourIsWildcardErr := writeBuffer.WriteVirtual(ctx, "hourIsWildcard", m.GetHourIsWildcard()); _hourIsWildcardErr != nil {
		return errors.Wrap(_hourIsWildcardErr, "Error serializing 'hourIsWildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "minute", m.GetMinute(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'minute' field")
	}
	// Virtual field
	minuteIsWildcard := m.GetMinuteIsWildcard()
	_ = minuteIsWildcard
	if _minuteIsWildcardErr := writeBuffer.WriteVirtual(ctx, "minuteIsWildcard", m.GetMinuteIsWildcard()); _minuteIsWildcardErr != nil {
		return errors.Wrap(_minuteIsWildcardErr, "Error serializing 'minuteIsWildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "second", m.GetSecond(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'second' field")
	}
	// Virtual field
	secondIsWildcard := m.GetSecondIsWildcard()
	_ = secondIsWildcard
	if _secondIsWildcardErr := writeBuffer.WriteVirtual(ctx, "secondIsWildcard", m.GetSecondIsWildcard()); _secondIsWildcardErr != nil {
		return errors.Wrap(_secondIsWildcardErr, "Error serializing 'secondIsWildcard' field")
	}

	if err := WriteSimpleField[uint8](ctx, "fractional", m.GetFractional(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'fractional' field")
	}
	// Virtual field
	fractionalIsWildcard := m.GetFractionalIsWildcard()
	_ = fractionalIsWildcard
	if _fractionalIsWildcardErr := writeBuffer.WriteVirtual(ctx, "fractionalIsWildcard", m.GetFractionalIsWildcard()); _fractionalIsWildcardErr != nil {
		return errors.Wrap(_fractionalIsWildcardErr, "Error serializing 'fractionalIsWildcard' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadTime")
	}
	return nil
}

func (m *_BACnetTagPayloadTime) IsBACnetTagPayloadTime() {}

func (m *_BACnetTagPayloadTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadTime) deepCopy() *_BACnetTagPayloadTime {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadTimeCopy := &_BACnetTagPayloadTime{
		m.Hour,
		m.Minute,
		m.Second,
		m.Fractional,
	}
	return _BACnetTagPayloadTimeCopy
}

func (m *_BACnetTagPayloadTime) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
