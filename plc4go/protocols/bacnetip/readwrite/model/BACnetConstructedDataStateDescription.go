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

// BACnetConstructedDataStateDescription is the corresponding interface of BACnetConstructedDataStateDescription
type BACnetConstructedDataStateDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetStateDescription returns StateDescription (property field)
	GetStateDescription() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataStateDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStateDescription()
	// CreateBuilder creates a BACnetConstructedDataStateDescriptionBuilder
	CreateBACnetConstructedDataStateDescriptionBuilder() BACnetConstructedDataStateDescriptionBuilder
}

// _BACnetConstructedDataStateDescription is the data-structure of this message
type _BACnetConstructedDataStateDescription struct {
	BACnetConstructedDataContract
	StateDescription BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataStateDescription = (*_BACnetConstructedDataStateDescription)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStateDescription)(nil)

// NewBACnetConstructedDataStateDescription factory function for _BACnetConstructedDataStateDescription
func NewBACnetConstructedDataStateDescription(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, stateDescription BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStateDescription {
	if stateDescription == nil {
		panic("stateDescription of type BACnetApplicationTagCharacterString for BACnetConstructedDataStateDescription must not be nil")
	}
	_result := &_BACnetConstructedDataStateDescription{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StateDescription:              stateDescription,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStateDescriptionBuilder is a builder for BACnetConstructedDataStateDescription
type BACnetConstructedDataStateDescriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(stateDescription BACnetApplicationTagCharacterString) BACnetConstructedDataStateDescriptionBuilder
	// WithStateDescription adds StateDescription (property field)
	WithStateDescription(BACnetApplicationTagCharacterString) BACnetConstructedDataStateDescriptionBuilder
	// WithStateDescriptionBuilder adds StateDescription (property field) which is build by the builder
	WithStateDescriptionBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataStateDescriptionBuilder
	// Build builds the BACnetConstructedDataStateDescription or returns an error if something is wrong
	Build() (BACnetConstructedDataStateDescription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStateDescription
}

// NewBACnetConstructedDataStateDescriptionBuilder() creates a BACnetConstructedDataStateDescriptionBuilder
func NewBACnetConstructedDataStateDescriptionBuilder() BACnetConstructedDataStateDescriptionBuilder {
	return &_BACnetConstructedDataStateDescriptionBuilder{_BACnetConstructedDataStateDescription: new(_BACnetConstructedDataStateDescription)}
}

type _BACnetConstructedDataStateDescriptionBuilder struct {
	*_BACnetConstructedDataStateDescription

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStateDescriptionBuilder) = (*_BACnetConstructedDataStateDescriptionBuilder)(nil)

func (b *_BACnetConstructedDataStateDescriptionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) WithMandatoryFields(stateDescription BACnetApplicationTagCharacterString) BACnetConstructedDataStateDescriptionBuilder {
	return b.WithStateDescription(stateDescription)
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) WithStateDescription(stateDescription BACnetApplicationTagCharacterString) BACnetConstructedDataStateDescriptionBuilder {
	b.StateDescription = stateDescription
	return b
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) WithStateDescriptionBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataStateDescriptionBuilder {
	builder := builderSupplier(b.StateDescription.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.StateDescription, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) Build() (BACnetConstructedDataStateDescription, error) {
	if b.StateDescription == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'stateDescription' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStateDescription.deepCopy(), nil
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) MustBuild() BACnetConstructedDataStateDescription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataStateDescriptionBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStateDescriptionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStateDescriptionBuilder().(*_BACnetConstructedDataStateDescriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStateDescriptionBuilder creates a BACnetConstructedDataStateDescriptionBuilder
func (b *_BACnetConstructedDataStateDescription) CreateBACnetConstructedDataStateDescriptionBuilder() BACnetConstructedDataStateDescriptionBuilder {
	if b == nil {
		return NewBACnetConstructedDataStateDescriptionBuilder()
	}
	return &_BACnetConstructedDataStateDescriptionBuilder{_BACnetConstructedDataStateDescription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStateDescription) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStateDescription) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STATE_DESCRIPTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStateDescription) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStateDescription) GetStateDescription() BACnetApplicationTagCharacterString {
	return m.StateDescription
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStateDescription) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetStateDescription())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStateDescription(structType any) BACnetConstructedDataStateDescription {
	if casted, ok := structType.(BACnetConstructedDataStateDescription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStateDescription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStateDescription) GetTypeName() string {
	return "BACnetConstructedDataStateDescription"
}

func (m *_BACnetConstructedDataStateDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (stateDescription)
	lengthInBits += m.StateDescription.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStateDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStateDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStateDescription BACnetConstructedDataStateDescription, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStateDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStateDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	stateDescription, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "stateDescription", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stateDescription' field"))
	}
	m.StateDescription = stateDescription

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), stateDescription)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStateDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStateDescription")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStateDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStateDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStateDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStateDescription")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "stateDescription", m.GetStateDescription(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'stateDescription' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStateDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStateDescription")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStateDescription) IsBACnetConstructedDataStateDescription() {}

func (m *_BACnetConstructedDataStateDescription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStateDescription) deepCopy() *_BACnetConstructedDataStateDescription {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStateDescriptionCopy := &_BACnetConstructedDataStateDescription{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.StateDescription.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStateDescriptionCopy
}

func (m *_BACnetConstructedDataStateDescription) String() string {
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
