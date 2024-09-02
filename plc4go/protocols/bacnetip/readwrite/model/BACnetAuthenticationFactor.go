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

// BACnetAuthenticationFactor is the corresponding interface of BACnetAuthenticationFactor
type BACnetAuthenticationFactor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFormatType returns FormatType (property field)
	GetFormatType() BACnetAuthenticationFactorTypeTagged
	// GetFormatClass returns FormatClass (property field)
	GetFormatClass() BACnetContextTagUnsignedInteger
	// GetValue returns Value (property field)
	GetValue() BACnetContextTagOctetString
}

// BACnetAuthenticationFactorExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationFactor.
// This is useful for switch cases.
type BACnetAuthenticationFactorExactly interface {
	BACnetAuthenticationFactor
	isBACnetAuthenticationFactor() bool
}

// _BACnetAuthenticationFactor is the data-structure of this message
type _BACnetAuthenticationFactor struct {
	FormatType  BACnetAuthenticationFactorTypeTagged
	FormatClass BACnetContextTagUnsignedInteger
	Value       BACnetContextTagOctetString
}

var _ BACnetAuthenticationFactor = (*_BACnetAuthenticationFactor)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactor) GetFormatType() BACnetAuthenticationFactorTypeTagged {
	return m.FormatType
}

func (m *_BACnetAuthenticationFactor) GetFormatClass() BACnetContextTagUnsignedInteger {
	return m.FormatClass
}

func (m *_BACnetAuthenticationFactor) GetValue() BACnetContextTagOctetString {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationFactor factory function for _BACnetAuthenticationFactor
func NewBACnetAuthenticationFactor(formatType BACnetAuthenticationFactorTypeTagged, formatClass BACnetContextTagUnsignedInteger, value BACnetContextTagOctetString) *_BACnetAuthenticationFactor {
	return &_BACnetAuthenticationFactor{FormatType: formatType, FormatClass: formatClass, Value: value}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactor(structType any) BACnetAuthenticationFactor {
	if casted, ok := structType.(BACnetAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactor) GetTypeName() string {
	return "BACnetAuthenticationFactor"
}

func (m *_BACnetAuthenticationFactor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (formatType)
	lengthInBits += m.FormatType.GetLengthInBits(ctx)

	// Simple field (formatClass)
	lengthInBits += m.FormatClass.GetLengthInBits(ctx)

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationFactor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationFactor, error) {
	return BACnetAuthenticationFactorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationFactorParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactor, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactor, error) {
		return BACnetAuthenticationFactorParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAuthenticationFactorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactor, error) {
	v, err := (&_BACnetAuthenticationFactor{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetAuthenticationFactor) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAuthenticationFactor BACnetAuthenticationFactor, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	formatType, err := ReadSimpleField[BACnetAuthenticationFactorTypeTagged](ctx, "formatType", ReadComplex[BACnetAuthenticationFactorTypeTagged](BACnetAuthenticationFactorTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'formatType' field"))
	}
	m.FormatType = formatType

	formatClass, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "formatClass", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'formatClass' field"))
	}
	m.FormatClass = formatClass

	value, err := ReadSimpleField[BACnetContextTagOctetString](ctx, "value", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactor")
	}

	return m, nil
}

func (m *_BACnetAuthenticationFactor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactor")
	}

	if err := WriteSimpleField[BACnetAuthenticationFactorTypeTagged](ctx, "formatType", m.GetFormatType(), WriteComplex[BACnetAuthenticationFactorTypeTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'formatType' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "formatClass", m.GetFormatClass(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'formatClass' field")
	}

	if err := WriteSimpleField[BACnetContextTagOctetString](ctx, "value", m.GetValue(), WriteComplex[BACnetContextTagOctetString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactor")
	}
	return nil
}

func (m *_BACnetAuthenticationFactor) isBACnetAuthenticationFactor() bool {
	return true
}

func (m *_BACnetAuthenticationFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
