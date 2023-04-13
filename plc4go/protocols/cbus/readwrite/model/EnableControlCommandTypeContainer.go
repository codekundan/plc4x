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

// EnableControlCommandTypeContainer is an enum
type EnableControlCommandTypeContainer uint8

type IEnableControlCommandTypeContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumBytes() uint8
	CommandType() EnableControlCommandType
}

const (
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable0_2Bytes  EnableControlCommandTypeContainer = 0x02
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable1_2Bytes  EnableControlCommandTypeContainer = 0x0A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable2_2Bytes  EnableControlCommandTypeContainer = 0x12
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable3_2Bytes  EnableControlCommandTypeContainer = 0x1A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable4_2Bytes  EnableControlCommandTypeContainer = 0x22
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable5_2Bytes  EnableControlCommandTypeContainer = 0x2A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable6_2Bytes  EnableControlCommandTypeContainer = 0x32
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable7_2Bytes  EnableControlCommandTypeContainer = 0x3A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable8_2Bytes  EnableControlCommandTypeContainer = 0x42
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable9_2Bytes  EnableControlCommandTypeContainer = 0x4A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable10_2Bytes EnableControlCommandTypeContainer = 0x52
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable11_2Bytes EnableControlCommandTypeContainer = 0x5A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable12_2Bytes EnableControlCommandTypeContainer = 0x62
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable13_2Bytes EnableControlCommandTypeContainer = 0x6A
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable14_2Bytes EnableControlCommandTypeContainer = 0x72
	EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable15_2Bytes EnableControlCommandTypeContainer = 0x7A
)

var EnableControlCommandTypeContainerValues []EnableControlCommandTypeContainer

func init() {
	_ = errors.New
	EnableControlCommandTypeContainerValues = []EnableControlCommandTypeContainer{
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable0_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable1_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable2_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable3_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable4_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable5_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable6_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable7_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable8_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable9_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable10_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable11_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable12_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable13_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable14_2Bytes,
		EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable15_2Bytes,
	}
}

func (e EnableControlCommandTypeContainer) NumBytes() uint8 {
	switch e {
	case 0x02:
		{ /* '0x02' */
			return 2
		}
	case 0x0A:
		{ /* '0x0A' */
			return 2
		}
	case 0x12:
		{ /* '0x12' */
			return 2
		}
	case 0x1A:
		{ /* '0x1A' */
			return 2
		}
	case 0x22:
		{ /* '0x22' */
			return 2
		}
	case 0x2A:
		{ /* '0x2A' */
			return 2
		}
	case 0x32:
		{ /* '0x32' */
			return 2
		}
	case 0x3A:
		{ /* '0x3A' */
			return 2
		}
	case 0x42:
		{ /* '0x42' */
			return 2
		}
	case 0x4A:
		{ /* '0x4A' */
			return 2
		}
	case 0x52:
		{ /* '0x52' */
			return 2
		}
	case 0x5A:
		{ /* '0x5A' */
			return 2
		}
	case 0x62:
		{ /* '0x62' */
			return 2
		}
	case 0x6A:
		{ /* '0x6A' */
			return 2
		}
	case 0x72:
		{ /* '0x72' */
			return 2
		}
	case 0x7A:
		{ /* '0x7A' */
			return 2
		}
	default:
		{
			return 0
		}
	}
}

func EnableControlCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (EnableControlCommandTypeContainer, error) {
	for _, sizeValue := range EnableControlCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e EnableControlCommandTypeContainer) CommandType() EnableControlCommandType {
	switch e {
	case 0x02:
		{ /* '0x02' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x0A:
		{ /* '0x0A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x12:
		{ /* '0x12' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x1A:
		{ /* '0x1A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x22:
		{ /* '0x22' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x2A:
		{ /* '0x2A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x32:
		{ /* '0x32' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x3A:
		{ /* '0x3A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x42:
		{ /* '0x42' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x4A:
		{ /* '0x4A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x52:
		{ /* '0x52' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x5A:
		{ /* '0x5A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x62:
		{ /* '0x62' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x6A:
		{ /* '0x6A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x72:
		{ /* '0x72' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	case 0x7A:
		{ /* '0x7A' */
			return EnableControlCommandType_SET_NETWORK_VARIABLE
		}
	default:
		{
			return 0
		}
	}
}

func EnableControlCommandTypeContainerFirstEnumForFieldCommandType(value EnableControlCommandType) (EnableControlCommandTypeContainer, error) {
	for _, sizeValue := range EnableControlCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func EnableControlCommandTypeContainerByValue(value uint8) (enum EnableControlCommandTypeContainer, ok bool) {
	switch value {
	case 0x02:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable0_2Bytes, true
	case 0x0A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable1_2Bytes, true
	case 0x12:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable2_2Bytes, true
	case 0x1A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable3_2Bytes, true
	case 0x22:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable4_2Bytes, true
	case 0x2A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable5_2Bytes, true
	case 0x32:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable6_2Bytes, true
	case 0x3A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable7_2Bytes, true
	case 0x42:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable8_2Bytes, true
	case 0x4A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable9_2Bytes, true
	case 0x52:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable10_2Bytes, true
	case 0x5A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable11_2Bytes, true
	case 0x62:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable12_2Bytes, true
	case 0x6A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable13_2Bytes, true
	case 0x72:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable14_2Bytes, true
	case 0x7A:
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable15_2Bytes, true
	}
	return 0, false
}

func EnableControlCommandTypeContainerByName(value string) (enum EnableControlCommandTypeContainer, ok bool) {
	switch value {
	case "EnableControlCommandSetNetworkVariable0_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable0_2Bytes, true
	case "EnableControlCommandSetNetworkVariable1_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable1_2Bytes, true
	case "EnableControlCommandSetNetworkVariable2_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable2_2Bytes, true
	case "EnableControlCommandSetNetworkVariable3_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable3_2Bytes, true
	case "EnableControlCommandSetNetworkVariable4_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable4_2Bytes, true
	case "EnableControlCommandSetNetworkVariable5_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable5_2Bytes, true
	case "EnableControlCommandSetNetworkVariable6_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable6_2Bytes, true
	case "EnableControlCommandSetNetworkVariable7_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable7_2Bytes, true
	case "EnableControlCommandSetNetworkVariable8_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable8_2Bytes, true
	case "EnableControlCommandSetNetworkVariable9_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable9_2Bytes, true
	case "EnableControlCommandSetNetworkVariable10_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable10_2Bytes, true
	case "EnableControlCommandSetNetworkVariable11_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable11_2Bytes, true
	case "EnableControlCommandSetNetworkVariable12_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable12_2Bytes, true
	case "EnableControlCommandSetNetworkVariable13_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable13_2Bytes, true
	case "EnableControlCommandSetNetworkVariable14_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable14_2Bytes, true
	case "EnableControlCommandSetNetworkVariable15_2Bytes":
		return EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable15_2Bytes, true
	}
	return 0, false
}

func EnableControlCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range EnableControlCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastEnableControlCommandTypeContainer(structType interface{}) EnableControlCommandTypeContainer {
	castFunc := func(typ interface{}) EnableControlCommandTypeContainer {
		if sEnableControlCommandTypeContainer, ok := typ.(EnableControlCommandTypeContainer); ok {
			return sEnableControlCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m EnableControlCommandTypeContainer) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m EnableControlCommandTypeContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EnableControlCommandTypeContainerParse(ctx context.Context, theBytes []byte) (EnableControlCommandTypeContainer, error) {
	return EnableControlCommandTypeContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func EnableControlCommandTypeContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EnableControlCommandTypeContainer, error) {
	val, err := readBuffer.ReadUint8("EnableControlCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading EnableControlCommandTypeContainer")
	}
	if enum, ok := EnableControlCommandTypeContainerByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return EnableControlCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e EnableControlCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e EnableControlCommandTypeContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("EnableControlCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e EnableControlCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable0_2Bytes:
		return "EnableControlCommandSetNetworkVariable0_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable1_2Bytes:
		return "EnableControlCommandSetNetworkVariable1_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable2_2Bytes:
		return "EnableControlCommandSetNetworkVariable2_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable3_2Bytes:
		return "EnableControlCommandSetNetworkVariable3_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable4_2Bytes:
		return "EnableControlCommandSetNetworkVariable4_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable5_2Bytes:
		return "EnableControlCommandSetNetworkVariable5_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable6_2Bytes:
		return "EnableControlCommandSetNetworkVariable6_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable7_2Bytes:
		return "EnableControlCommandSetNetworkVariable7_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable8_2Bytes:
		return "EnableControlCommandSetNetworkVariable8_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable9_2Bytes:
		return "EnableControlCommandSetNetworkVariable9_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable10_2Bytes:
		return "EnableControlCommandSetNetworkVariable10_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable11_2Bytes:
		return "EnableControlCommandSetNetworkVariable11_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable12_2Bytes:
		return "EnableControlCommandSetNetworkVariable12_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable13_2Bytes:
		return "EnableControlCommandSetNetworkVariable13_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable14_2Bytes:
		return "EnableControlCommandSetNetworkVariable14_2Bytes"
	case EnableControlCommandTypeContainer_EnableControlCommandSetNetworkVariable15_2Bytes:
		return "EnableControlCommandSetNetworkVariable15_2Bytes"
	}
	return ""
}

func (e EnableControlCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
