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

// NullCommandResponse is the corresponding interface of NullCommandResponse
type NullCommandResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// IsNullCommandResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNullCommandResponse()
	// CreateBuilder creates a NullCommandResponseBuilder
	CreateNullCommandResponseBuilder() NullCommandResponseBuilder
}

// _NullCommandResponse is the data-structure of this message
type _NullCommandResponse struct {
	EipPacketContract
}

var _ NullCommandResponse = (*_NullCommandResponse)(nil)
var _ EipPacketRequirements = (*_NullCommandResponse)(nil)

// NewNullCommandResponse factory function for _NullCommandResponse
func NewNullCommandResponse(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_NullCommandResponse {
	_result := &_NullCommandResponse{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NullCommandResponseBuilder is a builder for NullCommandResponse
type NullCommandResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NullCommandResponseBuilder
	// Build builds the NullCommandResponse or returns an error if something is wrong
	Build() (NullCommandResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NullCommandResponse
}

// NewNullCommandResponseBuilder() creates a NullCommandResponseBuilder
func NewNullCommandResponseBuilder() NullCommandResponseBuilder {
	return &_NullCommandResponseBuilder{_NullCommandResponse: new(_NullCommandResponse)}
}

type _NullCommandResponseBuilder struct {
	*_NullCommandResponse

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (NullCommandResponseBuilder) = (*_NullCommandResponseBuilder)(nil)

func (b *_NullCommandResponseBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
}

func (b *_NullCommandResponseBuilder) WithMandatoryFields() NullCommandResponseBuilder {
	return b
}

func (b *_NullCommandResponseBuilder) Build() (NullCommandResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NullCommandResponse.deepCopy(), nil
}

func (b *_NullCommandResponseBuilder) MustBuild() NullCommandResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NullCommandResponseBuilder) Done() EipPacketBuilder {
	return b.parentBuilder
}

func (b *_NullCommandResponseBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_NullCommandResponseBuilder) DeepCopy() any {
	_copy := b.CreateNullCommandResponseBuilder().(*_NullCommandResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNullCommandResponseBuilder creates a NullCommandResponseBuilder
func (b *_NullCommandResponse) CreateNullCommandResponseBuilder() NullCommandResponseBuilder {
	if b == nil {
		return NewNullCommandResponseBuilder()
	}
	return &_NullCommandResponseBuilder{_NullCommandResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullCommandResponse) GetCommand() uint16 {
	return 0x0001
}

func (m *_NullCommandResponse) GetResponse() bool {
	return bool(true)
}

func (m *_NullCommandResponse) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullCommandResponse) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// Deprecated: use the interface for direct cast
func CastNullCommandResponse(structType any) NullCommandResponse {
	if casted, ok := structType.(NullCommandResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NullCommandResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NullCommandResponse) GetTypeName() string {
	return "NullCommandResponse"
}

func (m *_NullCommandResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullCommandResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NullCommandResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__nullCommandResponse NullCommandResponse, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullCommandResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullCommandResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullCommandResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullCommandResponse")
	}

	return m, nil
}

func (m *_NullCommandResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullCommandResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullCommandResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullCommandResponse")
		}

		if popErr := writeBuffer.PopContext("NullCommandResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullCommandResponse")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullCommandResponse) IsNullCommandResponse() {}

func (m *_NullCommandResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NullCommandResponse) deepCopy() *_NullCommandResponse {
	if m == nil {
		return nil
	}
	_NullCommandResponseCopy := &_NullCommandResponse{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
	}
	m.EipPacketContract.(*_EipPacket)._SubType = m
	return _NullCommandResponseCopy
}

func (m *_NullCommandResponse) String() string {
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
