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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ModbusTcpADU_PROTOCOLIDENTIFIER uint16 = 0x0000

// ModbusTcpADU is the corresponding interface of ModbusTcpADU
type ModbusTcpADU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusADU
	// GetTransactionIdentifier returns TransactionIdentifier (property field)
	GetTransactionIdentifier() uint16
	// GetUnitIdentifier returns UnitIdentifier (property field)
	GetUnitIdentifier() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
	// IsModbusTcpADU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusTcpADU()
	// CreateBuilder creates a ModbusTcpADUBuilder
	CreateModbusTcpADUBuilder() ModbusTcpADUBuilder
}

// _ModbusTcpADU is the data-structure of this message
type _ModbusTcpADU struct {
	ModbusADUContract
	TransactionIdentifier uint16
	UnitIdentifier        uint8
	Pdu                   ModbusPDU
}

var _ ModbusTcpADU = (*_ModbusTcpADU)(nil)
var _ ModbusADURequirements = (*_ModbusTcpADU)(nil)

// NewModbusTcpADU factory function for _ModbusTcpADU
func NewModbusTcpADU(transactionIdentifier uint16, unitIdentifier uint8, pdu ModbusPDU, response bool) *_ModbusTcpADU {
	if pdu == nil {
		panic("pdu of type ModbusPDU for ModbusTcpADU must not be nil")
	}
	_result := &_ModbusTcpADU{
		ModbusADUContract:     NewModbusADU(response),
		TransactionIdentifier: transactionIdentifier,
		UnitIdentifier:        unitIdentifier,
		Pdu:                   pdu,
	}
	_result.ModbusADUContract.(*_ModbusADU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusTcpADUBuilder is a builder for ModbusTcpADU
type ModbusTcpADUBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(transactionIdentifier uint16, unitIdentifier uint8, pdu ModbusPDU) ModbusTcpADUBuilder
	// WithTransactionIdentifier adds TransactionIdentifier (property field)
	WithTransactionIdentifier(uint16) ModbusTcpADUBuilder
	// WithUnitIdentifier adds UnitIdentifier (property field)
	WithUnitIdentifier(uint8) ModbusTcpADUBuilder
	// WithPdu adds Pdu (property field)
	WithPdu(ModbusPDU) ModbusTcpADUBuilder
	// WithPduBuilder adds Pdu (property field) which is build by the builder
	WithPduBuilder(func(ModbusPDUBuilder) ModbusPDUBuilder) ModbusTcpADUBuilder
	// Build builds the ModbusTcpADU or returns an error if something is wrong
	Build() (ModbusTcpADU, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusTcpADU
}

// NewModbusTcpADUBuilder() creates a ModbusTcpADUBuilder
func NewModbusTcpADUBuilder() ModbusTcpADUBuilder {
	return &_ModbusTcpADUBuilder{_ModbusTcpADU: new(_ModbusTcpADU)}
}

type _ModbusTcpADUBuilder struct {
	*_ModbusTcpADU

	parentBuilder *_ModbusADUBuilder

	err *utils.MultiError
}

var _ (ModbusTcpADUBuilder) = (*_ModbusTcpADUBuilder)(nil)

func (b *_ModbusTcpADUBuilder) setParent(contract ModbusADUContract) {
	b.ModbusADUContract = contract
}

func (b *_ModbusTcpADUBuilder) WithMandatoryFields(transactionIdentifier uint16, unitIdentifier uint8, pdu ModbusPDU) ModbusTcpADUBuilder {
	return b.WithTransactionIdentifier(transactionIdentifier).WithUnitIdentifier(unitIdentifier).WithPdu(pdu)
}

func (b *_ModbusTcpADUBuilder) WithTransactionIdentifier(transactionIdentifier uint16) ModbusTcpADUBuilder {
	b.TransactionIdentifier = transactionIdentifier
	return b
}

func (b *_ModbusTcpADUBuilder) WithUnitIdentifier(unitIdentifier uint8) ModbusTcpADUBuilder {
	b.UnitIdentifier = unitIdentifier
	return b
}

func (b *_ModbusTcpADUBuilder) WithPdu(pdu ModbusPDU) ModbusTcpADUBuilder {
	b.Pdu = pdu
	return b
}

func (b *_ModbusTcpADUBuilder) WithPduBuilder(builderSupplier func(ModbusPDUBuilder) ModbusPDUBuilder) ModbusTcpADUBuilder {
	builder := builderSupplier(b.Pdu.CreateModbusPDUBuilder())
	var err error
	b.Pdu, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ModbusPDUBuilder failed"))
	}
	return b
}

func (b *_ModbusTcpADUBuilder) Build() (ModbusTcpADU, error) {
	if b.Pdu == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'pdu' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusTcpADU.deepCopy(), nil
}

func (b *_ModbusTcpADUBuilder) MustBuild() ModbusTcpADU {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ModbusTcpADUBuilder) Done() ModbusADUBuilder {
	return b.parentBuilder
}

func (b *_ModbusTcpADUBuilder) buildForModbusADU() (ModbusADU, error) {
	return b.Build()
}

func (b *_ModbusTcpADUBuilder) DeepCopy() any {
	_copy := b.CreateModbusTcpADUBuilder().(*_ModbusTcpADUBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusTcpADUBuilder creates a ModbusTcpADUBuilder
func (b *_ModbusTcpADU) CreateModbusTcpADUBuilder() ModbusTcpADUBuilder {
	if b == nil {
		return NewModbusTcpADUBuilder()
	}
	return &_ModbusTcpADUBuilder{_ModbusTcpADU: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusTcpADU) GetDriverType() DriverType {
	return DriverType_MODBUS_TCP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusTcpADU) GetParent() ModbusADUContract {
	return m.ModbusADUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusTcpADU) GetTransactionIdentifier() uint16 {
	return m.TransactionIdentifier
}

func (m *_ModbusTcpADU) GetUnitIdentifier() uint8 {
	return m.UnitIdentifier
}

func (m *_ModbusTcpADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusTcpADU) GetProtocolIdentifier() uint16 {
	return ModbusTcpADU_PROTOCOLIDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusTcpADU(structType any) ModbusTcpADU {
	if casted, ok := structType.(ModbusTcpADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusTcpADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusTcpADU) GetTypeName() string {
	return "ModbusTcpADU"
}

func (m *_ModbusTcpADU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusADUContract.(*_ModbusADU).getLengthInBits(ctx))

	// Simple field (transactionIdentifier)
	lengthInBits += 16

	// Const Field (protocolIdentifier)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 16

	// Simple field (unitIdentifier)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ModbusTcpADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusTcpADU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusADU, driverType DriverType, response bool) (__modbusTcpADU ModbusTcpADU, err error) {
	m.ModbusADUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusTcpADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusTcpADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	transactionIdentifier, err := ReadSimpleField(ctx, "transactionIdentifier", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionIdentifier' field"))
	}
	m.TransactionIdentifier = transactionIdentifier

	protocolIdentifier, err := ReadConstField[uint16](ctx, "protocolIdentifier", ReadUnsignedShort(readBuffer, uint8(16)), ModbusTcpADU_PROTOCOLIDENTIFIER, codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolIdentifier' field"))
	}
	_ = protocolIdentifier

	length, err := ReadImplicitField[uint16](ctx, "length", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	_ = length

	unitIdentifier, err := ReadSimpleField(ctx, "unitIdentifier", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitIdentifier' field"))
	}
	m.UnitIdentifier = unitIdentifier

	pdu, err := ReadSimpleField[ModbusPDU](ctx, "pdu", ReadComplex[ModbusPDU](ModbusPDUParseWithBufferProducer[ModbusPDU]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pdu' field"))
	}
	m.Pdu = pdu

	if closeErr := readBuffer.CloseContext("ModbusTcpADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusTcpADU")
	}

	return m, nil
}

func (m *_ModbusTcpADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusTcpADU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusTcpADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusTcpADU")
		}

		if err := WriteSimpleField[uint16](ctx, "transactionIdentifier", m.GetTransactionIdentifier(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'transactionIdentifier' field")
		}

		if err := WriteConstField(ctx, "protocolIdentifier", ModbusTcpADU_PROTOCOLIDENTIFIER, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolIdentifier' field")
		}
		length := uint16(uint16(m.GetPdu().GetLengthInBytes(ctx)) + uint16(uint16(1)))
		if err := WriteImplicitField(ctx, "length", length, WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleField[uint8](ctx, "unitIdentifier", m.GetUnitIdentifier(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'unitIdentifier' field")
		}

		if err := WriteSimpleField[ModbusPDU](ctx, "pdu", m.GetPdu(), WriteComplex[ModbusPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pdu' field")
		}

		if popErr := writeBuffer.PopContext("ModbusTcpADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusTcpADU")
		}
		return nil
	}
	return m.ModbusADUContract.(*_ModbusADU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusTcpADU) IsModbusTcpADU() {}

func (m *_ModbusTcpADU) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusTcpADU) deepCopy() *_ModbusTcpADU {
	if m == nil {
		return nil
	}
	_ModbusTcpADUCopy := &_ModbusTcpADU{
		m.ModbusADUContract.(*_ModbusADU).deepCopy(),
		m.TransactionIdentifier,
		m.UnitIdentifier,
		m.Pdu.DeepCopy().(ModbusPDU),
	}
	m.ModbusADUContract.(*_ModbusADU)._SubType = m
	return _ModbusTcpADUCopy
}

func (m *_ModbusTcpADU) String() string {
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
