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

// BACnetProcessIdSelection is the corresponding interface of BACnetProcessIdSelection
type BACnetProcessIdSelection interface {
	BACnetProcessIdSelectionContract
	BACnetProcessIdSelectionRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetProcessIdSelectionContract provides a set of functions which can be overwritten by a sub struct
type BACnetProcessIdSelectionContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetProcessIdSelectionRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetProcessIdSelectionRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetProcessIdSelectionExactly can be used when we want exactly this type and not a type which fulfills BACnetProcessIdSelection.
// This is useful for switch cases.
type BACnetProcessIdSelectionExactly interface {
	BACnetProcessIdSelection
	isBACnetProcessIdSelection() bool
}

// _BACnetProcessIdSelection is the data-structure of this message
type _BACnetProcessIdSelection struct {
	_SubType        BACnetProcessIdSelection
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetProcessIdSelectionContract = (*_BACnetProcessIdSelection)(nil)

type BACnetProcessIdSelectionChild interface {
	utils.Serializable

	GetParent() *BACnetProcessIdSelection

	GetTypeName() string
	BACnetProcessIdSelection
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelection) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetProcessIdSelection) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProcessIdSelection factory function for _BACnetProcessIdSelection
func NewBACnetProcessIdSelection(peekedTagHeader BACnetTagHeader) *_BACnetProcessIdSelection {
	return &_BACnetProcessIdSelection{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelection(structType any) BACnetProcessIdSelection {
	if casted, ok := structType.(BACnetProcessIdSelection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelection) GetTypeName() string {
	return "BACnetProcessIdSelection"
}

func (m *_BACnetProcessIdSelection) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetProcessIdSelection) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetProcessIdSelectionParse[T BACnetProcessIdSelection](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetProcessIdSelectionParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetProcessIdSelectionParseWithBufferProducer[T BACnetProcessIdSelection]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetProcessIdSelectionParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetProcessIdSelectionParseWithBuffer[T BACnetProcessIdSelection](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetProcessIdSelection{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetProcessIdSelection) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetProcessIdSelection BACnetProcessIdSelection, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetProcessIdSelection
	switch {
	case peekedTagNumber == uint8(0): // BACnetProcessIdSelectionNull
		if _child, err = (&_BACnetProcessIdSelectionNull{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetProcessIdSelectionNull for type-switch of BACnetProcessIdSelection")
		}
	case 0 == 0: // BACnetProcessIdSelectionValue
		if _child, err = (&_BACnetProcessIdSelectionValue{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetProcessIdSelectionValue for type-switch of BACnetProcessIdSelection")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelection")
	}

	return _child, nil
}

func (pm *_BACnetProcessIdSelection) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetProcessIdSelection, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProcessIdSelection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelection")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetProcessIdSelection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelection")
	}
	return nil
}

func (m *_BACnetProcessIdSelection) isBACnetProcessIdSelection() bool {
	return true
}
