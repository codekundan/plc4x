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

// BACnetConstructedDataDeviceMaxMaster is the corresponding interface of BACnetConstructedDataDeviceMaxMaster
type BACnetConstructedDataDeviceMaxMaster interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxMaster returns MaxMaster (property field)
	GetMaxMaster() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataDeviceMaxMaster is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDeviceMaxMaster()
	// CreateBuilder creates a BACnetConstructedDataDeviceMaxMasterBuilder
	CreateBACnetConstructedDataDeviceMaxMasterBuilder() BACnetConstructedDataDeviceMaxMasterBuilder
}

// _BACnetConstructedDataDeviceMaxMaster is the data-structure of this message
type _BACnetConstructedDataDeviceMaxMaster struct {
	BACnetConstructedDataContract
	MaxMaster BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataDeviceMaxMaster = (*_BACnetConstructedDataDeviceMaxMaster)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDeviceMaxMaster)(nil)

// NewBACnetConstructedDataDeviceMaxMaster factory function for _BACnetConstructedDataDeviceMaxMaster
func NewBACnetConstructedDataDeviceMaxMaster(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxMaster BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeviceMaxMaster {
	if maxMaster == nil {
		panic("maxMaster of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataDeviceMaxMaster must not be nil")
	}
	_result := &_BACnetConstructedDataDeviceMaxMaster{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxMaster:                     maxMaster,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDeviceMaxMasterBuilder is a builder for BACnetConstructedDataDeviceMaxMaster
type BACnetConstructedDataDeviceMaxMasterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxMaster BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDeviceMaxMasterBuilder
	// WithMaxMaster adds MaxMaster (property field)
	WithMaxMaster(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDeviceMaxMasterBuilder
	// WithMaxMasterBuilder adds MaxMaster (property field) which is build by the builder
	WithMaxMasterBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDeviceMaxMasterBuilder
	// Build builds the BACnetConstructedDataDeviceMaxMaster or returns an error if something is wrong
	Build() (BACnetConstructedDataDeviceMaxMaster, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDeviceMaxMaster
}

// NewBACnetConstructedDataDeviceMaxMasterBuilder() creates a BACnetConstructedDataDeviceMaxMasterBuilder
func NewBACnetConstructedDataDeviceMaxMasterBuilder() BACnetConstructedDataDeviceMaxMasterBuilder {
	return &_BACnetConstructedDataDeviceMaxMasterBuilder{_BACnetConstructedDataDeviceMaxMaster: new(_BACnetConstructedDataDeviceMaxMaster)}
}

type _BACnetConstructedDataDeviceMaxMasterBuilder struct {
	*_BACnetConstructedDataDeviceMaxMaster

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDeviceMaxMasterBuilder) = (*_BACnetConstructedDataDeviceMaxMasterBuilder)(nil)

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) WithMandatoryFields(maxMaster BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDeviceMaxMasterBuilder {
	return b.WithMaxMaster(maxMaster)
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) WithMaxMaster(maxMaster BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDeviceMaxMasterBuilder {
	b.MaxMaster = maxMaster
	return b
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) WithMaxMasterBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDeviceMaxMasterBuilder {
	builder := builderSupplier(b.MaxMaster.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaxMaster, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) Build() (BACnetConstructedDataDeviceMaxMaster, error) {
	if b.MaxMaster == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxMaster' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDeviceMaxMaster.deepCopy(), nil
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) MustBuild() BACnetConstructedDataDeviceMaxMaster {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDeviceMaxMasterBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDeviceMaxMasterBuilder().(*_BACnetConstructedDataDeviceMaxMasterBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDeviceMaxMasterBuilder creates a BACnetConstructedDataDeviceMaxMasterBuilder
func (b *_BACnetConstructedDataDeviceMaxMaster) CreateBACnetConstructedDataDeviceMaxMasterBuilder() BACnetConstructedDataDeviceMaxMasterBuilder {
	if b == nil {
		return NewBACnetConstructedDataDeviceMaxMasterBuilder()
	}
	return &_BACnetConstructedDataDeviceMaxMasterBuilder{_BACnetConstructedDataDeviceMaxMaster: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxMaster) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DEVICE
}

func (m *_BACnetConstructedDataDeviceMaxMaster) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_MASTER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeviceMaxMaster) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxMaster) GetMaxMaster() BACnetApplicationTagUnsignedInteger {
	return m.MaxMaster
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxMaster) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxMaster())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeviceMaxMaster(structType any) BACnetConstructedDataDeviceMaxMaster {
	if casted, ok := structType.(BACnetConstructedDataDeviceMaxMaster); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceMaxMaster); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeviceMaxMaster) GetTypeName() string {
	return "BACnetConstructedDataDeviceMaxMaster"
}

func (m *_BACnetConstructedDataDeviceMaxMaster) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxMaster)
	lengthInBits += m.MaxMaster.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDeviceMaxMaster) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDeviceMaxMaster) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDeviceMaxMaster BACnetConstructedDataDeviceMaxMaster, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceMaxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceMaxMaster")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxMaster, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxMaster", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxMaster' field"))
	}
	m.MaxMaster = maxMaster

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxMaster)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceMaxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceMaxMaster")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDeviceMaxMaster) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeviceMaxMaster) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceMaxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceMaxMaster")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxMaster", m.GetMaxMaster(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxMaster' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceMaxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceMaxMaster")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeviceMaxMaster) IsBACnetConstructedDataDeviceMaxMaster() {}

func (m *_BACnetConstructedDataDeviceMaxMaster) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDeviceMaxMaster) deepCopy() *_BACnetConstructedDataDeviceMaxMaster {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDeviceMaxMasterCopy := &_BACnetConstructedDataDeviceMaxMaster{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MaxMaster.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDeviceMaxMasterCopy
}

func (m *_BACnetConstructedDataDeviceMaxMaster) String() string {
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
