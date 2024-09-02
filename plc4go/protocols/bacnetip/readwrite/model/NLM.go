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

// NLM is the corresponding interface of NLM
type NLM interface {
	NLMContract
	NLMRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// NLMContract provides a set of functions which can be overwritten by a sub struct
type NLMContract interface {
	// GetIsVendorProprietaryMessage returns IsVendorProprietaryMessage (virtual field)
	GetIsVendorProprietaryMessage() bool
	// GetApduLength() returns a parser argument
	GetApduLength() uint16
}

// NLMRequirements provides a set of functions which need to be implemented by a sub struct
type NLMRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsVendorProprietaryMessage returns IsVendorProprietaryMessage (discriminator field)
	GetIsVendorProprietaryMessage() bool
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
}

// NLMExactly can be used when we want exactly this type and not a type which fulfills NLM.
// This is useful for switch cases.
type NLMExactly interface {
	NLM
	isNLM() bool
}

// _NLM is the data-structure of this message
type _NLM struct {
	_SubType NLM

	// Arguments.
	ApduLength uint16
}

var _ NLMContract = (*_NLM)(nil)

type NLMChild interface {
	utils.Serializable

	GetParent() *NLM

	GetTypeName() string
	NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_NLM) GetIsVendorProprietaryMessage() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMessageType()) >= (128)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLM factory function for _NLM
func NewNLM(apduLength uint16) *_NLM {
	return &_NLM{ApduLength: apduLength}
}

// Deprecated: use the interface for direct cast
func CastNLM(structType any) NLM {
	if casted, ok := structType.(NLM); ok {
		return casted
	}
	if casted, ok := structType.(*NLM); ok {
		return *casted
	}
	return nil
}

func (m *_NLM) GetTypeName() string {
	return "NLM"
}

func (m *_NLM) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NLM) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func NLMParse[T NLM](ctx context.Context, theBytes []byte, apduLength uint16) (T, error) {
	return NLMParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMParseWithBufferProducer[T NLM](apduLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := NLMParseWithBuffer[T](ctx, readBuffer, apduLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func NLMParseWithBuffer[T NLM](ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (T, error) {
	v, err := (&_NLM{ApduLength: apduLength}).parse(ctx, readBuffer, apduLength)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_NLM) parse(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (__nLM NLM, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLM"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLM")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageType, err := ReadDiscriminatorField[uint8](ctx, "messageType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	isVendorProprietaryMessage, err := ReadVirtualField[bool](ctx, "isVendorProprietaryMessage", (*bool)(nil), bool((messageType) >= (128)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isVendorProprietaryMessage' field"))
	}
	_ = isVendorProprietaryMessage

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child NLM
	switch {
	case messageType == 0x00: // NLMWhoIsRouterToNetwork
		if _child, err = (&_NLMWhoIsRouterToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMWhoIsRouterToNetwork for type-switch of NLM")
		}
	case messageType == 0x01: // NLMIAmRouterToNetwork
		if _child, err = (&_NLMIAmRouterToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMIAmRouterToNetwork for type-switch of NLM")
		}
	case messageType == 0x02: // NLMICouldBeRouterToNetwork
		if _child, err = (&_NLMICouldBeRouterToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMICouldBeRouterToNetwork for type-switch of NLM")
		}
	case messageType == 0x03: // NLMRejectMessageToNetwork
		if _child, err = (&_NLMRejectMessageToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRejectMessageToNetwork for type-switch of NLM")
		}
	case messageType == 0x04: // NLMRouterBusyToNetwork
		if _child, err = (&_NLMRouterBusyToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRouterBusyToNetwork for type-switch of NLM")
		}
	case messageType == 0x05: // NLMRouterAvailableToNetwork
		if _child, err = (&_NLMRouterAvailableToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRouterAvailableToNetwork for type-switch of NLM")
		}
	case messageType == 0x06: // NLMInitializeRoutingTable
		if _child, err = (&_NLMInitializeRoutingTable{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMInitializeRoutingTable for type-switch of NLM")
		}
	case messageType == 0x07: // NLMInitializeRoutingTableAck
		if _child, err = (&_NLMInitializeRoutingTableAck{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMInitializeRoutingTableAck for type-switch of NLM")
		}
	case messageType == 0x08: // NLMEstablishConnectionToNetwork
		if _child, err = (&_NLMEstablishConnectionToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMEstablishConnectionToNetwork for type-switch of NLM")
		}
	case messageType == 0x09: // NLMDisconnectConnectionToNetwork
		if _child, err = (&_NLMDisconnectConnectionToNetwork{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMDisconnectConnectionToNetwork for type-switch of NLM")
		}
	case messageType == 0x0A: // NLMChallengeRequest
		if _child, err = (&_NLMChallengeRequest{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMChallengeRequest for type-switch of NLM")
		}
	case messageType == 0x0B: // NLMSecurityPayload
		if _child, err = (&_NLMSecurityPayload{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMSecurityPayload for type-switch of NLM")
		}
	case messageType == 0x0C: // NLMSecurityResponse
		if _child, err = (&_NLMSecurityResponse{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMSecurityResponse for type-switch of NLM")
		}
	case messageType == 0x0D: // NLMRequestKeyUpdate
		if _child, err = (&_NLMRequestKeyUpdate{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRequestKeyUpdate for type-switch of NLM")
		}
	case messageType == 0x0E: // NLMUpdateKeyUpdate
		if _child, err = (&_NLMUpdateKeyUpdate{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMUpdateKeyUpdate for type-switch of NLM")
		}
	case messageType == 0x0F: // NLMUpdateKeyDistributionKey
		if _child, err = (&_NLMUpdateKeyDistributionKey{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMUpdateKeyDistributionKey for type-switch of NLM")
		}
	case messageType == 0x10: // NLMRequestMasterKey
		if _child, err = (&_NLMRequestMasterKey{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMRequestMasterKey for type-switch of NLM")
		}
	case messageType == 0x11: // NLMSetMasterKey
		if _child, err = (&_NLMSetMasterKey{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMSetMasterKey for type-switch of NLM")
		}
	case messageType == 0x12: // NLMWhatIsNetworkNumber
		if _child, err = (&_NLMWhatIsNetworkNumber{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMWhatIsNetworkNumber for type-switch of NLM")
		}
	case messageType == 0x13: // NLMNetworkNumberIs
		if _child, err = (&_NLMNetworkNumberIs{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMNetworkNumberIs for type-switch of NLM")
		}
	case 0 == 0 && isVendorProprietaryMessage == bool(false): // NLMReserved
		if _child, err = (&_NLMReserved{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMReserved for type-switch of NLM")
		}
	case 0 == 0: // NLMVendorProprietaryMessage
		if _child, err = (&_NLMVendorProprietaryMessage{}).parse(ctx, readBuffer, m, apduLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NLMVendorProprietaryMessage for type-switch of NLM")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageType=%v, isVendorProprietaryMessage=%v]", messageType, isVendorProprietaryMessage)
	}

	if closeErr := readBuffer.CloseContext("NLM"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLM")
	}

	return _child, nil
}

func (pm *_NLM) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NLM, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NLM"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLM")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}
	// Virtual field
	isVendorProprietaryMessage := m.GetIsVendorProprietaryMessage()
	_ = isVendorProprietaryMessage
	if _isVendorProprietaryMessageErr := writeBuffer.WriteVirtual(ctx, "isVendorProprietaryMessage", m.GetIsVendorProprietaryMessage()); _isVendorProprietaryMessageErr != nil {
		return errors.Wrap(_isVendorProprietaryMessageErr, "Error serializing 'isVendorProprietaryMessage' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NLM"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLM")
	}
	return nil
}

////
// Arguments Getter

func (m *_NLM) GetApduLength() uint16 {
	return m.ApduLength
}

//
////

func (m *_NLM) isNLM() bool {
	return true
}
