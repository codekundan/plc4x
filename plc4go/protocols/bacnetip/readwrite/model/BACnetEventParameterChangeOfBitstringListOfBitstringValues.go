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

// BACnetEventParameterChangeOfBitstringListOfBitstringValues is the corresponding interface of BACnetEventParameterChangeOfBitstringListOfBitstringValues
type BACnetEventParameterChangeOfBitstringListOfBitstringValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfBitstringValues returns ListOfBitstringValues (property field)
	GetListOfBitstringValues() []BACnetApplicationTagBitString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfBitstringListOfBitstringValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfBitstringListOfBitstringValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfBitstringListOfBitstringValuesExactly interface {
	BACnetEventParameterChangeOfBitstringListOfBitstringValues
	isBACnetEventParameterChangeOfBitstringListOfBitstringValues() bool
}

// _BACnetEventParameterChangeOfBitstringListOfBitstringValues is the data-structure of this message
type _BACnetEventParameterChangeOfBitstringListOfBitstringValues struct {
	OpeningTag            BACnetOpeningTag
	ListOfBitstringValues []BACnetApplicationTagBitString
	ClosingTag            BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfBitstringListOfBitstringValues = (*_BACnetEventParameterChangeOfBitstringListOfBitstringValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetListOfBitstringValues() []BACnetApplicationTagBitString {
	return m.ListOfBitstringValues
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfBitstringListOfBitstringValues factory function for _BACnetEventParameterChangeOfBitstringListOfBitstringValues
func NewBACnetEventParameterChangeOfBitstringListOfBitstringValues(openingTag BACnetOpeningTag, listOfBitstringValues []BACnetApplicationTagBitString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	return &_BACnetEventParameterChangeOfBitstringListOfBitstringValues{OpeningTag: openingTag, ListOfBitstringValues: listOfBitstringValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfBitstringListOfBitstringValues(structType any) BACnetEventParameterChangeOfBitstringListOfBitstringValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfBitstringListOfBitstringValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfBitstringListOfBitstringValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfBitstringListOfBitstringValues"
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfBitstringValues) > 0 {
		for _, element := range m.ListOfBitstringValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfBitstringListOfBitstringValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfBitstringListOfBitstringValues, error) {
	return BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfBitstringListOfBitstringValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfBitstringListOfBitstringValues, error) {
		return BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfBitstringListOfBitstringValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfBitstringListOfBitstringValues, error) {
	v, err := (&_BACnetEventParameterChangeOfBitstringListOfBitstringValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterChangeOfBitstringListOfBitstringValues BACnetEventParameterChangeOfBitstringListOfBitstringValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfBitstringValues, err := ReadTerminatedArrayField[BACnetApplicationTagBitString](ctx, "listOfBitstringValues", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfBitstringValues' field"))
	}
	m.ListOfBitstringValues = listOfBitstringValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfBitstringValues", m.GetListOfBitstringValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfBitstringValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfBitstringListOfBitstringValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfBitstringListOfBitstringValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) isBACnetEventParameterChangeOfBitstringListOfBitstringValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfBitstringListOfBitstringValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
