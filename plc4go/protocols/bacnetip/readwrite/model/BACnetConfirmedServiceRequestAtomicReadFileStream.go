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

// BACnetConfirmedServiceRequestAtomicReadFileStream is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileStream
type BACnetConfirmedServiceRequestAtomicReadFileStream interface {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() BACnetApplicationTagSignedInteger
	// GetRequestOctetCount returns RequestOctetCount (property field)
	GetRequestOctetCount() BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConfirmedServiceRequestAtomicReadFileStream is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileStream struct {
	*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	FileStartPosition BACnetApplicationTagSignedInteger
	RequestOctetCount BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) InitializeParent(parent BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) {
	m.PeekedTagHeader = peekedTagHeader
	m.OpeningTag = openingTag
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetParent() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	return m._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetFileStartPosition() BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetRequestOctetCount() BACnetApplicationTagUnsignedInteger {
	return m.RequestOctetCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicReadFileStream factory function for _BACnetConfirmedServiceRequestAtomicReadFileStream
func NewBACnetConfirmedServiceRequestAtomicReadFileStream(fileStartPosition BACnetApplicationTagSignedInteger, requestOctetCount BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestAtomicReadFileStream {
	_result := &_BACnetConfirmedServiceRequestAtomicReadFileStream{
		FileStartPosition: fileStartPosition,
		RequestOctetCount: requestOctetCount,
		_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord: NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
	}
	_result._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFileStream(structType interface{}) BACnetConfirmedServiceRequestAtomicReadFileStream {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileStream); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileStream); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileStream"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits()

	// Simple field (requestOctetCount)
	lengthInBits += m.RequestOctetCount.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicReadFileStreamParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestAtomicReadFileStream, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileStream")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileStartPosition")
	}
	_fileStartPosition, _fileStartPositionErr := BACnetApplicationTagParse(readBuffer)
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field")
	}
	fileStartPosition := _fileStartPosition.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileStartPosition")
	}

	// Simple Field (requestOctetCount)
	if pullErr := readBuffer.PullContext("requestOctetCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestOctetCount")
	}
	_requestOctetCount, _requestOctetCountErr := BACnetApplicationTagParse(readBuffer)
	if _requestOctetCountErr != nil {
		return nil, errors.Wrap(_requestOctetCountErr, "Error parsing 'requestOctetCount' field")
	}
	requestOctetCount := _requestOctetCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("requestOctetCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestOctetCount")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileStream")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestAtomicReadFileStream{
		FileStartPosition: fileStartPosition,
		RequestOctetCount: requestOctetCount,
		_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord: &_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord{},
	}
	_child._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileStream")
		}

		// Simple Field (fileStartPosition)
		if pushErr := writeBuffer.PushContext("fileStartPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileStartPosition")
		}
		_fileStartPositionErr := writeBuffer.WriteSerializable(m.GetFileStartPosition())
		if popErr := writeBuffer.PopContext("fileStartPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileStartPosition")
		}
		if _fileStartPositionErr != nil {
			return errors.Wrap(_fileStartPositionErr, "Error serializing 'fileStartPosition' field")
		}

		// Simple Field (requestOctetCount)
		if pushErr := writeBuffer.PushContext("requestOctetCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestOctetCount")
		}
		_requestOctetCountErr := writeBuffer.WriteSerializable(m.GetRequestOctetCount())
		if popErr := writeBuffer.PopContext("requestOctetCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestOctetCount")
		}
		if _requestOctetCountErr != nil {
			return errors.Wrap(_requestOctetCountErr, "Error serializing 'requestOctetCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileStream"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileStream")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileStream) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
