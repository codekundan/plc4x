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

// DF1RequestProtectedTypedLogicalRead is the corresponding interface of DF1RequestProtectedTypedLogicalRead
type DF1RequestProtectedTypedLogicalRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1RequestCommand
	// GetByteSize returns ByteSize (property field)
	GetByteSize() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint8
	// GetFileType returns FileType (property field)
	GetFileType() uint8
	// GetElementNumber returns ElementNumber (property field)
	GetElementNumber() uint8
	// GetSubElementNumber returns SubElementNumber (property field)
	GetSubElementNumber() uint8
	// IsDF1RequestProtectedTypedLogicalRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1RequestProtectedTypedLogicalRead()
	// CreateBuilder creates a DF1RequestProtectedTypedLogicalReadBuilder
	CreateDF1RequestProtectedTypedLogicalReadBuilder() DF1RequestProtectedTypedLogicalReadBuilder
}

// _DF1RequestProtectedTypedLogicalRead is the data-structure of this message
type _DF1RequestProtectedTypedLogicalRead struct {
	DF1RequestCommandContract
	ByteSize         uint8
	FileNumber       uint8
	FileType         uint8
	ElementNumber    uint8
	SubElementNumber uint8
}

var _ DF1RequestProtectedTypedLogicalRead = (*_DF1RequestProtectedTypedLogicalRead)(nil)
var _ DF1RequestCommandRequirements = (*_DF1RequestProtectedTypedLogicalRead)(nil)

// NewDF1RequestProtectedTypedLogicalRead factory function for _DF1RequestProtectedTypedLogicalRead
func NewDF1RequestProtectedTypedLogicalRead(byteSize uint8, fileNumber uint8, fileType uint8, elementNumber uint8, subElementNumber uint8) *_DF1RequestProtectedTypedLogicalRead {
	_result := &_DF1RequestProtectedTypedLogicalRead{
		DF1RequestCommandContract: NewDF1RequestCommand(),
		ByteSize:                  byteSize,
		FileNumber:                fileNumber,
		FileType:                  fileType,
		ElementNumber:             elementNumber,
		SubElementNumber:          subElementNumber,
	}
	_result.DF1RequestCommandContract.(*_DF1RequestCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1RequestProtectedTypedLogicalReadBuilder is a builder for DF1RequestProtectedTypedLogicalRead
type DF1RequestProtectedTypedLogicalReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(byteSize uint8, fileNumber uint8, fileType uint8, elementNumber uint8, subElementNumber uint8) DF1RequestProtectedTypedLogicalReadBuilder
	// WithByteSize adds ByteSize (property field)
	WithByteSize(uint8) DF1RequestProtectedTypedLogicalReadBuilder
	// WithFileNumber adds FileNumber (property field)
	WithFileNumber(uint8) DF1RequestProtectedTypedLogicalReadBuilder
	// WithFileType adds FileType (property field)
	WithFileType(uint8) DF1RequestProtectedTypedLogicalReadBuilder
	// WithElementNumber adds ElementNumber (property field)
	WithElementNumber(uint8) DF1RequestProtectedTypedLogicalReadBuilder
	// WithSubElementNumber adds SubElementNumber (property field)
	WithSubElementNumber(uint8) DF1RequestProtectedTypedLogicalReadBuilder
	// Build builds the DF1RequestProtectedTypedLogicalRead or returns an error if something is wrong
	Build() (DF1RequestProtectedTypedLogicalRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1RequestProtectedTypedLogicalRead
}

// NewDF1RequestProtectedTypedLogicalReadBuilder() creates a DF1RequestProtectedTypedLogicalReadBuilder
func NewDF1RequestProtectedTypedLogicalReadBuilder() DF1RequestProtectedTypedLogicalReadBuilder {
	return &_DF1RequestProtectedTypedLogicalReadBuilder{_DF1RequestProtectedTypedLogicalRead: new(_DF1RequestProtectedTypedLogicalRead)}
}

type _DF1RequestProtectedTypedLogicalReadBuilder struct {
	*_DF1RequestProtectedTypedLogicalRead

	parentBuilder *_DF1RequestCommandBuilder

	err *utils.MultiError
}

var _ (DF1RequestProtectedTypedLogicalReadBuilder) = (*_DF1RequestProtectedTypedLogicalReadBuilder)(nil)

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) setParent(contract DF1RequestCommandContract) {
	b.DF1RequestCommandContract = contract
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) WithMandatoryFields(byteSize uint8, fileNumber uint8, fileType uint8, elementNumber uint8, subElementNumber uint8) DF1RequestProtectedTypedLogicalReadBuilder {
	return b.WithByteSize(byteSize).WithFileNumber(fileNumber).WithFileType(fileType).WithElementNumber(elementNumber).WithSubElementNumber(subElementNumber)
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) WithByteSize(byteSize uint8) DF1RequestProtectedTypedLogicalReadBuilder {
	b.ByteSize = byteSize
	return b
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) WithFileNumber(fileNumber uint8) DF1RequestProtectedTypedLogicalReadBuilder {
	b.FileNumber = fileNumber
	return b
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) WithFileType(fileType uint8) DF1RequestProtectedTypedLogicalReadBuilder {
	b.FileType = fileType
	return b
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) WithElementNumber(elementNumber uint8) DF1RequestProtectedTypedLogicalReadBuilder {
	b.ElementNumber = elementNumber
	return b
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) WithSubElementNumber(subElementNumber uint8) DF1RequestProtectedTypedLogicalReadBuilder {
	b.SubElementNumber = subElementNumber
	return b
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) Build() (DF1RequestProtectedTypedLogicalRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1RequestProtectedTypedLogicalRead.deepCopy(), nil
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) MustBuild() DF1RequestProtectedTypedLogicalRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DF1RequestProtectedTypedLogicalReadBuilder) Done() DF1RequestCommandBuilder {
	return b.parentBuilder
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) buildForDF1RequestCommand() (DF1RequestCommand, error) {
	return b.Build()
}

func (b *_DF1RequestProtectedTypedLogicalReadBuilder) DeepCopy() any {
	_copy := b.CreateDF1RequestProtectedTypedLogicalReadBuilder().(*_DF1RequestProtectedTypedLogicalReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1RequestProtectedTypedLogicalReadBuilder creates a DF1RequestProtectedTypedLogicalReadBuilder
func (b *_DF1RequestProtectedTypedLogicalRead) CreateDF1RequestProtectedTypedLogicalReadBuilder() DF1RequestProtectedTypedLogicalReadBuilder {
	if b == nil {
		return NewDF1RequestProtectedTypedLogicalReadBuilder()
	}
	return &_DF1RequestProtectedTypedLogicalReadBuilder{_DF1RequestProtectedTypedLogicalRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1RequestProtectedTypedLogicalRead) GetFunctionCode() uint8 {
	return 0xA2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1RequestProtectedTypedLogicalRead) GetParent() DF1RequestCommandContract {
	return m.DF1RequestCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1RequestProtectedTypedLogicalRead) GetByteSize() uint8 {
	return m.ByteSize
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetFileNumber() uint8 {
	return m.FileNumber
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetFileType() uint8 {
	return m.FileType
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetElementNumber() uint8 {
	return m.ElementNumber
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetSubElementNumber() uint8 {
	return m.SubElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1RequestProtectedTypedLogicalRead(structType any) DF1RequestProtectedTypedLogicalRead {
	if casted, ok := structType.(DF1RequestProtectedTypedLogicalRead); ok {
		return casted
	}
	if casted, ok := structType.(*DF1RequestProtectedTypedLogicalRead); ok {
		return *casted
	}
	return nil
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1RequestProtectedTypedLogicalRead"
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1RequestCommandContract.(*_DF1RequestCommand).getLengthInBits(ctx))

	// Simple field (byteSize)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 8

	// Simple field (fileType)
	lengthInBits += 8

	// Simple field (elementNumber)
	lengthInBits += 8

	// Simple field (subElementNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DF1RequestProtectedTypedLogicalRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1RequestProtectedTypedLogicalRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1RequestCommand) (__dF1RequestProtectedTypedLogicalRead DF1RequestProtectedTypedLogicalRead, err error) {
	m.DF1RequestCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1RequestProtectedTypedLogicalRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1RequestProtectedTypedLogicalRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteSize, err := ReadSimpleField(ctx, "byteSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteSize' field"))
	}
	m.ByteSize = byteSize

	fileNumber, err := ReadSimpleField(ctx, "fileNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileNumber' field"))
	}
	m.FileNumber = fileNumber

	fileType, err := ReadSimpleField(ctx, "fileType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fileType' field"))
	}
	m.FileType = fileType

	elementNumber, err := ReadSimpleField(ctx, "elementNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementNumber' field"))
	}
	m.ElementNumber = elementNumber

	subElementNumber, err := ReadSimpleField(ctx, "subElementNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subElementNumber' field"))
	}
	m.SubElementNumber = subElementNumber

	if closeErr := readBuffer.CloseContext("DF1RequestProtectedTypedLogicalRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1RequestProtectedTypedLogicalRead")
	}

	return m, nil
}

func (m *_DF1RequestProtectedTypedLogicalRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1RequestProtectedTypedLogicalRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1RequestProtectedTypedLogicalRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1RequestProtectedTypedLogicalRead")
		}

		if err := WriteSimpleField[uint8](ctx, "byteSize", m.GetByteSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteSize' field")
		}

		if err := WriteSimpleField[uint8](ctx, "fileNumber", m.GetFileNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileNumber' field")
		}

		if err := WriteSimpleField[uint8](ctx, "fileType", m.GetFileType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'fileType' field")
		}

		if err := WriteSimpleField[uint8](ctx, "elementNumber", m.GetElementNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'elementNumber' field")
		}

		if err := WriteSimpleField[uint8](ctx, "subElementNumber", m.GetSubElementNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'subElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("DF1RequestProtectedTypedLogicalRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1RequestProtectedTypedLogicalRead")
		}
		return nil
	}
	return m.DF1RequestCommandContract.(*_DF1RequestCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1RequestProtectedTypedLogicalRead) IsDF1RequestProtectedTypedLogicalRead() {}

func (m *_DF1RequestProtectedTypedLogicalRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1RequestProtectedTypedLogicalRead) deepCopy() *_DF1RequestProtectedTypedLogicalRead {
	if m == nil {
		return nil
	}
	_DF1RequestProtectedTypedLogicalReadCopy := &_DF1RequestProtectedTypedLogicalRead{
		m.DF1RequestCommandContract.(*_DF1RequestCommand).deepCopy(),
		m.ByteSize,
		m.FileNumber,
		m.FileType,
		m.ElementNumber,
		m.SubElementNumber,
	}
	m.DF1RequestCommandContract.(*_DF1RequestCommand)._SubType = m
	return _DF1RequestProtectedTypedLogicalReadCopy
}

func (m *_DF1RequestProtectedTypedLogicalRead) String() string {
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
