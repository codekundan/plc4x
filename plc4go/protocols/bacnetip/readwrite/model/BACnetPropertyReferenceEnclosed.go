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

// BACnetPropertyReferenceEnclosed is the corresponding interface of BACnetPropertyReferenceEnclosed
type BACnetPropertyReferenceEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetReference returns Reference (property field)
	GetReference() BACnetPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetPropertyReferenceEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyReferenceEnclosed.
// This is useful for switch cases.
type BACnetPropertyReferenceEnclosedExactly interface {
	BACnetPropertyReferenceEnclosed
	isBACnetPropertyReferenceEnclosed() bool
}

// _BACnetPropertyReferenceEnclosed is the data-structure of this message
type _BACnetPropertyReferenceEnclosed struct {
	OpeningTag BACnetOpeningTag
	Reference  BACnetPropertyReference
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetPropertyReferenceEnclosed = (*_BACnetPropertyReferenceEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetPropertyReferenceEnclosed) GetReference() BACnetPropertyReference {
	return m.Reference
}

func (m *_BACnetPropertyReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyReferenceEnclosed factory function for _BACnetPropertyReferenceEnclosed
func NewBACnetPropertyReferenceEnclosed(openingTag BACnetOpeningTag, reference BACnetPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetPropertyReferenceEnclosed {
	return &_BACnetPropertyReferenceEnclosed{OpeningTag: openingTag, Reference: reference, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyReferenceEnclosed(structType any) BACnetPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetPropertyReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetPropertyReferenceEnclosed"
}

func (m *_BACnetPropertyReferenceEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (reference)
	lengthInBits += m.Reference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyReferenceEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyReferenceEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetPropertyReferenceEnclosed, error) {
	return BACnetPropertyReferenceEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetPropertyReferenceEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReferenceEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReferenceEnclosed, error) {
		return BACnetPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetPropertyReferenceEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetPropertyReferenceEnclosed, error) {
	v, err := (&_BACnetPropertyReferenceEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetPropertyReferenceEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetPropertyReferenceEnclosed BACnetPropertyReferenceEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	reference, err := ReadSimpleField[BACnetPropertyReference](ctx, "reference", ReadComplex[BACnetPropertyReference](BACnetPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reference' field"))
	}
	m.Reference = reference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyReferenceEnclosed")
	}

	return m, nil
}

func (m *_BACnetPropertyReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyReferenceEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyReferenceEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetPropertyReference](ctx, "reference", m.GetReference(), WriteComplex[BACnetPropertyReference](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reference' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetPropertyReferenceEnclosed) isBACnetPropertyReferenceEnclosed() bool {
	return true
}

func (m *_BACnetPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
