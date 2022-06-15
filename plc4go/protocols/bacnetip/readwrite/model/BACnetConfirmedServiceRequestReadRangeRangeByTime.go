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

// BACnetConfirmedServiceRequestReadRangeRangeByTime is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeByTime
type BACnetConfirmedServiceRequestReadRangeRangeByTime interface {
	BACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceTime returns ReferenceTime (property field)
	GetReferenceTime() BACnetDateTime
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagSignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConfirmedServiceRequestReadRangeRangeByTime is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRangeByTime struct {
	*_BACnetConfirmedServiceRequestReadRangeRange
	ReferenceTime BACnetDateTime
	Count         BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) InitializeParent(parent BACnetConfirmedServiceRequestReadRangeRange, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) {
	m.PeekedTagHeader = peekedTagHeader
	m.OpeningTag = openingTag
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetParent() BACnetConfirmedServiceRequestReadRangeRange {
	return m._BACnetConfirmedServiceRequestReadRangeRange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetReferenceTime() BACnetDateTime {
	return m.ReferenceTime
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetCount() BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRangeByTime factory function for _BACnetConfirmedServiceRequestReadRangeRangeByTime
func NewBACnetConfirmedServiceRequestReadRangeRangeByTime(referenceTime BACnetDateTime, count BACnetApplicationTagSignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestReadRangeRangeByTime {
	_result := &_BACnetConfirmedServiceRequestReadRangeRangeByTime{
		ReferenceTime: referenceTime,
		Count:         count,
		_BACnetConfirmedServiceRequestReadRangeRange: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
	}
	_result._BACnetConfirmedServiceRequestReadRangeRange._BACnetConfirmedServiceRequestReadRangeRangeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRangeByTime(structType interface{}) BACnetConfirmedServiceRequestReadRangeRangeByTime {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeByTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeByTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeByTime"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referenceTime)
	lengthInBits += m.ReferenceTime.GetLengthInBits()

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeByTimeParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReadRangeRangeByTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeByTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceTime)
	if pullErr := readBuffer.PullContext("referenceTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTime")
	}
	_referenceTime, _referenceTimeErr := BACnetDateTimeParse(readBuffer)
	if _referenceTimeErr != nil {
		return nil, errors.Wrap(_referenceTimeErr, "Error parsing 'referenceTime' field")
	}
	referenceTime := _referenceTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("referenceTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTime")
	}

	// Simple Field (count)
	if pullErr := readBuffer.PullContext("count"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for count")
	}
	_count, _countErr := BACnetApplicationTagParse(readBuffer)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := _count.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("count"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for count")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeByTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadRangeRangeByTime{
		ReferenceTime: referenceTime,
		Count:         count,
		_BACnetConfirmedServiceRequestReadRangeRange: &_BACnetConfirmedServiceRequestReadRangeRange{},
	}
	_child._BACnetConfirmedServiceRequestReadRangeRange._BACnetConfirmedServiceRequestReadRangeRangeChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeByTime")
		}

		// Simple Field (referenceTime)
		if pushErr := writeBuffer.PushContext("referenceTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceTime")
		}
		_referenceTimeErr := writeBuffer.WriteSerializable(m.GetReferenceTime())
		if popErr := writeBuffer.PopContext("referenceTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceTime")
		}
		if _referenceTimeErr != nil {
			return errors.Wrap(_referenceTimeErr, "Error serializing 'referenceTime' field")
		}

		// Simple Field (count)
		if pushErr := writeBuffer.PushContext("count"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for count")
		}
		_countErr := writeBuffer.WriteSerializable(m.GetCount())
		if popErr := writeBuffer.PopContext("count"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for count")
		}
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeByTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
