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

// BACnetTagPayloadTime is the corresponding interface of BACnetTagPayloadTime
type BACnetTagPayloadTime interface {
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
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetTagPayloadTime is the data-structure of this message
type _BACnetTagPayloadTime struct {
	Hour       uint8
	Minute     uint8
	Second     uint8
	Fractional uint8
}

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
	return uint8(0xFF)
}

func (m *_BACnetTagPayloadTime) GetHourIsWildcard() bool {
	return bool(bool((m.GetHour()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadTime) GetMinuteIsWildcard() bool {
	return bool(bool((m.GetMinute()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadTime) GetSecondIsWildcard() bool {
	return bool(bool((m.GetSecond()) == (m.GetWildcard())))
}

func (m *_BACnetTagPayloadTime) GetFractionalIsWildcard() bool {
	return bool(bool((m.GetFractional()) == (m.GetWildcard())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadTime factory function for _BACnetTagPayloadTime
func NewBACnetTagPayloadTime(hour uint8, minute uint8, second uint8, fractional uint8) *_BACnetTagPayloadTime {
	return &_BACnetTagPayloadTime{Hour: hour, Minute: minute, Second: second, Fractional: fractional}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadTime(structType interface{}) BACnetTagPayloadTime {
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

func (m *_BACnetTagPayloadTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadTime) GetLengthInBitsConditional(lastItem bool) uint16 {
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

func (m *_BACnetTagPayloadTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadTimeParse(readBuffer utils.ReadBuffer) (BACnetTagPayloadTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_wildcard := 0xFF
	wildcard := uint8(_wildcard)
	_ = wildcard

	// Simple Field (hour)
	_hour, _hourErr := readBuffer.ReadUint8("hour", 8)
	if _hourErr != nil {
		return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
	}
	hour := _hour

	// Virtual field
	_hourIsWildcard := bool((hour) == (wildcard))
	hourIsWildcard := bool(_hourIsWildcard)
	_ = hourIsWildcard

	// Simple Field (minute)
	_minute, _minuteErr := readBuffer.ReadUint8("minute", 8)
	if _minuteErr != nil {
		return nil, errors.Wrap(_minuteErr, "Error parsing 'minute' field")
	}
	minute := _minute

	// Virtual field
	_minuteIsWildcard := bool((minute) == (wildcard))
	minuteIsWildcard := bool(_minuteIsWildcard)
	_ = minuteIsWildcard

	// Simple Field (second)
	_second, _secondErr := readBuffer.ReadUint8("second", 8)
	if _secondErr != nil {
		return nil, errors.Wrap(_secondErr, "Error parsing 'second' field")
	}
	second := _second

	// Virtual field
	_secondIsWildcard := bool((second) == (wildcard))
	secondIsWildcard := bool(_secondIsWildcard)
	_ = secondIsWildcard

	// Simple Field (fractional)
	_fractional, _fractionalErr := readBuffer.ReadUint8("fractional", 8)
	if _fractionalErr != nil {
		return nil, errors.Wrap(_fractionalErr, "Error parsing 'fractional' field")
	}
	fractional := _fractional

	// Virtual field
	_fractionalIsWildcard := bool((fractional) == (wildcard))
	fractionalIsWildcard := bool(_fractionalIsWildcard)
	_ = fractionalIsWildcard

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadTime")
	}

	// Create the instance
	return NewBACnetTagPayloadTime(hour, minute, second, fractional), nil
}

func (m *_BACnetTagPayloadTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadTime")
	}
	// Virtual field
	if _wildcardErr := writeBuffer.WriteVirtual("wildcard", m.GetWildcard()); _wildcardErr != nil {
		return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
	}

	// Simple Field (hour)
	hour := uint8(m.GetHour())
	_hourErr := writeBuffer.WriteUint8("hour", 8, (hour))
	if _hourErr != nil {
		return errors.Wrap(_hourErr, "Error serializing 'hour' field")
	}
	// Virtual field
	if _hourIsWildcardErr := writeBuffer.WriteVirtual("hourIsWildcard", m.GetHourIsWildcard()); _hourIsWildcardErr != nil {
		return errors.Wrap(_hourIsWildcardErr, "Error serializing 'hourIsWildcard' field")
	}

	// Simple Field (minute)
	minute := uint8(m.GetMinute())
	_minuteErr := writeBuffer.WriteUint8("minute", 8, (minute))
	if _minuteErr != nil {
		return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
	}
	// Virtual field
	if _minuteIsWildcardErr := writeBuffer.WriteVirtual("minuteIsWildcard", m.GetMinuteIsWildcard()); _minuteIsWildcardErr != nil {
		return errors.Wrap(_minuteIsWildcardErr, "Error serializing 'minuteIsWildcard' field")
	}

	// Simple Field (second)
	second := uint8(m.GetSecond())
	_secondErr := writeBuffer.WriteUint8("second", 8, (second))
	if _secondErr != nil {
		return errors.Wrap(_secondErr, "Error serializing 'second' field")
	}
	// Virtual field
	if _secondIsWildcardErr := writeBuffer.WriteVirtual("secondIsWildcard", m.GetSecondIsWildcard()); _secondIsWildcardErr != nil {
		return errors.Wrap(_secondIsWildcardErr, "Error serializing 'secondIsWildcard' field")
	}

	// Simple Field (fractional)
	fractional := uint8(m.GetFractional())
	_fractionalErr := writeBuffer.WriteUint8("fractional", 8, (fractional))
	if _fractionalErr != nil {
		return errors.Wrap(_fractionalErr, "Error serializing 'fractional' field")
	}
	// Virtual field
	if _fractionalIsWildcardErr := writeBuffer.WriteVirtual("fractionalIsWildcard", m.GetFractionalIsWildcard()); _fractionalIsWildcardErr != nil {
		return errors.Wrap(_fractionalIsWildcardErr, "Error serializing 'fractionalIsWildcard' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadTime")
	}
	return nil
}

func (m *_BACnetTagPayloadTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
