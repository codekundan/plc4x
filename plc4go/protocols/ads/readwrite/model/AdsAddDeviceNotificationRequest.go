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

// AdsAddDeviceNotificationRequest is the corresponding interface of AdsAddDeviceNotificationRequest
type AdsAddDeviceNotificationRequest interface {
	AdsData
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetLength returns Length (property field)
	GetLength() uint32
	// GetTransmissionMode returns TransmissionMode (property field)
	GetTransmissionMode() uint32
	// GetMaxDelay returns MaxDelay (property field)
	GetMaxDelay() uint32
	// GetCycleTime returns CycleTime (property field)
	GetCycleTime() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _AdsAddDeviceNotificationRequest is the data-structure of this message
type _AdsAddDeviceNotificationRequest struct {
	*_AdsData
	IndexGroup       uint32
	IndexOffset      uint32
	Length           uint32
	TransmissionMode uint32
	MaxDelay         uint32
	CycleTime        uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetCommandId() CommandId {
	return CommandId_ADS_ADD_DEVICE_NOTIFICATION
}

func (m *_AdsAddDeviceNotificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsAddDeviceNotificationRequest) InitializeParent(parent AdsData) {}

func (m *_AdsAddDeviceNotificationRequest) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsAddDeviceNotificationRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsAddDeviceNotificationRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsAddDeviceNotificationRequest) GetLength() uint32 {
	return m.Length
}

func (m *_AdsAddDeviceNotificationRequest) GetTransmissionMode() uint32 {
	return m.TransmissionMode
}

func (m *_AdsAddDeviceNotificationRequest) GetMaxDelay() uint32 {
	return m.MaxDelay
}

func (m *_AdsAddDeviceNotificationRequest) GetCycleTime() uint32 {
	return m.CycleTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsAddDeviceNotificationRequest factory function for _AdsAddDeviceNotificationRequest
func NewAdsAddDeviceNotificationRequest(indexGroup uint32, indexOffset uint32, length uint32, transmissionMode uint32, maxDelay uint32, cycleTime uint32) *_AdsAddDeviceNotificationRequest {
	_result := &_AdsAddDeviceNotificationRequest{
		IndexGroup:       indexGroup,
		IndexOffset:      indexOffset,
		Length:           length,
		TransmissionMode: transmissionMode,
		MaxDelay:         maxDelay,
		CycleTime:        cycleTime,
		_AdsData:         NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsAddDeviceNotificationRequest(structType interface{}) AdsAddDeviceNotificationRequest {
	if casted, ok := structType.(AdsAddDeviceNotificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsAddDeviceNotificationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsAddDeviceNotificationRequest) GetTypeName() string {
	return "AdsAddDeviceNotificationRequest"
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	// Simple field (transmissionMode)
	lengthInBits += 32

	// Simple field (maxDelay)
	lengthInBits += 32

	// Simple field (cycleTime)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 64

	// Reserved Field (reserved)
	lengthInBits += 64

	return lengthInBits
}

func (m *_AdsAddDeviceNotificationRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsAddDeviceNotificationRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsAddDeviceNotificationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsAddDeviceNotificationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsAddDeviceNotificationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (indexGroup)
	_indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field")
	}
	indexGroup := _indexGroup

	// Simple Field (indexOffset)
	_indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field")
	}
	indexOffset := _indexOffset

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadUint32("length", 32)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	length := _length

	// Simple Field (transmissionMode)
	_transmissionMode, _transmissionModeErr := readBuffer.ReadUint32("transmissionMode", 32)
	if _transmissionModeErr != nil {
		return nil, errors.Wrap(_transmissionModeErr, "Error parsing 'transmissionMode' field")
	}
	transmissionMode := _transmissionMode

	// Simple Field (maxDelay)
	_maxDelay, _maxDelayErr := readBuffer.ReadUint32("maxDelay", 32)
	if _maxDelayErr != nil {
		return nil, errors.Wrap(_maxDelayErr, "Error parsing 'maxDelay' field")
	}
	maxDelay := _maxDelay

	// Simple Field (cycleTime)
	_cycleTime, _cycleTimeErr := readBuffer.ReadUint32("cycleTime", 32)
	if _cycleTimeErr != nil {
		return nil, errors.Wrap(_cycleTimeErr, "Error parsing 'cycleTime' field")
	}
	cycleTime := _cycleTime

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint64("reserved", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint64(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint64(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint64("reserved", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint64(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint64(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("AdsAddDeviceNotificationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsAddDeviceNotificationRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsAddDeviceNotificationRequest{
		IndexGroup:       indexGroup,
		IndexOffset:      indexOffset,
		Length:           length,
		TransmissionMode: transmissionMode,
		MaxDelay:         maxDelay,
		CycleTime:        cycleTime,
		_AdsData:         &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsAddDeviceNotificationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsAddDeviceNotificationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsAddDeviceNotificationRequest")
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.GetIndexGroup())
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.GetIndexOffset())
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (length)
		length := uint32(m.GetLength())
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Simple Field (transmissionMode)
		transmissionMode := uint32(m.GetTransmissionMode())
		_transmissionModeErr := writeBuffer.WriteUint32("transmissionMode", 32, (transmissionMode))
		if _transmissionModeErr != nil {
			return errors.Wrap(_transmissionModeErr, "Error serializing 'transmissionMode' field")
		}

		// Simple Field (maxDelay)
		maxDelay := uint32(m.GetMaxDelay())
		_maxDelayErr := writeBuffer.WriteUint32("maxDelay", 32, (maxDelay))
		if _maxDelayErr != nil {
			return errors.Wrap(_maxDelayErr, "Error serializing 'maxDelay' field")
		}

		// Simple Field (cycleTime)
		cycleTime := uint32(m.GetCycleTime())
		_cycleTimeErr := writeBuffer.WriteUint32("cycleTime", 32, (cycleTime))
		if _cycleTimeErr != nil {
			return errors.Wrap(_cycleTimeErr, "Error serializing 'cycleTime' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint64("reserved", 64, uint64(0x0000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint64("reserved", 64, uint64(0x0000))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsAddDeviceNotificationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsAddDeviceNotificationRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsAddDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
