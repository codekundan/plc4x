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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// DeviceConfigurationRequestDataBlock is the corresponding interface of DeviceConfigurationRequestDataBlock
type DeviceConfigurationRequestDataBlock interface {
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _DeviceConfigurationRequestDataBlock is the data-structure of this message
type _DeviceConfigurationRequestDataBlock struct {
	CommunicationChannelId uint8
	SequenceCounter        uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationRequestDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DeviceConfigurationRequestDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceConfigurationRequestDataBlock factory function for _DeviceConfigurationRequestDataBlock
func NewDeviceConfigurationRequestDataBlock(communicationChannelId uint8, sequenceCounter uint8) *_DeviceConfigurationRequestDataBlock {
	return &_DeviceConfigurationRequestDataBlock{CommunicationChannelId: communicationChannelId, SequenceCounter: sequenceCounter}
}

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationRequestDataBlock(structType interface{}) DeviceConfigurationRequestDataBlock {
	if casted, ok := structType.(DeviceConfigurationRequestDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationRequestDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationRequestDataBlock) GetTypeName() string {
	return "DeviceConfigurationRequestDataBlock"
}

func (m *_DeviceConfigurationRequestDataBlock) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DeviceConfigurationRequestDataBlock) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8

	// Simple field (sequenceCounter)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_DeviceConfigurationRequestDataBlock) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceConfigurationRequestDataBlockParse(readBuffer utils.ReadBuffer) (DeviceConfigurationRequestDataBlock, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequestDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationRequestDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (sequenceCounter)
	_sequenceCounter, _sequenceCounterErr := readBuffer.ReadUint8("sequenceCounter", 8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field")
	}
	sequenceCounter := _sequenceCounter

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequestDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationRequestDataBlock")
	}

	// Create the instance
	return NewDeviceConfigurationRequestDataBlock(communicationChannelId, sequenceCounter), nil
}

func (m *_DeviceConfigurationRequestDataBlock) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("DeviceConfigurationRequestDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationRequestDataBlock")
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.GetCommunicationChannelId())
	_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter := uint8(m.GetSequenceCounter())
	_sequenceCounterErr := writeBuffer.WriteUint8("sequenceCounter", 8, (sequenceCounter))
	if _sequenceCounterErr != nil {
		return errors.Wrap(_sequenceCounterErr, "Error serializing 'sequenceCounter' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	if popErr := writeBuffer.PopContext("DeviceConfigurationRequestDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceConfigurationRequestDataBlock")
	}
	return nil
}

func (m *_DeviceConfigurationRequestDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
