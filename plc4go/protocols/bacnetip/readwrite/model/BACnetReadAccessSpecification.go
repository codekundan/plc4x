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

// BACnetReadAccessSpecification is the corresponding interface of BACnetReadAccessSpecification
type BACnetReadAccessSpecification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfPropertyReferences returns ListOfPropertyReferences (property field)
	GetListOfPropertyReferences() []BACnetPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetReadAccessSpecificationExactly can be used when we want exactly this type and not a type which fulfills BACnetReadAccessSpecification.
// This is useful for switch cases.
type BACnetReadAccessSpecificationExactly interface {
	BACnetReadAccessSpecification
	isBACnetReadAccessSpecification() bool
}

// _BACnetReadAccessSpecification is the data-structure of this message
type _BACnetReadAccessSpecification struct {
	ObjectIdentifier         BACnetContextTagObjectIdentifier
	OpeningTag               BACnetOpeningTag
	ListOfPropertyReferences []BACnetPropertyReference
	ClosingTag               BACnetClosingTag
}

var _ BACnetReadAccessSpecification = (*_BACnetReadAccessSpecification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetReadAccessSpecification) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetReadAccessSpecification) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetReadAccessSpecification) GetListOfPropertyReferences() []BACnetPropertyReference {
	return m.ListOfPropertyReferences
}

func (m *_BACnetReadAccessSpecification) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetReadAccessSpecification factory function for _BACnetReadAccessSpecification
func NewBACnetReadAccessSpecification(objectIdentifier BACnetContextTagObjectIdentifier, openingTag BACnetOpeningTag, listOfPropertyReferences []BACnetPropertyReference, closingTag BACnetClosingTag) *_BACnetReadAccessSpecification {
	return &_BACnetReadAccessSpecification{ObjectIdentifier: objectIdentifier, OpeningTag: openingTag, ListOfPropertyReferences: listOfPropertyReferences, ClosingTag: closingTag}
}

// Deprecated: use the interface for direct cast
func CastBACnetReadAccessSpecification(structType any) BACnetReadAccessSpecification {
	if casted, ok := structType.(BACnetReadAccessSpecification); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetReadAccessSpecification); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetReadAccessSpecification) GetTypeName() string {
	return "BACnetReadAccessSpecification"
}

func (m *_BACnetReadAccessSpecification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfPropertyReferences) > 0 {
		for _, element := range m.ListOfPropertyReferences {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetReadAccessSpecification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetReadAccessSpecificationParse(ctx context.Context, theBytes []byte) (BACnetReadAccessSpecification, error) {
	return BACnetReadAccessSpecificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetReadAccessSpecificationParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessSpecification, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessSpecification, error) {
		return BACnetReadAccessSpecificationParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetReadAccessSpecificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReadAccessSpecification, error) {
	v, err := (&_BACnetReadAccessSpecification{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetReadAccessSpecification) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetReadAccessSpecification BACnetReadAccessSpecification, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessSpecification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetReadAccessSpecification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfPropertyReferences, err := ReadTerminatedArrayField[BACnetPropertyReference](ctx, "listOfPropertyReferences", ReadComplex[BACnetPropertyReference](BACnetPropertyReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfPropertyReferences' field"))
	}
	m.ListOfPropertyReferences = listOfPropertyReferences

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetReadAccessSpecification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetReadAccessSpecification")
	}

	return m, nil
}

func (m *_BACnetReadAccessSpecification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetReadAccessSpecification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetReadAccessSpecification"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetReadAccessSpecification")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfPropertyReferences", m.GetListOfPropertyReferences(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfPropertyReferences' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessSpecification"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetReadAccessSpecification")
	}
	return nil
}

func (m *_BACnetReadAccessSpecification) isBACnetReadAccessSpecification() bool {
	return true
}

func (m *_BACnetReadAccessSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
