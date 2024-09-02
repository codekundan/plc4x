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

// LDataFrame is the corresponding interface of LDataFrame
type LDataFrame interface {
	LDataFrameContract
	LDataFrameRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// LDataFrameContract provides a set of functions which can be overwritten by a sub struct
type LDataFrameContract interface {
	// GetFrameType returns FrameType (property field)
	GetFrameType() bool
	// GetNotRepeated returns NotRepeated (property field)
	GetNotRepeated() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetAcknowledgeRequested returns AcknowledgeRequested (property field)
	GetAcknowledgeRequested() bool
	// GetErrorFlag returns ErrorFlag (property field)
	GetErrorFlag() bool
}

// LDataFrameRequirements provides a set of functions which need to be implemented by a sub struct
type LDataFrameRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetNotAckFrame returns NotAckFrame (discriminator field)
	GetNotAckFrame() bool
	// GetPolling returns Polling (discriminator field)
	GetPolling() bool
}

// LDataFrameExactly can be used when we want exactly this type and not a type which fulfills LDataFrame.
// This is useful for switch cases.
type LDataFrameExactly interface {
	LDataFrame
	isLDataFrame() bool
}

// _LDataFrame is the data-structure of this message
type _LDataFrame struct {
	_SubType             LDataFrame
	FrameType            bool
	NotRepeated          bool
	Priority             CEMIPriority
	AcknowledgeRequested bool
	ErrorFlag            bool
}

var _ LDataFrameContract = (*_LDataFrame)(nil)

type LDataFrameChild interface {
	utils.Serializable

	GetParent() *LDataFrame

	GetTypeName() string
	LDataFrame
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataFrame) GetFrameType() bool {
	return m.FrameType
}

func (m *_LDataFrame) GetNotRepeated() bool {
	return m.NotRepeated
}

func (m *_LDataFrame) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_LDataFrame) GetAcknowledgeRequested() bool {
	return m.AcknowledgeRequested
}

func (m *_LDataFrame) GetErrorFlag() bool {
	return m.ErrorFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLDataFrame factory function for _LDataFrame
func NewLDataFrame(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *_LDataFrame {
	return &_LDataFrame{FrameType: frameType, NotRepeated: notRepeated, Priority: priority, AcknowledgeRequested: acknowledgeRequested, ErrorFlag: errorFlag}
}

// Deprecated: use the interface for direct cast
func CastLDataFrame(structType any) LDataFrame {
	if casted, ok := structType.(LDataFrame); ok {
		return casted
	}
	if casted, ok := structType.(*LDataFrame); ok {
		return *casted
	}
	return nil
}

func (m *_LDataFrame) GetTypeName() string {
	return "LDataFrame"
}

func (m *_LDataFrame) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (frameType)
	lengthInBits += 1
	// Discriminator Field (polling)
	lengthInBits += 1

	// Simple field (notRepeated)
	lengthInBits += 1
	// Discriminator Field (notAckFrame)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (acknowledgeRequested)
	lengthInBits += 1

	// Simple field (errorFlag)
	lengthInBits += 1

	return lengthInBits
}

func (m *_LDataFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func LDataFrameParse[T LDataFrame](ctx context.Context, theBytes []byte) (T, error) {
	return LDataFrameParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func LDataFrameParseWithBufferProducer[T LDataFrame]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := LDataFrameParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func LDataFrameParseWithBuffer[T LDataFrame](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_LDataFrame{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_LDataFrame) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__lDataFrame LDataFrame, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	frameType, err := ReadSimpleField(ctx, "frameType", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'frameType' field"))
	}
	m.FrameType = frameType

	polling, err := ReadDiscriminatorField[bool](ctx, "polling", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'polling' field"))
	}

	notRepeated, err := ReadSimpleField(ctx, "notRepeated", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notRepeated' field"))
	}
	m.NotRepeated = notRepeated

	notAckFrame, err := ReadDiscriminatorField[bool](ctx, "notAckFrame", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notAckFrame' field"))
	}

	priority, err := ReadEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", ReadEnum(CEMIPriorityByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	acknowledgeRequested, err := ReadSimpleField(ctx, "acknowledgeRequested", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgeRequested' field"))
	}
	m.AcknowledgeRequested = acknowledgeRequested

	errorFlag, err := ReadSimpleField(ctx, "errorFlag", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorFlag' field"))
	}
	m.ErrorFlag = errorFlag

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child LDataFrame
	switch {
	case notAckFrame == bool(true) && polling == bool(false): // LDataExtended
		if _child, err = (&_LDataExtended{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataExtended for type-switch of LDataFrame")
		}
	case notAckFrame == bool(true) && polling == bool(true): // LPollData
		if _child, err = (&_LPollData{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LPollData for type-switch of LDataFrame")
		}
	case notAckFrame == bool(false): // LDataFrameACK
		if _child, err = (&_LDataFrameACK{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type LDataFrameACK for type-switch of LDataFrame")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [notAckFrame=%v, polling=%v]", notAckFrame, polling)
	}

	if closeErr := readBuffer.CloseContext("LDataFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataFrame")
	}

	return _child, nil
}

func (pm *_LDataFrame) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LDataFrame, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LDataFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LDataFrame")
	}

	if err := WriteSimpleField[bool](ctx, "frameType", m.GetFrameType(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'frameType' field")
	}

	if err := WriteDiscriminatorField(ctx, "polling", m.GetPolling(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'polling' field")
	}

	if err := WriteSimpleField[bool](ctx, "notRepeated", m.GetNotRepeated(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'notRepeated' field")
	}

	if err := WriteDiscriminatorField(ctx, "notAckFrame", m.GetNotAckFrame(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'notAckFrame' field")
	}

	if err := WriteSimpleEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", m.GetPriority(), WriteEnum[CEMIPriority, uint8](CEMIPriority.GetValue, CEMIPriority.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if err := WriteSimpleField[bool](ctx, "acknowledgeRequested", m.GetAcknowledgeRequested(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'acknowledgeRequested' field")
	}

	if err := WriteSimpleField[bool](ctx, "errorFlag", m.GetErrorFlag(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorFlag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LDataFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LDataFrame")
	}
	return nil
}

func (m *_LDataFrame) isLDataFrame() bool {
	return true
}
