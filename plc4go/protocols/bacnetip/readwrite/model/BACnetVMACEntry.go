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

// BACnetVMACEntry is the corresponding interface of BACnetVMACEntry
type BACnetVMACEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetVirtualMacAddress returns VirtualMacAddress (property field)
	GetVirtualMacAddress() BACnetContextTagOctetString
	// GetNativeMacAddress returns NativeMacAddress (property field)
	GetNativeMacAddress() BACnetContextTagOctetString
}

// BACnetVMACEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetVMACEntry.
// This is useful for switch cases.
type BACnetVMACEntryExactly interface {
	BACnetVMACEntry
	isBACnetVMACEntry() bool
}

// _BACnetVMACEntry is the data-structure of this message
type _BACnetVMACEntry struct {
	VirtualMacAddress BACnetContextTagOctetString
	NativeMacAddress  BACnetContextTagOctetString
}

var _ BACnetVMACEntry = (*_BACnetVMACEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetVMACEntry) GetVirtualMacAddress() BACnetContextTagOctetString {
	return m.VirtualMacAddress
}

func (m *_BACnetVMACEntry) GetNativeMacAddress() BACnetContextTagOctetString {
	return m.NativeMacAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetVMACEntry factory function for _BACnetVMACEntry
func NewBACnetVMACEntry(virtualMacAddress BACnetContextTagOctetString, nativeMacAddress BACnetContextTagOctetString) *_BACnetVMACEntry {
	return &_BACnetVMACEntry{VirtualMacAddress: virtualMacAddress, NativeMacAddress: nativeMacAddress}
}

// Deprecated: use the interface for direct cast
func CastBACnetVMACEntry(structType any) BACnetVMACEntry {
	if casted, ok := structType.(BACnetVMACEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetVMACEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetVMACEntry) GetTypeName() string {
	return "BACnetVMACEntry"
}

func (m *_BACnetVMACEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (virtualMacAddress)
	if m.VirtualMacAddress != nil {
		lengthInBits += m.VirtualMacAddress.GetLengthInBits(ctx)
	}

	// Optional Field (nativeMacAddress)
	if m.NativeMacAddress != nil {
		lengthInBits += m.NativeMacAddress.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetVMACEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetVMACEntryParse(ctx context.Context, theBytes []byte) (BACnetVMACEntry, error) {
	return BACnetVMACEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetVMACEntryParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVMACEntry, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVMACEntry, error) {
		return BACnetVMACEntryParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetVMACEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetVMACEntry, error) {
	v, err := (&_BACnetVMACEntry{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetVMACEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetVMACEntry BACnetVMACEntry, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetVMACEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetVMACEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var virtualMacAddress BACnetContextTagOctetString
	_virtualMacAddress, err := ReadOptionalField[BACnetContextTagOctetString](ctx, "virtualMacAddress", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'virtualMacAddress' field"))
	}
	if _virtualMacAddress != nil {
		virtualMacAddress = *_virtualMacAddress
		m.VirtualMacAddress = virtualMacAddress
	}

	var nativeMacAddress BACnetContextTagOctetString
	_nativeMacAddress, err := ReadOptionalField[BACnetContextTagOctetString](ctx, "nativeMacAddress", ReadComplex[BACnetContextTagOctetString](BACnetContextTagParseWithBufferProducer[BACnetContextTagOctetString]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_OCTET_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nativeMacAddress' field"))
	}
	if _nativeMacAddress != nil {
		nativeMacAddress = *_nativeMacAddress
		m.NativeMacAddress = nativeMacAddress
	}

	if closeErr := readBuffer.CloseContext("BACnetVMACEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetVMACEntry")
	}

	return m, nil
}

func (m *_BACnetVMACEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetVMACEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetVMACEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetVMACEntry")
	}

	if err := WriteOptionalField[BACnetContextTagOctetString](ctx, "virtualMacAddress", GetRef(m.GetVirtualMacAddress()), WriteComplex[BACnetContextTagOctetString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'virtualMacAddress' field")
	}

	if err := WriteOptionalField[BACnetContextTagOctetString](ctx, "nativeMacAddress", GetRef(m.GetNativeMacAddress()), WriteComplex[BACnetContextTagOctetString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'nativeMacAddress' field")
	}

	if popErr := writeBuffer.PopContext("BACnetVMACEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetVMACEntry")
	}
	return nil
}

func (m *_BACnetVMACEntry) isBACnetVMACEntry() bool {
	return true
}

func (m *_BACnetVMACEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
