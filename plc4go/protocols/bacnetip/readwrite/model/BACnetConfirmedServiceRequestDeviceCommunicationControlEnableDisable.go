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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable is an enum
type BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable uint8

type IBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE             BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable = 0
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE            BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable = 1
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable = 2
)

var BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableValues []BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable

func init() {
	_ = errors.New
	BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableValues = []BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable{
		BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE,
		BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE,
		BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION,
	}
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableByValue(value uint8) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
	switch value {
	case 0:
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE
	case 1:
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE
	case 2:
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION
	}
	return 0
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableByName(value string) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
	switch value {
	case "ENABLE":
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE
	case "DISABLE":
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE
	case "DISABLE_INITIATION":
		return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION
	}
	return 0
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable(structType interface{}) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable {
		if sBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, ok := typ.(BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable); ok {
			return sBACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisableByValue(val), nil
}

func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) name() string {
	switch e {
	case BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_ENABLE:
		return "ENABLE"
	case BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE:
		return "DISABLE"
	case BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable_DISABLE_INITIATION:
		return "DISABLE_INITIATION"
	}
	return ""
}

func (e BACnetConfirmedServiceRequestDeviceCommunicationControlEnableDisable) String() string {
	return e.name()
}
