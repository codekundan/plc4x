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

// SearchResponse is the corresponding interface of SearchResponse
type SearchResponse interface {
	KnxNetIpMessage
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
	// GetDibDeviceInfo returns DibDeviceInfo (property field)
	GetDibDeviceInfo() DIBDeviceInfo
	// GetDibSuppSvcFamilies returns DibSuppSvcFamilies (property field)
	GetDibSuppSvcFamilies() DIBSuppSvcFamilies
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _SearchResponse is the data-structure of this message
type _SearchResponse struct {
	*_KnxNetIpMessage
	HpaiControlEndpoint HPAIControlEndpoint
	DibDeviceInfo       DIBDeviceInfo
	DibSuppSvcFamilies  DIBSuppSvcFamilies
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SearchResponse) GetMsgType() uint16 {
	return 0x0202
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SearchResponse) InitializeParent(parent KnxNetIpMessage) {}

func (m *_SearchResponse) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SearchResponse) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

func (m *_SearchResponse) GetDibDeviceInfo() DIBDeviceInfo {
	return m.DibDeviceInfo
}

func (m *_SearchResponse) GetDibSuppSvcFamilies() DIBSuppSvcFamilies {
	return m.DibSuppSvcFamilies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSearchResponse factory function for _SearchResponse
func NewSearchResponse(hpaiControlEndpoint HPAIControlEndpoint, dibDeviceInfo DIBDeviceInfo, dibSuppSvcFamilies DIBSuppSvcFamilies) *_SearchResponse {
	_result := &_SearchResponse{
		HpaiControlEndpoint: hpaiControlEndpoint,
		DibDeviceInfo:       dibDeviceInfo,
		DibSuppSvcFamilies:  dibSuppSvcFamilies,
		_KnxNetIpMessage:    NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSearchResponse(structType interface{}) SearchResponse {
	if casted, ok := structType.(SearchResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SearchResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SearchResponse) GetTypeName() string {
	return "SearchResponse"
}

func (m *_SearchResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SearchResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits()

	// Simple field (dibDeviceInfo)
	lengthInBits += m.DibDeviceInfo.GetLengthInBits()

	// Simple field (dibSuppSvcFamilies)
	lengthInBits += m.DibSuppSvcFamilies.GetLengthInBits()

	return lengthInBits
}

func (m *_SearchResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SearchResponseParse(readBuffer utils.ReadBuffer) (SearchResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SearchResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SearchResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (hpaiControlEndpoint)
	if pullErr := readBuffer.PullContext("hpaiControlEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hpaiControlEndpoint")
	}
	_hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(readBuffer)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field")
	}
	hpaiControlEndpoint := _hpaiControlEndpoint.(HPAIControlEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiControlEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hpaiControlEndpoint")
	}

	// Simple Field (dibDeviceInfo)
	if pullErr := readBuffer.PullContext("dibDeviceInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dibDeviceInfo")
	}
	_dibDeviceInfo, _dibDeviceInfoErr := DIBDeviceInfoParse(readBuffer)
	if _dibDeviceInfoErr != nil {
		return nil, errors.Wrap(_dibDeviceInfoErr, "Error parsing 'dibDeviceInfo' field")
	}
	dibDeviceInfo := _dibDeviceInfo.(DIBDeviceInfo)
	if closeErr := readBuffer.CloseContext("dibDeviceInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dibDeviceInfo")
	}

	// Simple Field (dibSuppSvcFamilies)
	if pullErr := readBuffer.PullContext("dibSuppSvcFamilies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dibSuppSvcFamilies")
	}
	_dibSuppSvcFamilies, _dibSuppSvcFamiliesErr := DIBSuppSvcFamiliesParse(readBuffer)
	if _dibSuppSvcFamiliesErr != nil {
		return nil, errors.Wrap(_dibSuppSvcFamiliesErr, "Error parsing 'dibSuppSvcFamilies' field")
	}
	dibSuppSvcFamilies := _dibSuppSvcFamilies.(DIBSuppSvcFamilies)
	if closeErr := readBuffer.CloseContext("dibSuppSvcFamilies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dibSuppSvcFamilies")
	}

	if closeErr := readBuffer.CloseContext("SearchResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SearchResponse")
	}

	// Create a partially initialized instance
	_child := &_SearchResponse{
		HpaiControlEndpoint: hpaiControlEndpoint,
		DibDeviceInfo:       dibDeviceInfo,
		DibSuppSvcFamilies:  dibSuppSvcFamilies,
		_KnxNetIpMessage:    &_KnxNetIpMessage{},
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_SearchResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SearchResponse")
		}

		// Simple Field (hpaiControlEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiControlEndpoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hpaiControlEndpoint")
		}
		_hpaiControlEndpointErr := writeBuffer.WriteSerializable(m.GetHpaiControlEndpoint())
		if popErr := writeBuffer.PopContext("hpaiControlEndpoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hpaiControlEndpoint")
		}
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		// Simple Field (dibDeviceInfo)
		if pushErr := writeBuffer.PushContext("dibDeviceInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dibDeviceInfo")
		}
		_dibDeviceInfoErr := writeBuffer.WriteSerializable(m.GetDibDeviceInfo())
		if popErr := writeBuffer.PopContext("dibDeviceInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dibDeviceInfo")
		}
		if _dibDeviceInfoErr != nil {
			return errors.Wrap(_dibDeviceInfoErr, "Error serializing 'dibDeviceInfo' field")
		}

		// Simple Field (dibSuppSvcFamilies)
		if pushErr := writeBuffer.PushContext("dibSuppSvcFamilies"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dibSuppSvcFamilies")
		}
		_dibSuppSvcFamiliesErr := writeBuffer.WriteSerializable(m.GetDibSuppSvcFamilies())
		if popErr := writeBuffer.PopContext("dibSuppSvcFamilies"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dibSuppSvcFamilies")
		}
		if _dibSuppSvcFamiliesErr != nil {
			return errors.Wrap(_dibSuppSvcFamiliesErr, "Error serializing 'dibSuppSvcFamilies' field")
		}

		if popErr := writeBuffer.PopContext("SearchResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SearchResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SearchResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
