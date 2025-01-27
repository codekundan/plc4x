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

// HVACModeAndFlags is the corresponding interface of HVACModeAndFlags
type HVACModeAndFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetAuxiliaryLevel returns AuxiliaryLevel (property field)
	GetAuxiliaryLevel() bool
	// GetGuard returns Guard (property field)
	GetGuard() bool
	// GetSetback returns Setback (property field)
	GetSetback() bool
	// GetLevel returns Level (property field)
	GetLevel() bool
	// GetMode returns Mode (property field)
	GetMode() HVACModeAndFlagsMode
	// GetIsAuxLevelUnused returns IsAuxLevelUnused (virtual field)
	GetIsAuxLevelUnused() bool
	// GetIsAuxLevelUsed returns IsAuxLevelUsed (virtual field)
	GetIsAuxLevelUsed() bool
	// GetIsGuardDisabled returns IsGuardDisabled (virtual field)
	GetIsGuardDisabled() bool
	// GetIsGuardEnabled returns IsGuardEnabled (virtual field)
	GetIsGuardEnabled() bool
	// GetIsSetbackDisabled returns IsSetbackDisabled (virtual field)
	GetIsSetbackDisabled() bool
	// GetIsSetbackEnabled returns IsSetbackEnabled (virtual field)
	GetIsSetbackEnabled() bool
	// GetIsLevelTemperature returns IsLevelTemperature (virtual field)
	GetIsLevelTemperature() bool
	// GetIsLevelRaw returns IsLevelRaw (virtual field)
	GetIsLevelRaw() bool
	// IsHVACModeAndFlags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACModeAndFlags()
	// CreateBuilder creates a HVACModeAndFlagsBuilder
	CreateHVACModeAndFlagsBuilder() HVACModeAndFlagsBuilder
}

// _HVACModeAndFlags is the data-structure of this message
type _HVACModeAndFlags struct {
	AuxiliaryLevel bool
	Guard          bool
	Setback        bool
	Level          bool
	Mode           HVACModeAndFlagsMode
	// Reserved Fields
	reservedField0 *bool
}

var _ HVACModeAndFlags = (*_HVACModeAndFlags)(nil)

// NewHVACModeAndFlags factory function for _HVACModeAndFlags
func NewHVACModeAndFlags(auxiliaryLevel bool, guard bool, setback bool, level bool, mode HVACModeAndFlagsMode) *_HVACModeAndFlags {
	return &_HVACModeAndFlags{AuxiliaryLevel: auxiliaryLevel, Guard: guard, Setback: setback, Level: level, Mode: mode}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HVACModeAndFlagsBuilder is a builder for HVACModeAndFlags
type HVACModeAndFlagsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(auxiliaryLevel bool, guard bool, setback bool, level bool, mode HVACModeAndFlagsMode) HVACModeAndFlagsBuilder
	// WithAuxiliaryLevel adds AuxiliaryLevel (property field)
	WithAuxiliaryLevel(bool) HVACModeAndFlagsBuilder
	// WithGuard adds Guard (property field)
	WithGuard(bool) HVACModeAndFlagsBuilder
	// WithSetback adds Setback (property field)
	WithSetback(bool) HVACModeAndFlagsBuilder
	// WithLevel adds Level (property field)
	WithLevel(bool) HVACModeAndFlagsBuilder
	// WithMode adds Mode (property field)
	WithMode(HVACModeAndFlagsMode) HVACModeAndFlagsBuilder
	// Build builds the HVACModeAndFlags or returns an error if something is wrong
	Build() (HVACModeAndFlags, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HVACModeAndFlags
}

// NewHVACModeAndFlagsBuilder() creates a HVACModeAndFlagsBuilder
func NewHVACModeAndFlagsBuilder() HVACModeAndFlagsBuilder {
	return &_HVACModeAndFlagsBuilder{_HVACModeAndFlags: new(_HVACModeAndFlags)}
}

type _HVACModeAndFlagsBuilder struct {
	*_HVACModeAndFlags

	err *utils.MultiError
}

var _ (HVACModeAndFlagsBuilder) = (*_HVACModeAndFlagsBuilder)(nil)

func (b *_HVACModeAndFlagsBuilder) WithMandatoryFields(auxiliaryLevel bool, guard bool, setback bool, level bool, mode HVACModeAndFlagsMode) HVACModeAndFlagsBuilder {
	return b.WithAuxiliaryLevel(auxiliaryLevel).WithGuard(guard).WithSetback(setback).WithLevel(level).WithMode(mode)
}

func (b *_HVACModeAndFlagsBuilder) WithAuxiliaryLevel(auxiliaryLevel bool) HVACModeAndFlagsBuilder {
	b.AuxiliaryLevel = auxiliaryLevel
	return b
}

func (b *_HVACModeAndFlagsBuilder) WithGuard(guard bool) HVACModeAndFlagsBuilder {
	b.Guard = guard
	return b
}

func (b *_HVACModeAndFlagsBuilder) WithSetback(setback bool) HVACModeAndFlagsBuilder {
	b.Setback = setback
	return b
}

func (b *_HVACModeAndFlagsBuilder) WithLevel(level bool) HVACModeAndFlagsBuilder {
	b.Level = level
	return b
}

func (b *_HVACModeAndFlagsBuilder) WithMode(mode HVACModeAndFlagsMode) HVACModeAndFlagsBuilder {
	b.Mode = mode
	return b
}

func (b *_HVACModeAndFlagsBuilder) Build() (HVACModeAndFlags, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HVACModeAndFlags.deepCopy(), nil
}

func (b *_HVACModeAndFlagsBuilder) MustBuild() HVACModeAndFlags {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HVACModeAndFlagsBuilder) DeepCopy() any {
	_copy := b.CreateHVACModeAndFlagsBuilder().(*_HVACModeAndFlagsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHVACModeAndFlagsBuilder creates a HVACModeAndFlagsBuilder
func (b *_HVACModeAndFlags) CreateHVACModeAndFlagsBuilder() HVACModeAndFlagsBuilder {
	if b == nil {
		return NewHVACModeAndFlagsBuilder()
	}
	return &_HVACModeAndFlagsBuilder{_HVACModeAndFlags: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACModeAndFlags) GetAuxiliaryLevel() bool {
	return m.AuxiliaryLevel
}

func (m *_HVACModeAndFlags) GetGuard() bool {
	return m.Guard
}

func (m *_HVACModeAndFlags) GetSetback() bool {
	return m.Setback
}

func (m *_HVACModeAndFlags) GetLevel() bool {
	return m.Level
}

func (m *_HVACModeAndFlags) GetMode() HVACModeAndFlagsMode {
	return m.Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACModeAndFlags) GetIsAuxLevelUnused() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetAuxiliaryLevel()))
}

func (m *_HVACModeAndFlags) GetIsAuxLevelUsed() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetAuxiliaryLevel())
}

func (m *_HVACModeAndFlags) GetIsGuardDisabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetGuard()))
}

func (m *_HVACModeAndFlags) GetIsGuardEnabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetGuard())
}

func (m *_HVACModeAndFlags) GetIsSetbackDisabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetSetback()))
}

func (m *_HVACModeAndFlags) GetIsSetbackEnabled() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetSetback())
}

func (m *_HVACModeAndFlags) GetIsLevelTemperature() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetLevel()))
}

func (m *_HVACModeAndFlags) GetIsLevelRaw() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHVACModeAndFlags(structType any) HVACModeAndFlags {
	if casted, ok := structType.(HVACModeAndFlags); ok {
		return casted
	}
	if casted, ok := structType.(*HVACModeAndFlags); ok {
		return *casted
	}
	return nil
}

func (m *_HVACModeAndFlags) GetTypeName() string {
	return "HVACModeAndFlags"
}

func (m *_HVACModeAndFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (auxiliaryLevel)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (guard)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (setback)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (level)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (mode)
	lengthInBits += 3

	return lengthInBits
}

func (m *_HVACModeAndFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACModeAndFlagsParse(ctx context.Context, theBytes []byte) (HVACModeAndFlags, error) {
	return HVACModeAndFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACModeAndFlagsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACModeAndFlags, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACModeAndFlags, error) {
		return HVACModeAndFlagsParseWithBuffer(ctx, readBuffer)
	}
}

func HVACModeAndFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACModeAndFlags, error) {
	v, err := (&_HVACModeAndFlags{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACModeAndFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACModeAndFlags HVACModeAndFlags, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACModeAndFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACModeAndFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	auxiliaryLevel, err := ReadSimpleField(ctx, "auxiliaryLevel", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auxiliaryLevel' field"))
	}
	m.AuxiliaryLevel = auxiliaryLevel

	isAuxLevelUnused, err := ReadVirtualField[bool](ctx, "isAuxLevelUnused", (*bool)(nil), !(auxiliaryLevel))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAuxLevelUnused' field"))
	}
	_ = isAuxLevelUnused

	isAuxLevelUsed, err := ReadVirtualField[bool](ctx, "isAuxLevelUsed", (*bool)(nil), auxiliaryLevel)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAuxLevelUsed' field"))
	}
	_ = isAuxLevelUsed

	guard, err := ReadSimpleField(ctx, "guard", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'guard' field"))
	}
	m.Guard = guard

	isGuardDisabled, err := ReadVirtualField[bool](ctx, "isGuardDisabled", (*bool)(nil), !(guard))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isGuardDisabled' field"))
	}
	_ = isGuardDisabled

	isGuardEnabled, err := ReadVirtualField[bool](ctx, "isGuardEnabled", (*bool)(nil), guard)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isGuardEnabled' field"))
	}
	_ = isGuardEnabled

	setback, err := ReadSimpleField(ctx, "setback", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'setback' field"))
	}
	m.Setback = setback

	isSetbackDisabled, err := ReadVirtualField[bool](ctx, "isSetbackDisabled", (*bool)(nil), !(setback))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetbackDisabled' field"))
	}
	_ = isSetbackDisabled

	isSetbackEnabled, err := ReadVirtualField[bool](ctx, "isSetbackEnabled", (*bool)(nil), setback)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetbackEnabled' field"))
	}
	_ = isSetbackEnabled

	level, err := ReadSimpleField(ctx, "level", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	m.Level = level

	isLevelTemperature, err := ReadVirtualField[bool](ctx, "isLevelTemperature", (*bool)(nil), !(level))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isLevelTemperature' field"))
	}
	_ = isLevelTemperature

	isLevelRaw, err := ReadVirtualField[bool](ctx, "isLevelRaw", (*bool)(nil), level)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isLevelRaw' field"))
	}
	_ = isLevelRaw

	mode, err := ReadEnumField[HVACModeAndFlagsMode](ctx, "mode", "HVACModeAndFlagsMode", ReadEnum(HVACModeAndFlagsModeByValue, ReadUnsignedByte(readBuffer, uint8(3))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mode' field"))
	}
	m.Mode = mode

	if closeErr := readBuffer.CloseContext("HVACModeAndFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACModeAndFlags")
	}

	return m, nil
}

func (m *_HVACModeAndFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACModeAndFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACModeAndFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACModeAndFlags")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "auxiliaryLevel", m.GetAuxiliaryLevel(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'auxiliaryLevel' field")
	}
	// Virtual field
	isAuxLevelUnused := m.GetIsAuxLevelUnused()
	_ = isAuxLevelUnused
	if _isAuxLevelUnusedErr := writeBuffer.WriteVirtual(ctx, "isAuxLevelUnused", m.GetIsAuxLevelUnused()); _isAuxLevelUnusedErr != nil {
		return errors.Wrap(_isAuxLevelUnusedErr, "Error serializing 'isAuxLevelUnused' field")
	}
	// Virtual field
	isAuxLevelUsed := m.GetIsAuxLevelUsed()
	_ = isAuxLevelUsed
	if _isAuxLevelUsedErr := writeBuffer.WriteVirtual(ctx, "isAuxLevelUsed", m.GetIsAuxLevelUsed()); _isAuxLevelUsedErr != nil {
		return errors.Wrap(_isAuxLevelUsedErr, "Error serializing 'isAuxLevelUsed' field")
	}

	if err := WriteSimpleField[bool](ctx, "guard", m.GetGuard(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'guard' field")
	}
	// Virtual field
	isGuardDisabled := m.GetIsGuardDisabled()
	_ = isGuardDisabled
	if _isGuardDisabledErr := writeBuffer.WriteVirtual(ctx, "isGuardDisabled", m.GetIsGuardDisabled()); _isGuardDisabledErr != nil {
		return errors.Wrap(_isGuardDisabledErr, "Error serializing 'isGuardDisabled' field")
	}
	// Virtual field
	isGuardEnabled := m.GetIsGuardEnabled()
	_ = isGuardEnabled
	if _isGuardEnabledErr := writeBuffer.WriteVirtual(ctx, "isGuardEnabled", m.GetIsGuardEnabled()); _isGuardEnabledErr != nil {
		return errors.Wrap(_isGuardEnabledErr, "Error serializing 'isGuardEnabled' field")
	}

	if err := WriteSimpleField[bool](ctx, "setback", m.GetSetback(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'setback' field")
	}
	// Virtual field
	isSetbackDisabled := m.GetIsSetbackDisabled()
	_ = isSetbackDisabled
	if _isSetbackDisabledErr := writeBuffer.WriteVirtual(ctx, "isSetbackDisabled", m.GetIsSetbackDisabled()); _isSetbackDisabledErr != nil {
		return errors.Wrap(_isSetbackDisabledErr, "Error serializing 'isSetbackDisabled' field")
	}
	// Virtual field
	isSetbackEnabled := m.GetIsSetbackEnabled()
	_ = isSetbackEnabled
	if _isSetbackEnabledErr := writeBuffer.WriteVirtual(ctx, "isSetbackEnabled", m.GetIsSetbackEnabled()); _isSetbackEnabledErr != nil {
		return errors.Wrap(_isSetbackEnabledErr, "Error serializing 'isSetbackEnabled' field")
	}

	if err := WriteSimpleField[bool](ctx, "level", m.GetLevel(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'level' field")
	}
	// Virtual field
	isLevelTemperature := m.GetIsLevelTemperature()
	_ = isLevelTemperature
	if _isLevelTemperatureErr := writeBuffer.WriteVirtual(ctx, "isLevelTemperature", m.GetIsLevelTemperature()); _isLevelTemperatureErr != nil {
		return errors.Wrap(_isLevelTemperatureErr, "Error serializing 'isLevelTemperature' field")
	}
	// Virtual field
	isLevelRaw := m.GetIsLevelRaw()
	_ = isLevelRaw
	if _isLevelRawErr := writeBuffer.WriteVirtual(ctx, "isLevelRaw", m.GetIsLevelRaw()); _isLevelRawErr != nil {
		return errors.Wrap(_isLevelRawErr, "Error serializing 'isLevelRaw' field")
	}

	if err := WriteSimpleEnumField[HVACModeAndFlagsMode](ctx, "mode", "HVACModeAndFlagsMode", m.GetMode(), WriteEnum[HVACModeAndFlagsMode, uint8](HVACModeAndFlagsMode.GetValue, HVACModeAndFlagsMode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 3))); err != nil {
		return errors.Wrap(err, "Error serializing 'mode' field")
	}

	if popErr := writeBuffer.PopContext("HVACModeAndFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACModeAndFlags")
	}
	return nil
}

func (m *_HVACModeAndFlags) IsHVACModeAndFlags() {}

func (m *_HVACModeAndFlags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACModeAndFlags) deepCopy() *_HVACModeAndFlags {
	if m == nil {
		return nil
	}
	_HVACModeAndFlagsCopy := &_HVACModeAndFlags{
		m.AuxiliaryLevel,
		m.Guard,
		m.Setback,
		m.Level,
		m.Mode,
		m.reservedField0,
	}
	return _HVACModeAndFlagsCopy
}

func (m *_HVACModeAndFlags) String() string {
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
