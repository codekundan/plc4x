/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestSubscribeCOV is the data-structure of this message
type BACnetConfirmedServiceRequestSubscribeCOV struct {
	*BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier *BACnetContextTagUnsignedInteger
	MonitoredObjectIdentifier   *BACnetContextTagObjectIdentifier
	IssueConfirmed              *BACnetContextTagBoolean
	LifetimeInSeconds           *BACnetContextTagUnsignedInteger

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestSubscribeCOV is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOV
type IBACnetConfirmedServiceRequestSubscribeCOV interface {
	IBACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetIssueConfirmed returns IssueConfirmed (property field)
	GetIssueConfirmed() *BACnetContextTagBoolean
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() *BACnetContextTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestSubscribeCOV) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetIssueConfirmed() *BACnetContextTagBoolean {
	return m.IssueConfirmed
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetLifetimeInSeconds() *BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOV factory function for BACnetConfirmedServiceRequestSubscribeCOV
func NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier *BACnetContextTagUnsignedInteger, monitoredObjectIdentifier *BACnetContextTagObjectIdentifier, issueConfirmed *BACnetContextTagBoolean, lifetimeInSeconds *BACnetContextTagUnsignedInteger, serviceRequestLength uint16) *BACnetConfirmedServiceRequestSubscribeCOV {
	_result := &BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		MonitoredObjectIdentifier:     monitoredObjectIdentifier,
		IssueConfirmed:                issueConfirmed,
		LifetimeInSeconds:             lifetimeInSeconds,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestSubscribeCOV(structType interface{}) *BACnetConfirmedServiceRequestSubscribeCOV {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOV); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOV); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestSubscribeCOV(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestSubscribeCOV(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOV"
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits()

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Simple field (issueConfirmed)
	lengthInBits += m.IssueConfirmed.GetLengthInBits()

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestSubscribeCOV, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOV"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}
	subscriberProcessIdentifier := CastBACnetContextTagUnsignedInteger(_subscriberProcessIdentifier)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := CastBACnetContextTagObjectIdentifier(_monitoredObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (issueConfirmed)
	if pullErr := readBuffer.PullContext("issueConfirmed"); pullErr != nil {
		return nil, pullErr
	}
	_issueConfirmed, _issueConfirmedErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _issueConfirmedErr != nil {
		return nil, errors.Wrap(_issueConfirmedErr, "Error parsing 'issueConfirmed' field")
	}
	issueConfirmed := CastBACnetContextTagBoolean(_issueConfirmed)
	if closeErr := readBuffer.CloseContext("issueConfirmed"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lifetimeInSeconds)
	if pullErr := readBuffer.PullContext("lifetimeInSeconds"); pullErr != nil {
		return nil, pullErr
	}
	_lifetimeInSeconds, _lifetimeInSecondsErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _lifetimeInSecondsErr != nil {
		return nil, errors.Wrap(_lifetimeInSecondsErr, "Error parsing 'lifetimeInSeconds' field")
	}
	lifetimeInSeconds := CastBACnetContextTagUnsignedInteger(_lifetimeInSeconds)
	if closeErr := readBuffer.CloseContext("lifetimeInSeconds"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOV"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:   CastBACnetContextTagUnsignedInteger(subscriberProcessIdentifier),
		MonitoredObjectIdentifier:     CastBACnetContextTagObjectIdentifier(monitoredObjectIdentifier),
		IssueConfirmed:                CastBACnetContextTagBoolean(issueConfirmed),
		LifetimeInSeconds:             CastBACnetContextTagUnsignedInteger(lifetimeInSeconds),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOV"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return pushErr
		}
		_subscriberProcessIdentifierErr := m.SubscriberProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return popErr
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return pushErr
		}
		_monitoredObjectIdentifierErr := m.MonitoredObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return popErr
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Simple Field (issueConfirmed)
		if pushErr := writeBuffer.PushContext("issueConfirmed"); pushErr != nil {
			return pushErr
		}
		_issueConfirmedErr := m.IssueConfirmed.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("issueConfirmed"); popErr != nil {
			return popErr
		}
		if _issueConfirmedErr != nil {
			return errors.Wrap(_issueConfirmedErr, "Error serializing 'issueConfirmed' field")
		}

		// Simple Field (lifetimeInSeconds)
		if pushErr := writeBuffer.PushContext("lifetimeInSeconds"); pushErr != nil {
			return pushErr
		}
		_lifetimeInSecondsErr := m.LifetimeInSeconds.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lifetimeInSeconds"); popErr != nil {
			return popErr
		}
		if _lifetimeInSecondsErr != nil {
			return errors.Wrap(_lifetimeInSecondsErr, "Error serializing 'lifetimeInSeconds' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOV"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
