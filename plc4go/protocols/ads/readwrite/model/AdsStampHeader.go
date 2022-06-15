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

// AdsStampHeader is the corresponding interface of AdsStampHeader
type AdsStampHeader interface {
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() uint64
	// GetSamples returns Samples (property field)
	GetSamples() uint32
	// GetAdsNotificationSamples returns AdsNotificationSamples (property field)
	GetAdsNotificationSamples() []AdsNotificationSample
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _AdsStampHeader is the data-structure of this message
type _AdsStampHeader struct {
	Timestamp              uint64
	Samples                uint32
	AdsNotificationSamples []AdsNotificationSample
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsStampHeader) GetTimestamp() uint64 {
	return m.Timestamp
}

func (m *_AdsStampHeader) GetSamples() uint32 {
	return m.Samples
}

func (m *_AdsStampHeader) GetAdsNotificationSamples() []AdsNotificationSample {
	return m.AdsNotificationSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsStampHeader factory function for _AdsStampHeader
func NewAdsStampHeader(timestamp uint64, samples uint32, adsNotificationSamples []AdsNotificationSample) *_AdsStampHeader {
	return &_AdsStampHeader{Timestamp: timestamp, Samples: samples, AdsNotificationSamples: adsNotificationSamples}
}

// Deprecated: use the interface for direct cast
func CastAdsStampHeader(structType interface{}) AdsStampHeader {
	if casted, ok := structType.(AdsStampHeader); ok {
		return casted
	}
	if casted, ok := structType.(*AdsStampHeader); ok {
		return *casted
	}
	return nil
}

func (m *_AdsStampHeader) GetTypeName() string {
	return "AdsStampHeader"
}

func (m *_AdsStampHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsStampHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (samples)
	lengthInBits += 32

	// Array field
	if len(m.AdsNotificationSamples) > 0 {
		for i, element := range m.AdsNotificationSamples {
			last := i == len(m.AdsNotificationSamples)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_AdsStampHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsStampHeaderParse(readBuffer utils.ReadBuffer) (AdsStampHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsStampHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsStampHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	_timestamp, _timestampErr := readBuffer.ReadUint64("timestamp", 64)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := _timestamp

	// Simple Field (samples)
	_samples, _samplesErr := readBuffer.ReadUint32("samples", 32)
	if _samplesErr != nil {
		return nil, errors.Wrap(_samplesErr, "Error parsing 'samples' field")
	}
	samples := _samples

	// Array field (adsNotificationSamples)
	if pullErr := readBuffer.PullContext("adsNotificationSamples", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for adsNotificationSamples")
	}
	// Count array
	adsNotificationSamples := make([]AdsNotificationSample, samples)
	{
		for curItem := uint16(0); curItem < uint16(samples); curItem++ {
			_item, _err := AdsNotificationSampleParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'adsNotificationSamples' field")
			}
			adsNotificationSamples[curItem] = _item.(AdsNotificationSample)
		}
	}
	if closeErr := readBuffer.CloseContext("adsNotificationSamples", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for adsNotificationSamples")
	}

	if closeErr := readBuffer.CloseContext("AdsStampHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsStampHeader")
	}

	// Create the instance
	return NewAdsStampHeader(timestamp, samples, adsNotificationSamples), nil
}

func (m *_AdsStampHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsStampHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsStampHeader")
	}

	// Simple Field (timestamp)
	timestamp := uint64(m.GetTimestamp())
	_timestampErr := writeBuffer.WriteUint64("timestamp", 64, (timestamp))
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (samples)
	samples := uint32(m.GetSamples())
	_samplesErr := writeBuffer.WriteUint32("samples", 32, (samples))
	if _samplesErr != nil {
		return errors.Wrap(_samplesErr, "Error serializing 'samples' field")
	}

	// Array Field (adsNotificationSamples)
	if m.GetAdsNotificationSamples() != nil {
		if pushErr := writeBuffer.PushContext("adsNotificationSamples", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for adsNotificationSamples")
		}
		for _, _element := range m.GetAdsNotificationSamples() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'adsNotificationSamples' field")
			}
		}
		if popErr := writeBuffer.PopContext("adsNotificationSamples", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for adsNotificationSamples")
		}
	}

	if popErr := writeBuffer.PopContext("AdsStampHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsStampHeader")
	}
	return nil
}

func (m *_AdsStampHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
