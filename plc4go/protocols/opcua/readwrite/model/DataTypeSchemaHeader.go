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

// DataTypeSchemaHeader is the corresponding interface of DataTypeSchemaHeader
type DataTypeSchemaHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNoOfNamespaces returns NoOfNamespaces (property field)
	GetNoOfNamespaces() int32
	// GetNamespaces returns Namespaces (property field)
	GetNamespaces() []PascalString
	// GetNoOfStructureDataTypes returns NoOfStructureDataTypes (property field)
	GetNoOfStructureDataTypes() int32
	// GetStructureDataTypes returns StructureDataTypes (property field)
	GetStructureDataTypes() []DataTypeDescription
	// GetNoOfEnumDataTypes returns NoOfEnumDataTypes (property field)
	GetNoOfEnumDataTypes() int32
	// GetEnumDataTypes returns EnumDataTypes (property field)
	GetEnumDataTypes() []DataTypeDescription
	// GetNoOfSimpleDataTypes returns NoOfSimpleDataTypes (property field)
	GetNoOfSimpleDataTypes() int32
	// GetSimpleDataTypes returns SimpleDataTypes (property field)
	GetSimpleDataTypes() []DataTypeDescription
	// IsDataTypeSchemaHeader is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataTypeSchemaHeader()
	// CreateBuilder creates a DataTypeSchemaHeaderBuilder
	CreateDataTypeSchemaHeaderBuilder() DataTypeSchemaHeaderBuilder
}

// _DataTypeSchemaHeader is the data-structure of this message
type _DataTypeSchemaHeader struct {
	ExtensionObjectDefinitionContract
	NoOfNamespaces         int32
	Namespaces             []PascalString
	NoOfStructureDataTypes int32
	StructureDataTypes     []DataTypeDescription
	NoOfEnumDataTypes      int32
	EnumDataTypes          []DataTypeDescription
	NoOfSimpleDataTypes    int32
	SimpleDataTypes        []DataTypeDescription
}

var _ DataTypeSchemaHeader = (*_DataTypeSchemaHeader)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataTypeSchemaHeader)(nil)

// NewDataTypeSchemaHeader factory function for _DataTypeSchemaHeader
func NewDataTypeSchemaHeader(noOfNamespaces int32, namespaces []PascalString, noOfStructureDataTypes int32, structureDataTypes []DataTypeDescription, noOfEnumDataTypes int32, enumDataTypes []DataTypeDescription, noOfSimpleDataTypes int32, simpleDataTypes []DataTypeDescription) *_DataTypeSchemaHeader {
	_result := &_DataTypeSchemaHeader{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfNamespaces:                    noOfNamespaces,
		Namespaces:                        namespaces,
		NoOfStructureDataTypes:            noOfStructureDataTypes,
		StructureDataTypes:                structureDataTypes,
		NoOfEnumDataTypes:                 noOfEnumDataTypes,
		EnumDataTypes:                     enumDataTypes,
		NoOfSimpleDataTypes:               noOfSimpleDataTypes,
		SimpleDataTypes:                   simpleDataTypes,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataTypeSchemaHeaderBuilder is a builder for DataTypeSchemaHeader
type DataTypeSchemaHeaderBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(noOfNamespaces int32, namespaces []PascalString, noOfStructureDataTypes int32, structureDataTypes []DataTypeDescription, noOfEnumDataTypes int32, enumDataTypes []DataTypeDescription, noOfSimpleDataTypes int32, simpleDataTypes []DataTypeDescription) DataTypeSchemaHeaderBuilder
	// WithNoOfNamespaces adds NoOfNamespaces (property field)
	WithNoOfNamespaces(int32) DataTypeSchemaHeaderBuilder
	// WithNamespaces adds Namespaces (property field)
	WithNamespaces(...PascalString) DataTypeSchemaHeaderBuilder
	// WithNoOfStructureDataTypes adds NoOfStructureDataTypes (property field)
	WithNoOfStructureDataTypes(int32) DataTypeSchemaHeaderBuilder
	// WithStructureDataTypes adds StructureDataTypes (property field)
	WithStructureDataTypes(...DataTypeDescription) DataTypeSchemaHeaderBuilder
	// WithNoOfEnumDataTypes adds NoOfEnumDataTypes (property field)
	WithNoOfEnumDataTypes(int32) DataTypeSchemaHeaderBuilder
	// WithEnumDataTypes adds EnumDataTypes (property field)
	WithEnumDataTypes(...DataTypeDescription) DataTypeSchemaHeaderBuilder
	// WithNoOfSimpleDataTypes adds NoOfSimpleDataTypes (property field)
	WithNoOfSimpleDataTypes(int32) DataTypeSchemaHeaderBuilder
	// WithSimpleDataTypes adds SimpleDataTypes (property field)
	WithSimpleDataTypes(...DataTypeDescription) DataTypeSchemaHeaderBuilder
	// Build builds the DataTypeSchemaHeader or returns an error if something is wrong
	Build() (DataTypeSchemaHeader, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataTypeSchemaHeader
}

// NewDataTypeSchemaHeaderBuilder() creates a DataTypeSchemaHeaderBuilder
func NewDataTypeSchemaHeaderBuilder() DataTypeSchemaHeaderBuilder {
	return &_DataTypeSchemaHeaderBuilder{_DataTypeSchemaHeader: new(_DataTypeSchemaHeader)}
}

type _DataTypeSchemaHeaderBuilder struct {
	*_DataTypeSchemaHeader

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DataTypeSchemaHeaderBuilder) = (*_DataTypeSchemaHeaderBuilder)(nil)

func (b *_DataTypeSchemaHeaderBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DataTypeSchemaHeaderBuilder) WithMandatoryFields(noOfNamespaces int32, namespaces []PascalString, noOfStructureDataTypes int32, structureDataTypes []DataTypeDescription, noOfEnumDataTypes int32, enumDataTypes []DataTypeDescription, noOfSimpleDataTypes int32, simpleDataTypes []DataTypeDescription) DataTypeSchemaHeaderBuilder {
	return b.WithNoOfNamespaces(noOfNamespaces).WithNamespaces(namespaces...).WithNoOfStructureDataTypes(noOfStructureDataTypes).WithStructureDataTypes(structureDataTypes...).WithNoOfEnumDataTypes(noOfEnumDataTypes).WithEnumDataTypes(enumDataTypes...).WithNoOfSimpleDataTypes(noOfSimpleDataTypes).WithSimpleDataTypes(simpleDataTypes...)
}

func (b *_DataTypeSchemaHeaderBuilder) WithNoOfNamespaces(noOfNamespaces int32) DataTypeSchemaHeaderBuilder {
	b.NoOfNamespaces = noOfNamespaces
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithNamespaces(namespaces ...PascalString) DataTypeSchemaHeaderBuilder {
	b.Namespaces = namespaces
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithNoOfStructureDataTypes(noOfStructureDataTypes int32) DataTypeSchemaHeaderBuilder {
	b.NoOfStructureDataTypes = noOfStructureDataTypes
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithStructureDataTypes(structureDataTypes ...DataTypeDescription) DataTypeSchemaHeaderBuilder {
	b.StructureDataTypes = structureDataTypes
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithNoOfEnumDataTypes(noOfEnumDataTypes int32) DataTypeSchemaHeaderBuilder {
	b.NoOfEnumDataTypes = noOfEnumDataTypes
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithEnumDataTypes(enumDataTypes ...DataTypeDescription) DataTypeSchemaHeaderBuilder {
	b.EnumDataTypes = enumDataTypes
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithNoOfSimpleDataTypes(noOfSimpleDataTypes int32) DataTypeSchemaHeaderBuilder {
	b.NoOfSimpleDataTypes = noOfSimpleDataTypes
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) WithSimpleDataTypes(simpleDataTypes ...DataTypeDescription) DataTypeSchemaHeaderBuilder {
	b.SimpleDataTypes = simpleDataTypes
	return b
}

func (b *_DataTypeSchemaHeaderBuilder) Build() (DataTypeSchemaHeader, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataTypeSchemaHeader.deepCopy(), nil
}

func (b *_DataTypeSchemaHeaderBuilder) MustBuild() DataTypeSchemaHeader {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DataTypeSchemaHeaderBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_DataTypeSchemaHeaderBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DataTypeSchemaHeaderBuilder) DeepCopy() any {
	_copy := b.CreateDataTypeSchemaHeaderBuilder().(*_DataTypeSchemaHeaderBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataTypeSchemaHeaderBuilder creates a DataTypeSchemaHeaderBuilder
func (b *_DataTypeSchemaHeader) CreateDataTypeSchemaHeaderBuilder() DataTypeSchemaHeaderBuilder {
	if b == nil {
		return NewDataTypeSchemaHeaderBuilder()
	}
	return &_DataTypeSchemaHeaderBuilder{_DataTypeSchemaHeader: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataTypeSchemaHeader) GetIdentifier() string {
	return "15536"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataTypeSchemaHeader) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataTypeSchemaHeader) GetNoOfNamespaces() int32 {
	return m.NoOfNamespaces
}

func (m *_DataTypeSchemaHeader) GetNamespaces() []PascalString {
	return m.Namespaces
}

func (m *_DataTypeSchemaHeader) GetNoOfStructureDataTypes() int32 {
	return m.NoOfStructureDataTypes
}

func (m *_DataTypeSchemaHeader) GetStructureDataTypes() []DataTypeDescription {
	return m.StructureDataTypes
}

func (m *_DataTypeSchemaHeader) GetNoOfEnumDataTypes() int32 {
	return m.NoOfEnumDataTypes
}

func (m *_DataTypeSchemaHeader) GetEnumDataTypes() []DataTypeDescription {
	return m.EnumDataTypes
}

func (m *_DataTypeSchemaHeader) GetNoOfSimpleDataTypes() int32 {
	return m.NoOfSimpleDataTypes
}

func (m *_DataTypeSchemaHeader) GetSimpleDataTypes() []DataTypeDescription {
	return m.SimpleDataTypes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataTypeSchemaHeader(structType any) DataTypeSchemaHeader {
	if casted, ok := structType.(DataTypeSchemaHeader); ok {
		return casted
	}
	if casted, ok := structType.(*DataTypeSchemaHeader); ok {
		return *casted
	}
	return nil
}

func (m *_DataTypeSchemaHeader) GetTypeName() string {
	return "DataTypeSchemaHeader"
}

func (m *_DataTypeSchemaHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfNamespaces)
	lengthInBits += 32

	// Array field
	if len(m.Namespaces) > 0 {
		for _curItem, element := range m.Namespaces {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Namespaces), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfStructureDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.StructureDataTypes) > 0 {
		for _curItem, element := range m.StructureDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StructureDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfEnumDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.EnumDataTypes) > 0 {
		for _curItem, element := range m.EnumDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EnumDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfSimpleDataTypes)
	lengthInBits += 32

	// Array field
	if len(m.SimpleDataTypes) > 0 {
		for _curItem, element := range m.SimpleDataTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SimpleDataTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DataTypeSchemaHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataTypeSchemaHeader) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataTypeSchemaHeader DataTypeSchemaHeader, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataTypeSchemaHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataTypeSchemaHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfNamespaces, err := ReadSimpleField(ctx, "noOfNamespaces", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNamespaces' field"))
	}
	m.NoOfNamespaces = noOfNamespaces

	namespaces, err := ReadCountArrayField[PascalString](ctx, "namespaces", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfNamespaces))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaces' field"))
	}
	m.Namespaces = namespaces

	noOfStructureDataTypes, err := ReadSimpleField(ctx, "noOfStructureDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfStructureDataTypes' field"))
	}
	m.NoOfStructureDataTypes = noOfStructureDataTypes

	structureDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "structureDataTypes", ReadComplex[DataTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[DataTypeDescription]((string)("14525")), readBuffer), uint64(noOfStructureDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'structureDataTypes' field"))
	}
	m.StructureDataTypes = structureDataTypes

	noOfEnumDataTypes, err := ReadSimpleField(ctx, "noOfEnumDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEnumDataTypes' field"))
	}
	m.NoOfEnumDataTypes = noOfEnumDataTypes

	enumDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "enumDataTypes", ReadComplex[DataTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[DataTypeDescription]((string)("14525")), readBuffer), uint64(noOfEnumDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumDataTypes' field"))
	}
	m.EnumDataTypes = enumDataTypes

	noOfSimpleDataTypes, err := ReadSimpleField(ctx, "noOfSimpleDataTypes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSimpleDataTypes' field"))
	}
	m.NoOfSimpleDataTypes = noOfSimpleDataTypes

	simpleDataTypes, err := ReadCountArrayField[DataTypeDescription](ctx, "simpleDataTypes", ReadComplex[DataTypeDescription](ExtensionObjectDefinitionParseWithBufferProducer[DataTypeDescription]((string)("14525")), readBuffer), uint64(noOfSimpleDataTypes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'simpleDataTypes' field"))
	}
	m.SimpleDataTypes = simpleDataTypes

	if closeErr := readBuffer.CloseContext("DataTypeSchemaHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataTypeSchemaHeader")
	}

	return m, nil
}

func (m *_DataTypeSchemaHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataTypeSchemaHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataTypeSchemaHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataTypeSchemaHeader")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNamespaces", m.GetNoOfNamespaces(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNamespaces' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "namespaces", m.GetNamespaces(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaces' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfStructureDataTypes", m.GetNoOfStructureDataTypes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfStructureDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "structureDataTypes", m.GetStructureDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'structureDataTypes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfEnumDataTypes", m.GetNoOfEnumDataTypes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEnumDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "enumDataTypes", m.GetEnumDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'enumDataTypes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfSimpleDataTypes", m.GetNoOfSimpleDataTypes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSimpleDataTypes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "simpleDataTypes", m.GetSimpleDataTypes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'simpleDataTypes' field")
		}

		if popErr := writeBuffer.PopContext("DataTypeSchemaHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataTypeSchemaHeader")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataTypeSchemaHeader) IsDataTypeSchemaHeader() {}

func (m *_DataTypeSchemaHeader) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataTypeSchemaHeader) deepCopy() *_DataTypeSchemaHeader {
	if m == nil {
		return nil
	}
	_DataTypeSchemaHeaderCopy := &_DataTypeSchemaHeader{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NoOfNamespaces,
		utils.DeepCopySlice[PascalString, PascalString](m.Namespaces),
		m.NoOfStructureDataTypes,
		utils.DeepCopySlice[DataTypeDescription, DataTypeDescription](m.StructureDataTypes),
		m.NoOfEnumDataTypes,
		utils.DeepCopySlice[DataTypeDescription, DataTypeDescription](m.EnumDataTypes),
		m.NoOfSimpleDataTypes,
		utils.DeepCopySlice[DataTypeDescription, DataTypeDescription](m.SimpleDataTypes),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataTypeSchemaHeaderCopy
}

func (m *_DataTypeSchemaHeader) String() string {
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
