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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const KnxNetIpMessage_PROTOCOLVERSION uint8 = 0x10

// KnxNetIpMessage is the corresponding interface of KnxNetIpMessage
type KnxNetIpMessage interface {
	// GetMsgType returns MsgType (discriminator field)
	GetMsgType() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _KnxNetIpMessage is the data-structure of this message
type _KnxNetIpMessage struct {
	_KnxNetIpMessageChildRequirements
}

type _KnxNetIpMessageChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetMsgType() uint16
}

type KnxNetIpMessageParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child KnxNetIpMessage, serializeChildFunction func() error) error
	GetTypeName() string
}

type KnxNetIpMessageChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent KnxNetIpMessage)
	GetParent() *KnxNetIpMessage

	GetTypeName() string
	KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_KnxNetIpMessage) GetProtocolVersion() uint8 {
	return KnxNetIpMessage_PROTOCOLVERSION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpMessage factory function for _KnxNetIpMessage
func NewKnxNetIpMessage() *_KnxNetIpMessage {
	return &_KnxNetIpMessage{}
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpMessage(structType interface{}) KnxNetIpMessage {
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpMessage) GetTypeName() string {
	return "KnxNetIpMessage"
}

func (m *_KnxNetIpMessage) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8

	// Const Field (protocolVersion)
	lengthInBits += 8
	// Discriminator Field (msgType)
	lengthInBits += 16

	// Implicit Field (totalLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_KnxNetIpMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpMessageParse(readBuffer utils.ReadBuffer) (KnxNetIpMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (headerLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	headerLength, _headerLengthErr := readBuffer.ReadUint8("headerLength", 8)
	_ = headerLength
	if _headerLengthErr != nil {
		return nil, errors.Wrap(_headerLengthErr, "Error parsing 'headerLength' field")
	}

	// Const Field (protocolVersion)
	protocolVersion, _protocolVersionErr := readBuffer.ReadUint8("protocolVersion", 8)
	if _protocolVersionErr != nil {
		return nil, errors.Wrap(_protocolVersionErr, "Error parsing 'protocolVersion' field")
	}
	if protocolVersion != KnxNetIpMessage_PROTOCOLVERSION {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", KnxNetIpMessage_PROTOCOLVERSION) + " but got " + fmt.Sprintf("%d", protocolVersion))
	}

	// Discriminator Field (msgType) (Used as input to a switch field)
	msgType, _msgTypeErr := readBuffer.ReadUint16("msgType", 16)
	if _msgTypeErr != nil {
		return nil, errors.Wrap(_msgTypeErr, "Error parsing 'msgType' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	totalLength, _totalLengthErr := readBuffer.ReadUint16("totalLength", 16)
	_ = totalLength
	if _totalLengthErr != nil {
		return nil, errors.Wrap(_totalLengthErr, "Error parsing 'totalLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type KnxNetIpMessageChildSerializeRequirement interface {
		KnxNetIpMessage
		InitializeParent(KnxNetIpMessage)
		GetParent() KnxNetIpMessage
	}
	var _childTemp interface{}
	var _child KnxNetIpMessageChildSerializeRequirement
	var typeSwitchError error
	switch {
	case msgType == 0x0201: // SearchRequest
		_childTemp, typeSwitchError = SearchRequestParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0202: // SearchResponse
		_childTemp, typeSwitchError = SearchResponseParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0203: // DescriptionRequest
		_childTemp, typeSwitchError = DescriptionRequestParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0204: // DescriptionResponse
		_childTemp, typeSwitchError = DescriptionResponseParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0205: // ConnectionRequest
		_childTemp, typeSwitchError = ConnectionRequestParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0206: // ConnectionResponse
		_childTemp, typeSwitchError = ConnectionResponseParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0207: // ConnectionStateRequest
		_childTemp, typeSwitchError = ConnectionStateRequestParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0208: // ConnectionStateResponse
		_childTemp, typeSwitchError = ConnectionStateResponseParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0209: // DisconnectRequest
		_childTemp, typeSwitchError = DisconnectRequestParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x020A: // DisconnectResponse
		_childTemp, typeSwitchError = DisconnectResponseParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x020B: // UnknownMessage
		_childTemp, typeSwitchError = UnknownMessageParse(readBuffer, totalLength)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0310: // DeviceConfigurationRequest
		_childTemp, typeSwitchError = DeviceConfigurationRequestParse(readBuffer, totalLength)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0311: // DeviceConfigurationAck
		_childTemp, typeSwitchError = DeviceConfigurationAckParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0420: // TunnelingRequest
		_childTemp, typeSwitchError = TunnelingRequestParse(readBuffer, totalLength)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0421: // TunnelingResponse
		_childTemp, typeSwitchError = TunnelingResponseParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	case msgType == 0x0530: // RoutingIndication
		_childTemp, typeSwitchError = RoutingIndicationParse(readBuffer)
		_child = _childTemp.(KnxNetIpMessageChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("KnxNetIpMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpMessage")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (m *_KnxNetIpMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	panic("Required method Serialize not implemented")
}

func (pm *_KnxNetIpMessage) SerializeParent(writeBuffer utils.WriteBuffer, child KnxNetIpMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("KnxNetIpMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for KnxNetIpMessage")
	}

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(6))
	_headerLengthErr := writeBuffer.WriteUint8("headerLength", 8, (headerLength))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	// Const Field (protocolVersion)
	_protocolVersionErr := writeBuffer.WriteUint8("protocolVersion", 8, 0x10)
	if _protocolVersionErr != nil {
		return errors.Wrap(_protocolVersionErr, "Error serializing 'protocolVersion' field")
	}

	// Discriminator Field (msgType) (Used as input to a switch field)
	msgType := uint16(child.GetMsgType())
	_msgTypeErr := writeBuffer.WriteUint16("msgType", 16, (msgType))

	if _msgTypeErr != nil {
		return errors.Wrap(_msgTypeErr, "Error serializing 'msgType' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	totalLength := uint16(uint16(m.GetLengthInBytes()))
	_totalLengthErr := writeBuffer.WriteUint16("totalLength", 16, (totalLength))
	if _totalLengthErr != nil {
		return errors.Wrap(_totalLengthErr, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("KnxNetIpMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for KnxNetIpMessage")
	}
	return nil
}

func (m *_KnxNetIpMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
