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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsMultiRequestItem is the corresponding interface of AdsMultiRequestItem
type AdsMultiRequestItem interface {
	// GetIndexGroup returns IndexGroup (discriminator field)
	GetIndexGroup() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _AdsMultiRequestItem is the data-structure of this message
type _AdsMultiRequestItem struct {
	_AdsMultiRequestItemChildRequirements
}

type _AdsMultiRequestItemChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetIndexGroup() uint32
}

type AdsMultiRequestItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child AdsMultiRequestItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type AdsMultiRequestItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent AdsMultiRequestItem)
	GetParent() *AdsMultiRequestItem

	GetTypeName() string
	AdsMultiRequestItem
}

// NewAdsMultiRequestItem factory function for _AdsMultiRequestItem
func NewAdsMultiRequestItem() *_AdsMultiRequestItem {
	return &_AdsMultiRequestItem{}
}

// Deprecated: use the interface for direct cast
func CastAdsMultiRequestItem(structType interface{}) AdsMultiRequestItem {
	if casted, ok := structType.(AdsMultiRequestItem); ok {
		return casted
	}
	if casted, ok := structType.(*AdsMultiRequestItem); ok {
		return *casted
	}
	return nil
}

func (m *_AdsMultiRequestItem) GetTypeName() string {
	return "AdsMultiRequestItem"
}

func (m *_AdsMultiRequestItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_AdsMultiRequestItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsMultiRequestItemParse(readBuffer utils.ReadBuffer, indexGroup uint32) (AdsMultiRequestItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsMultiRequestItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AdsMultiRequestItemChildSerializeRequirement interface {
		AdsMultiRequestItem
		InitializeParent(AdsMultiRequestItem)
		GetParent() AdsMultiRequestItem
	}
	var _childTemp interface{}
	var _child AdsMultiRequestItemChildSerializeRequirement
	var typeSwitchError error
	switch {
	case indexGroup == uint32(61568): // AdsMultiRequestItemRead
		_childTemp, typeSwitchError = AdsMultiRequestItemReadParse(readBuffer, indexGroup)
		_child = _childTemp.(AdsMultiRequestItemChildSerializeRequirement)
	case indexGroup == uint32(61569): // AdsMultiRequestItemWrite
		_childTemp, typeSwitchError = AdsMultiRequestItemWriteParse(readBuffer, indexGroup)
		_child = _childTemp.(AdsMultiRequestItemChildSerializeRequirement)
	case indexGroup == uint32(61570): // AdsMultiRequestItemReadWrite
		_childTemp, typeSwitchError = AdsMultiRequestItemReadWriteParse(readBuffer, indexGroup)
		_child = _childTemp.(AdsMultiRequestItemChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsMultiRequestItem")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (m *_AdsMultiRequestItem) Serialize(writeBuffer utils.WriteBuffer) error {
	panic("Required method Serialize not implemented")
}

func (pm *_AdsMultiRequestItem) SerializeParent(writeBuffer utils.WriteBuffer, child AdsMultiRequestItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsMultiRequestItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsMultiRequestItem")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsMultiRequestItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsMultiRequestItem")
	}
	return nil
}

func (m *_AdsMultiRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
