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

// HPAIDataEndpoint is the corresponding interface of HPAIDataEndpoint
type HPAIDataEndpoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHostProtocolCode returns HostProtocolCode (property field)
	GetHostProtocolCode() HostProtocolCode
	// GetIpAddress returns IpAddress (property field)
	GetIpAddress() IPAddress
	// GetIpPort returns IpPort (property field)
	GetIpPort() uint16
}

// HPAIDataEndpointExactly can be used when we want exactly this type and not a type which fulfills HPAIDataEndpoint.
// This is useful for switch cases.
type HPAIDataEndpointExactly interface {
	HPAIDataEndpoint
	isHPAIDataEndpoint() bool
}

// _HPAIDataEndpoint is the data-structure of this message
type _HPAIDataEndpoint struct {
	HostProtocolCode HostProtocolCode
	IpAddress        IPAddress
	IpPort           uint16
}

var _ HPAIDataEndpoint = (*_HPAIDataEndpoint)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HPAIDataEndpoint) GetHostProtocolCode() HostProtocolCode {
	return m.HostProtocolCode
}

func (m *_HPAIDataEndpoint) GetIpAddress() IPAddress {
	return m.IpAddress
}

func (m *_HPAIDataEndpoint) GetIpPort() uint16 {
	return m.IpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHPAIDataEndpoint factory function for _HPAIDataEndpoint
func NewHPAIDataEndpoint(hostProtocolCode HostProtocolCode, ipAddress IPAddress, ipPort uint16) *_HPAIDataEndpoint {
	return &_HPAIDataEndpoint{HostProtocolCode: hostProtocolCode, IpAddress: ipAddress, IpPort: ipPort}
}

// Deprecated: use the interface for direct cast
func CastHPAIDataEndpoint(structType any) HPAIDataEndpoint {
	if casted, ok := structType.(HPAIDataEndpoint); ok {
		return casted
	}
	if casted, ok := structType.(*HPAIDataEndpoint); ok {
		return *casted
	}
	return nil
}

func (m *_HPAIDataEndpoint) GetTypeName() string {
	return "HPAIDataEndpoint"
}

func (m *_HPAIDataEndpoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits(ctx)

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_HPAIDataEndpoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HPAIDataEndpointParse(ctx context.Context, theBytes []byte) (HPAIDataEndpoint, error) {
	return HPAIDataEndpointParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HPAIDataEndpointParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HPAIDataEndpoint, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HPAIDataEndpoint, error) {
		return HPAIDataEndpointParseWithBuffer(ctx, readBuffer)
	}
}

func HPAIDataEndpointParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HPAIDataEndpoint, error) {
	v, err := (&_HPAIDataEndpoint{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_HPAIDataEndpoint) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hPAIDataEndpoint HPAIDataEndpoint, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HPAIDataEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HPAIDataEndpoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	structureLength, err := ReadImplicitField[uint8](ctx, "structureLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureLength' field"))
	}
	_ = structureLength

	hostProtocolCode, err := ReadEnumField[HostProtocolCode](ctx, "hostProtocolCode", "HostProtocolCode", ReadEnum(HostProtocolCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hostProtocolCode' field"))
	}
	m.HostProtocolCode = hostProtocolCode

	ipAddress, err := ReadSimpleField[IPAddress](ctx, "ipAddress", ReadComplex[IPAddress](IPAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipAddress' field"))
	}
	m.IpAddress = ipAddress

	ipPort, err := ReadSimpleField(ctx, "ipPort", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipPort' field"))
	}
	m.IpPort = ipPort

	if closeErr := readBuffer.CloseContext("HPAIDataEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HPAIDataEndpoint")
	}

	return m, nil
}

func (m *_HPAIDataEndpoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HPAIDataEndpoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HPAIDataEndpoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HPAIDataEndpoint")
	}
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	if err := WriteImplicitField(ctx, "structureLength", structureLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'structureLength' field")
	}

	if err := WriteSimpleEnumField[HostProtocolCode](ctx, "hostProtocolCode", "HostProtocolCode", m.GetHostProtocolCode(), WriteEnum[HostProtocolCode, uint8](HostProtocolCode.GetValue, HostProtocolCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'hostProtocolCode' field")
	}

	if err := WriteSimpleField[IPAddress](ctx, "ipAddress", m.GetIpAddress(), WriteComplex[IPAddress](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'ipAddress' field")
	}

	if err := WriteSimpleField[uint16](ctx, "ipPort", m.GetIpPort(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'ipPort' field")
	}

	if popErr := writeBuffer.PopContext("HPAIDataEndpoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HPAIDataEndpoint")
	}
	return nil
}

func (m *_HPAIDataEndpoint) isHPAIDataEndpoint() bool {
	return true
}

func (m *_HPAIDataEndpoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
