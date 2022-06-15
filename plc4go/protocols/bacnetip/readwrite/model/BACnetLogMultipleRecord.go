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

// BACnetLogMultipleRecord is the corresponding interface of BACnetLogMultipleRecord
type BACnetLogMultipleRecord interface {
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetLogData returns LogData (property field)
	GetLogData() BACnetLogData
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetLogMultipleRecord is the data-structure of this message
type _BACnetLogMultipleRecord struct {
	Timestamp BACnetDateTimeEnclosed
	LogData   BACnetLogData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogMultipleRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetLogMultipleRecord) GetLogData() BACnetLogData {
	return m.LogData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogMultipleRecord factory function for _BACnetLogMultipleRecord
func NewBACnetLogMultipleRecord(timestamp BACnetDateTimeEnclosed, logData BACnetLogData) *_BACnetLogMultipleRecord {
	return &_BACnetLogMultipleRecord{Timestamp: timestamp, LogData: logData}
}

// Deprecated: use the interface for direct cast
func CastBACnetLogMultipleRecord(structType interface{}) BACnetLogMultipleRecord {
	if casted, ok := structType.(BACnetLogMultipleRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogMultipleRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogMultipleRecord) GetTypeName() string {
	return "BACnetLogMultipleRecord"
}

func (m *_BACnetLogMultipleRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogMultipleRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (logData)
	lengthInBits += m.LogData.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogMultipleRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogMultipleRecordParse(readBuffer utils.ReadBuffer) (BACnetLogMultipleRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogMultipleRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogMultipleRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
	}
	_timestamp, _timestampErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(0)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := _timestamp.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamp")
	}

	// Simple Field (logData)
	if pullErr := readBuffer.PullContext("logData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logData")
	}
	_logData, _logDataErr := BACnetLogDataParse(readBuffer, uint8(uint8(1)))
	if _logDataErr != nil {
		return nil, errors.Wrap(_logDataErr, "Error parsing 'logData' field")
	}
	logData := _logData.(BACnetLogData)
	if closeErr := readBuffer.CloseContext("logData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logData")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogMultipleRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogMultipleRecord")
	}

	// Create the instance
	return NewBACnetLogMultipleRecord(timestamp, logData), nil
}

func (m *_BACnetLogMultipleRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLogMultipleRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogMultipleRecord")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamp")
	}
	_timestampErr := writeBuffer.WriteSerializable(m.GetTimestamp())
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamp")
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (logData)
	if pushErr := writeBuffer.PushContext("logData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for logData")
	}
	_logDataErr := writeBuffer.WriteSerializable(m.GetLogData())
	if popErr := writeBuffer.PopContext("logData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for logData")
	}
	if _logDataErr != nil {
		return errors.Wrap(_logDataErr, "Error serializing 'logData' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogMultipleRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogMultipleRecord")
	}
	return nil
}

func (m *_BACnetLogMultipleRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
