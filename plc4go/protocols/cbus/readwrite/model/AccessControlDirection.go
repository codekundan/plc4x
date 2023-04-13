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

// AccessControlDirection is an enum
type AccessControlDirection uint8

type IAccessControlDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	AccessControlDirection_NOT_USED AccessControlDirection = 0x00
	AccessControlDirection_IN       AccessControlDirection = 0x01
	AccessControlDirection_OUT      AccessControlDirection = 0x02
)

var AccessControlDirectionValues []AccessControlDirection

func init() {
	_ = errors.New
	AccessControlDirectionValues = []AccessControlDirection{
		AccessControlDirection_NOT_USED,
		AccessControlDirection_IN,
		AccessControlDirection_OUT,
	}
}

func AccessControlDirectionByValue(value uint8) (enum AccessControlDirection, ok bool) {
	switch value {
	case 0x00:
		return AccessControlDirection_NOT_USED, true
	case 0x01:
		return AccessControlDirection_IN, true
	case 0x02:
		return AccessControlDirection_OUT, true
	}
	return 0, false
}

func AccessControlDirectionByName(value string) (enum AccessControlDirection, ok bool) {
	switch value {
	case "NOT_USED":
		return AccessControlDirection_NOT_USED, true
	case "IN":
		return AccessControlDirection_IN, true
	case "OUT":
		return AccessControlDirection_OUT, true
	}
	return 0, false
}

func AccessControlDirectionKnows(value uint8) bool {
	for _, typeValue := range AccessControlDirectionValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAccessControlDirection(structType interface{}) AccessControlDirection {
	castFunc := func(typ interface{}) AccessControlDirection {
		if sAccessControlDirection, ok := typ.(AccessControlDirection); ok {
			return sAccessControlDirection
		}
		return 0
	}
	return castFunc(structType)
}

func (m AccessControlDirection) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m AccessControlDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessControlDirectionParse(ctx context.Context, theBytes []byte) (AccessControlDirection, error) {
	return AccessControlDirectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AccessControlDirectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessControlDirection, error) {
	val, err := readBuffer.ReadUint8("AccessControlDirection", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AccessControlDirection")
	}
	if enum, ok := AccessControlDirectionByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AccessControlDirection(val), nil
	} else {
		return enum, nil
	}
}

func (e AccessControlDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AccessControlDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AccessControlDirection", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AccessControlDirection) PLC4XEnumName() string {
	switch e {
	case AccessControlDirection_NOT_USED:
		return "NOT_USED"
	case AccessControlDirection_IN:
		return "IN"
	case AccessControlDirection_OUT:
		return "OUT"
	}
	return ""
}

func (e AccessControlDirection) String() string {
	return e.PLC4XEnumName()
}
