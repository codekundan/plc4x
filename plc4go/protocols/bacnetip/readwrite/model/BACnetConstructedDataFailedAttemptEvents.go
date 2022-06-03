/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataFailedAttemptEvents is the data-structure of this message
type BACnetConstructedDataFailedAttemptEvents struct {
	*BACnetConstructedData
	FailedAttemptEvents []*BACnetAccessEventTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataFailedAttemptEvents is the corresponding interface of BACnetConstructedDataFailedAttemptEvents
type IBACnetConstructedDataFailedAttemptEvents interface {
	IBACnetConstructedData
	// GetFailedAttemptEvents returns FailedAttemptEvents (property field)
	GetFailedAttemptEvents() []*BACnetAccessEventTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataFailedAttemptEvents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataFailedAttemptEvents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAILED_ATTEMPT_EVENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFailedAttemptEvents) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFailedAttemptEvents) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFailedAttemptEvents) GetFailedAttemptEvents() []*BACnetAccessEventTagged {
	return m.FailedAttemptEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFailedAttemptEvents factory function for BACnetConstructedDataFailedAttemptEvents
func NewBACnetConstructedDataFailedAttemptEvents(failedAttemptEvents []*BACnetAccessEventTagged, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataFailedAttemptEvents {
	_result := &BACnetConstructedDataFailedAttemptEvents{
		FailedAttemptEvents:   failedAttemptEvents,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFailedAttemptEvents(structType interface{}) *BACnetConstructedDataFailedAttemptEvents {
	if casted, ok := structType.(BACnetConstructedDataFailedAttemptEvents); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFailedAttemptEvents); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFailedAttemptEvents(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFailedAttemptEvents(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFailedAttemptEvents) GetTypeName() string {
	return "BACnetConstructedDataFailedAttemptEvents"
}

func (m *BACnetConstructedDataFailedAttemptEvents) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFailedAttemptEvents) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.FailedAttemptEvents) > 0 {
		for _, element := range m.FailedAttemptEvents {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataFailedAttemptEvents) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFailedAttemptEventsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataFailedAttemptEvents, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFailedAttemptEvents"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (failedAttemptEvents)
	if pullErr := readBuffer.PullContext("failedAttemptEvents", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	failedAttemptEvents := make([]*BACnetAccessEventTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAccessEventTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'failedAttemptEvents' field")
			}
			failedAttemptEvents = append(failedAttemptEvents, CastBACnetAccessEventTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("failedAttemptEvents", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFailedAttemptEvents"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFailedAttemptEvents{
		FailedAttemptEvents:   failedAttemptEvents,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFailedAttemptEvents) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFailedAttemptEvents"); pushErr != nil {
			return pushErr
		}

		// Array Field (failedAttemptEvents)
		if m.FailedAttemptEvents != nil {
			if pushErr := writeBuffer.PushContext("failedAttemptEvents", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.FailedAttemptEvents {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'failedAttemptEvents' field")
				}
			}
			if popErr := writeBuffer.PopContext("failedAttemptEvents", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFailedAttemptEvents"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFailedAttemptEvents) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
