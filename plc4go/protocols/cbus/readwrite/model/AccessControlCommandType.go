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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AccessControlCommandType is an enum
type AccessControlCommandType uint8

type IAccessControlCommandType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	AccessControlCommandType_CLOSE_ACCESS_POINT       AccessControlCommandType = 0x00
	AccessControlCommandType_LOCK_ACCESS_POINT        AccessControlCommandType = 0x01
	AccessControlCommandType_ACCESS_POINT_LEFT_OPEN   AccessControlCommandType = 0x02
	AccessControlCommandType_ACCESS_POINT_FORCED_OPEN AccessControlCommandType = 0x03
	AccessControlCommandType_ACCESS_POINT_CLOSED      AccessControlCommandType = 0x04
	AccessControlCommandType_REQUEST_TO_EXIT          AccessControlCommandType = 0x05
	AccessControlCommandType_VALID_ACCESS             AccessControlCommandType = 0x06
	AccessControlCommandType_INVALID_ACCESS           AccessControlCommandType = 0x07
)

var AccessControlCommandTypeValues []AccessControlCommandType

func init() {
	_ = errors.New
	AccessControlCommandTypeValues = []AccessControlCommandType{
		AccessControlCommandType_CLOSE_ACCESS_POINT,
		AccessControlCommandType_LOCK_ACCESS_POINT,
		AccessControlCommandType_ACCESS_POINT_LEFT_OPEN,
		AccessControlCommandType_ACCESS_POINT_FORCED_OPEN,
		AccessControlCommandType_ACCESS_POINT_CLOSED,
		AccessControlCommandType_REQUEST_TO_EXIT,
		AccessControlCommandType_VALID_ACCESS,
		AccessControlCommandType_INVALID_ACCESS,
	}
}

func (e AccessControlCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0
		}
	case 0x01:
		{ /* '0x01' */
			return 0
		}
	case 0x02:
		{ /* '0x02' */
			return 0
		}
	case 0x03:
		{ /* '0x03' */
			return 0
		}
	case 0x04:
		{ /* '0x04' */
			return 0
		}
	case 0x05:
		{ /* '0x05' */
			return 0
		}
	case 0x06:
		{ /* '0x06' */
			return 2
		}
	case 0x07:
		{ /* '0x07' */
			return 2
		}
	default:
		{
			return 0
		}
	}
}

func AccessControlCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (AccessControlCommandType, error) {
	for _, sizeValue := range AccessControlCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func AccessControlCommandTypeByValue(value uint8) (enum AccessControlCommandType, ok bool) {
	switch value {
	case 0x00:
		return AccessControlCommandType_CLOSE_ACCESS_POINT, true
	case 0x01:
		return AccessControlCommandType_LOCK_ACCESS_POINT, true
	case 0x02:
		return AccessControlCommandType_ACCESS_POINT_LEFT_OPEN, true
	case 0x03:
		return AccessControlCommandType_ACCESS_POINT_FORCED_OPEN, true
	case 0x04:
		return AccessControlCommandType_ACCESS_POINT_CLOSED, true
	case 0x05:
		return AccessControlCommandType_REQUEST_TO_EXIT, true
	case 0x06:
		return AccessControlCommandType_VALID_ACCESS, true
	case 0x07:
		return AccessControlCommandType_INVALID_ACCESS, true
	}
	return 0, false
}

func AccessControlCommandTypeByName(value string) (enum AccessControlCommandType, ok bool) {
	switch value {
	case "CLOSE_ACCESS_POINT":
		return AccessControlCommandType_CLOSE_ACCESS_POINT, true
	case "LOCK_ACCESS_POINT":
		return AccessControlCommandType_LOCK_ACCESS_POINT, true
	case "ACCESS_POINT_LEFT_OPEN":
		return AccessControlCommandType_ACCESS_POINT_LEFT_OPEN, true
	case "ACCESS_POINT_FORCED_OPEN":
		return AccessControlCommandType_ACCESS_POINT_FORCED_OPEN, true
	case "ACCESS_POINT_CLOSED":
		return AccessControlCommandType_ACCESS_POINT_CLOSED, true
	case "REQUEST_TO_EXIT":
		return AccessControlCommandType_REQUEST_TO_EXIT, true
	case "VALID_ACCESS":
		return AccessControlCommandType_VALID_ACCESS, true
	case "INVALID_ACCESS":
		return AccessControlCommandType_INVALID_ACCESS, true
	}
	return 0, false
}

func AccessControlCommandTypeKnows(value uint8) bool {
	for _, typeValue := range AccessControlCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAccessControlCommandType(structType interface{}) AccessControlCommandType {
	castFunc := func(typ interface{}) AccessControlCommandType {
		if sAccessControlCommandType, ok := typ.(AccessControlCommandType); ok {
			return sAccessControlCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessControlCommandType) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m AccessControlCommandType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessControlCommandTypeParse(ctx context.Context, theBytes []byte) (AccessControlCommandType, error) {
	return AccessControlCommandTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AccessControlCommandTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessControlCommandType, error) {
	val, err := readBuffer.ReadUint8("AccessControlCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessControlCommandType")
	}
	if enum, ok := AccessControlCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AccessControlCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e AccessControlCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AccessControlCommandType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AccessControlCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessControlCommandType) PLC4XEnumName() string {
	switch e {
	case AccessControlCommandType_CLOSE_ACCESS_POINT:
		return "CLOSE_ACCESS_POINT"
	case AccessControlCommandType_LOCK_ACCESS_POINT:
		return "LOCK_ACCESS_POINT"
	case AccessControlCommandType_ACCESS_POINT_LEFT_OPEN:
		return "ACCESS_POINT_LEFT_OPEN"
	case AccessControlCommandType_ACCESS_POINT_FORCED_OPEN:
		return "ACCESS_POINT_FORCED_OPEN"
	case AccessControlCommandType_ACCESS_POINT_CLOSED:
		return "ACCESS_POINT_CLOSED"
	case AccessControlCommandType_REQUEST_TO_EXIT:
		return "REQUEST_TO_EXIT"
	case AccessControlCommandType_VALID_ACCESS:
		return "VALID_ACCESS"
	case AccessControlCommandType_INVALID_ACCESS:
		return "INVALID_ACCESS"
	}
	return ""
}

func (e AccessControlCommandType) String() string {
	return e.PLC4XEnumName()
}
