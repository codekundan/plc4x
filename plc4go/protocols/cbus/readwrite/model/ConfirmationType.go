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

// ConfirmationType is an enum
type ConfirmationType byte

type IConfirmationType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ConfirmationType_CONFIRMATION_SUCCESSFUL                  ConfirmationType = 0x2E
	ConfirmationType_NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS ConfirmationType = 0x23
	ConfirmationType_NOT_TRANSMITTED_CORRUPTION               ConfirmationType = 0x24
	ConfirmationType_NOT_TRANSMITTED_SYNC_LOSS                ConfirmationType = 0x25
	ConfirmationType_NOT_TRANSMITTED_TOO_LONG                 ConfirmationType = 0x27
	ConfirmationType_CHECKSUM_FAILURE                         ConfirmationType = 0x21
)

var ConfirmationTypeValues []ConfirmationType

func init() {
	_ = errors.New
	ConfirmationTypeValues = []ConfirmationType{
		ConfirmationType_CONFIRMATION_SUCCESSFUL,
		ConfirmationType_NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS,
		ConfirmationType_NOT_TRANSMITTED_CORRUPTION,
		ConfirmationType_NOT_TRANSMITTED_SYNC_LOSS,
		ConfirmationType_NOT_TRANSMITTED_TOO_LONG,
		ConfirmationType_CHECKSUM_FAILURE,
	}
}

func ConfirmationTypeByValue(value byte) (enum ConfirmationType, ok bool) {
	switch value {
	case 0x21:
		return ConfirmationType_CHECKSUM_FAILURE, true
	case 0x23:
		return ConfirmationType_NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS, true
	case 0x24:
		return ConfirmationType_NOT_TRANSMITTED_CORRUPTION, true
	case 0x25:
		return ConfirmationType_NOT_TRANSMITTED_SYNC_LOSS, true
	case 0x27:
		return ConfirmationType_NOT_TRANSMITTED_TOO_LONG, true
	case 0x2E:
		return ConfirmationType_CONFIRMATION_SUCCESSFUL, true
	}
	return 0, false
}

func ConfirmationTypeByName(value string) (enum ConfirmationType, ok bool) {
	switch value {
	case "CHECKSUM_FAILURE":
		return ConfirmationType_CHECKSUM_FAILURE, true
	case "NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS":
		return ConfirmationType_NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS, true
	case "NOT_TRANSMITTED_CORRUPTION":
		return ConfirmationType_NOT_TRANSMITTED_CORRUPTION, true
	case "NOT_TRANSMITTED_SYNC_LOSS":
		return ConfirmationType_NOT_TRANSMITTED_SYNC_LOSS, true
	case "NOT_TRANSMITTED_TOO_LONG":
		return ConfirmationType_NOT_TRANSMITTED_TOO_LONG, true
	case "CONFIRMATION_SUCCESSFUL":
		return ConfirmationType_CONFIRMATION_SUCCESSFUL, true
	}
	return 0, false
}

func ConfirmationTypeKnows(value byte) bool {
	for _, typeValue := range ConfirmationTypeValues {
		if byte(typeValue) == value {
			return true
		}
	}
	return false
}

func CastConfirmationType(structType interface{}) ConfirmationType {
	castFunc := func(typ interface{}) ConfirmationType {
		if sConfirmationType, ok := typ.(ConfirmationType); ok {
			return sConfirmationType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ConfirmationType) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ConfirmationType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConfirmationTypeParse(ctx context.Context, theBytes []byte) (ConfirmationType, error) {
	return ConfirmationTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ConfirmationTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConfirmationType, error) {
	val, err := readBuffer.ReadByte("ConfirmationType")
	if err != nil {
		return 0, errors.Wrap(err, "error reading ConfirmationType")
	}
	if enum, ok := ConfirmationTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ConfirmationType(val), nil
	} else {
		return enum, nil
	}
}

func (e ConfirmationType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ConfirmationType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteByte("ConfirmationType", byte(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ConfirmationType) PLC4XEnumName() string {
	switch e {
	case ConfirmationType_CHECKSUM_FAILURE:
		return "CHECKSUM_FAILURE"
	case ConfirmationType_NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS:
		return "NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS"
	case ConfirmationType_NOT_TRANSMITTED_CORRUPTION:
		return "NOT_TRANSMITTED_CORRUPTION"
	case ConfirmationType_NOT_TRANSMITTED_SYNC_LOSS:
		return "NOT_TRANSMITTED_SYNC_LOSS"
	case ConfirmationType_NOT_TRANSMITTED_TOO_LONG:
		return "NOT_TRANSMITTED_TOO_LONG"
	case ConfirmationType_CONFIRMATION_SUCCESSFUL:
		return "CONFIRMATION_SUCCESSFUL"
	}
	return ""
}

func (e ConfirmationType) String() string {
	return e.PLC4XEnumName()
}
