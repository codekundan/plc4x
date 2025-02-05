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

// BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
type BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter
	// IsBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged()
	// CreateBuilder creates a BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
	CreateBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
}

// _BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged struct {
	Header BACnetTagHeader
	Value  BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged = (*_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged)(nil)

// NewBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged factory function for _BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
func NewBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged(header BACnetTagHeader, value BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter, tagNumber uint8, tagClass TagClass) *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged must not be nil")
	}
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder is a builder for BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
type BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
	// Build builds the BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
}

// NewBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder() creates a BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
func NewBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder {
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder{_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged: new(_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged)}
}

type _BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder struct {
	*_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) = (*_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) WithValue(value BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) Build() (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) MustBuild() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder().(*_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder creates a BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder
func (b *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) CreateBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder()
	}
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedBuilder{_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetValue() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged(structType any) BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged"
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, error) {
	return BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, error) {
		return BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, error) {
	v, err := (&_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter](ctx, "value", readBuffer, EnsureType[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter_ALL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilter](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) IsBACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged() {
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) deepCopy() *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedCopy := &_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedCopy
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged) String() string {
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
