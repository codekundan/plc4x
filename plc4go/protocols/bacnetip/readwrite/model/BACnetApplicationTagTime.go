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

// BACnetApplicationTagTime is the corresponding interface of BACnetApplicationTagTime
type BACnetApplicationTagTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadTime
	// IsBACnetApplicationTagTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagTime()
	// CreateBuilder creates a BACnetApplicationTagTimeBuilder
	CreateBACnetApplicationTagTimeBuilder() BACnetApplicationTagTimeBuilder
}

// _BACnetApplicationTagTime is the data-structure of this message
type _BACnetApplicationTagTime struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadTime
}

var _ BACnetApplicationTagTime = (*_BACnetApplicationTagTime)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagTime)(nil)

// NewBACnetApplicationTagTime factory function for _BACnetApplicationTagTime
func NewBACnetApplicationTagTime(header BACnetTagHeader, payload BACnetTagPayloadTime) *_BACnetApplicationTagTime {
	if payload == nil {
		panic("payload of type BACnetTagPayloadTime for BACnetApplicationTagTime must not be nil")
	}
	_result := &_BACnetApplicationTagTime{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagTimeBuilder is a builder for BACnetApplicationTagTime
type BACnetApplicationTagTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadTime) BACnetApplicationTagTimeBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadTime) BACnetApplicationTagTimeBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadTimeBuilder) BACnetTagPayloadTimeBuilder) BACnetApplicationTagTimeBuilder
	// Build builds the BACnetApplicationTagTime or returns an error if something is wrong
	Build() (BACnetApplicationTagTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagTime
}

// NewBACnetApplicationTagTimeBuilder() creates a BACnetApplicationTagTimeBuilder
func NewBACnetApplicationTagTimeBuilder() BACnetApplicationTagTimeBuilder {
	return &_BACnetApplicationTagTimeBuilder{_BACnetApplicationTagTime: new(_BACnetApplicationTagTime)}
}

type _BACnetApplicationTagTimeBuilder struct {
	*_BACnetApplicationTagTime

	parentBuilder *_BACnetApplicationTagBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagTimeBuilder) = (*_BACnetApplicationTagTimeBuilder)(nil)

func (b *_BACnetApplicationTagTimeBuilder) setParent(contract BACnetApplicationTagContract) {
	b.BACnetApplicationTagContract = contract
}

func (b *_BACnetApplicationTagTimeBuilder) WithMandatoryFields(payload BACnetTagPayloadTime) BACnetApplicationTagTimeBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetApplicationTagTimeBuilder) WithPayload(payload BACnetTagPayloadTime) BACnetApplicationTagTimeBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetApplicationTagTimeBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadTimeBuilder) BACnetTagPayloadTimeBuilder) BACnetApplicationTagTimeBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadTimeBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagTimeBuilder) Build() (BACnetApplicationTagTime, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTagTime.deepCopy(), nil
}

func (b *_BACnetApplicationTagTimeBuilder) MustBuild() BACnetApplicationTagTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetApplicationTagTimeBuilder) Done() BACnetApplicationTagBuilder {
	return b.parentBuilder
}

func (b *_BACnetApplicationTagTimeBuilder) buildForBACnetApplicationTag() (BACnetApplicationTag, error) {
	return b.Build()
}

func (b *_BACnetApplicationTagTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagTimeBuilder().(*_BACnetApplicationTagTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagTimeBuilder creates a BACnetApplicationTagTimeBuilder
func (b *_BACnetApplicationTagTime) CreateBACnetApplicationTagTimeBuilder() BACnetApplicationTagTimeBuilder {
	if b == nil {
		return NewBACnetApplicationTagTimeBuilder()
	}
	return &_BACnetApplicationTagTimeBuilder{_BACnetApplicationTagTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagTime) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagTime) GetPayload() BACnetTagPayloadTime {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagTime(structType any) BACnetApplicationTagTime {
	if casted, ok := structType.(BACnetApplicationTagTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagTime) GetTypeName() string {
	return "BACnetApplicationTagTime"
}

func (m *_BACnetApplicationTagTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetApplicationTagTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag) (__bACnetApplicationTagTime BACnetApplicationTagTime, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadTime](ctx, "payload", ReadComplex[BACnetTagPayloadTime](BACnetTagPayloadTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagTime")
	}

	return m, nil
}

func (m *_BACnetApplicationTagTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagTime")
		}

		if err := WriteSimpleField[BACnetTagPayloadTime](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagTime")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagTime) IsBACnetApplicationTagTime() {}

func (m *_BACnetApplicationTagTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagTime) deepCopy() *_BACnetApplicationTagTime {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagTimeCopy := &_BACnetApplicationTagTime{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadTime),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagTimeCopy
}

func (m *_BACnetApplicationTagTime) String() string {
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
