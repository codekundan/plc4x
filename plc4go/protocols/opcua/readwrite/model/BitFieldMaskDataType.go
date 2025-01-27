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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BitFieldMaskDataType is the corresponding interface of BitFieldMaskDataType
type BitFieldMaskDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBitFieldMaskDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBitFieldMaskDataType()
	// CreateBuilder creates a BitFieldMaskDataTypeBuilder
	CreateBitFieldMaskDataTypeBuilder() BitFieldMaskDataTypeBuilder
}

// _BitFieldMaskDataType is the data-structure of this message
type _BitFieldMaskDataType struct {
}

var _ BitFieldMaskDataType = (*_BitFieldMaskDataType)(nil)

// NewBitFieldMaskDataType factory function for _BitFieldMaskDataType
func NewBitFieldMaskDataType() *_BitFieldMaskDataType {
	return &_BitFieldMaskDataType{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BitFieldMaskDataTypeBuilder is a builder for BitFieldMaskDataType
type BitFieldMaskDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BitFieldMaskDataTypeBuilder
	// Build builds the BitFieldMaskDataType or returns an error if something is wrong
	Build() (BitFieldMaskDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BitFieldMaskDataType
}

// NewBitFieldMaskDataTypeBuilder() creates a BitFieldMaskDataTypeBuilder
func NewBitFieldMaskDataTypeBuilder() BitFieldMaskDataTypeBuilder {
	return &_BitFieldMaskDataTypeBuilder{_BitFieldMaskDataType: new(_BitFieldMaskDataType)}
}

type _BitFieldMaskDataTypeBuilder struct {
	*_BitFieldMaskDataType

	err *utils.MultiError
}

var _ (BitFieldMaskDataTypeBuilder) = (*_BitFieldMaskDataTypeBuilder)(nil)

func (b *_BitFieldMaskDataTypeBuilder) WithMandatoryFields() BitFieldMaskDataTypeBuilder {
	return b
}

func (b *_BitFieldMaskDataTypeBuilder) Build() (BitFieldMaskDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BitFieldMaskDataType.deepCopy(), nil
}

func (b *_BitFieldMaskDataTypeBuilder) MustBuild() BitFieldMaskDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BitFieldMaskDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateBitFieldMaskDataTypeBuilder().(*_BitFieldMaskDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBitFieldMaskDataTypeBuilder creates a BitFieldMaskDataTypeBuilder
func (b *_BitFieldMaskDataType) CreateBitFieldMaskDataTypeBuilder() BitFieldMaskDataTypeBuilder {
	if b == nil {
		return NewBitFieldMaskDataTypeBuilder()
	}
	return &_BitFieldMaskDataTypeBuilder{_BitFieldMaskDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBitFieldMaskDataType(structType any) BitFieldMaskDataType {
	if casted, ok := structType.(BitFieldMaskDataType); ok {
		return casted
	}
	if casted, ok := structType.(*BitFieldMaskDataType); ok {
		return *casted
	}
	return nil
}

func (m *_BitFieldMaskDataType) GetTypeName() string {
	return "BitFieldMaskDataType"
}

func (m *_BitFieldMaskDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_BitFieldMaskDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BitFieldMaskDataTypeParse(ctx context.Context, theBytes []byte) (BitFieldMaskDataType, error) {
	return BitFieldMaskDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BitFieldMaskDataTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BitFieldMaskDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BitFieldMaskDataType, error) {
		return BitFieldMaskDataTypeParseWithBuffer(ctx, readBuffer)
	}
}

func BitFieldMaskDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BitFieldMaskDataType, error) {
	v, err := (&_BitFieldMaskDataType{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BitFieldMaskDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bitFieldMaskDataType BitFieldMaskDataType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BitFieldMaskDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BitFieldMaskDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BitFieldMaskDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BitFieldMaskDataType")
	}

	return m, nil
}

func (m *_BitFieldMaskDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BitFieldMaskDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BitFieldMaskDataType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BitFieldMaskDataType")
	}

	if popErr := writeBuffer.PopContext("BitFieldMaskDataType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BitFieldMaskDataType")
	}
	return nil
}

func (m *_BitFieldMaskDataType) IsBitFieldMaskDataType() {}

func (m *_BitFieldMaskDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BitFieldMaskDataType) deepCopy() *_BitFieldMaskDataType {
	if m == nil {
		return nil
	}
	_BitFieldMaskDataTypeCopy := &_BitFieldMaskDataType{}
	return _BitFieldMaskDataTypeCopy
}

func (m *_BitFieldMaskDataType) String() string {
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
