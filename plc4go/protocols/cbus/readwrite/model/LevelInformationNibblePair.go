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

// LevelInformationNibblePair is an enum
type LevelInformationNibblePair uint8

type ILevelInformationNibblePair interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NibbleValue() uint8
}

const (
	LevelInformationNibblePair_Value_F LevelInformationNibblePair = 0x55
	LevelInformationNibblePair_Value_E LevelInformationNibblePair = 0x56
	LevelInformationNibblePair_Value_D LevelInformationNibblePair = 0x59
	LevelInformationNibblePair_Value_C LevelInformationNibblePair = 0x5A
	LevelInformationNibblePair_Value_B LevelInformationNibblePair = 0x65
	LevelInformationNibblePair_Value_A LevelInformationNibblePair = 0x66
	LevelInformationNibblePair_Value_9 LevelInformationNibblePair = 0x69
	LevelInformationNibblePair_Value_8 LevelInformationNibblePair = 0x6A
	LevelInformationNibblePair_Value_7 LevelInformationNibblePair = 0x95
	LevelInformationNibblePair_Value_6 LevelInformationNibblePair = 0x96
	LevelInformationNibblePair_Value_5 LevelInformationNibblePair = 0x99
	LevelInformationNibblePair_Value_4 LevelInformationNibblePair = 0x9A
	LevelInformationNibblePair_Value_3 LevelInformationNibblePair = 0xA5
	LevelInformationNibblePair_Value_2 LevelInformationNibblePair = 0xA6
	LevelInformationNibblePair_Value_1 LevelInformationNibblePair = 0xA9
	LevelInformationNibblePair_Value_0 LevelInformationNibblePair = 0xAA
)

var LevelInformationNibblePairValues []LevelInformationNibblePair

func init() {
	_ = errors.New
	LevelInformationNibblePairValues = []LevelInformationNibblePair{
		LevelInformationNibblePair_Value_F,
		LevelInformationNibblePair_Value_E,
		LevelInformationNibblePair_Value_D,
		LevelInformationNibblePair_Value_C,
		LevelInformationNibblePair_Value_B,
		LevelInformationNibblePair_Value_A,
		LevelInformationNibblePair_Value_9,
		LevelInformationNibblePair_Value_8,
		LevelInformationNibblePair_Value_7,
		LevelInformationNibblePair_Value_6,
		LevelInformationNibblePair_Value_5,
		LevelInformationNibblePair_Value_4,
		LevelInformationNibblePair_Value_3,
		LevelInformationNibblePair_Value_2,
		LevelInformationNibblePair_Value_1,
		LevelInformationNibblePair_Value_0,
	}
}

func (e LevelInformationNibblePair) NibbleValue() uint8 {
	switch e {
	case 0x55:
		{ /* '0x55' */
			return 0xF
		}
	case 0x56:
		{ /* '0x56' */
			return 0xE
		}
	case 0x59:
		{ /* '0x59' */
			return 0xD
		}
	case 0x5A:
		{ /* '0x5A' */
			return 0xC
		}
	case 0x65:
		{ /* '0x65' */
			return 0xB
		}
	case 0x66:
		{ /* '0x66' */
			return 0xA
		}
	case 0x69:
		{ /* '0x69' */
			return 0x9
		}
	case 0x6A:
		{ /* '0x6A' */
			return 0x8
		}
	case 0x95:
		{ /* '0x95' */
			return 0x7
		}
	case 0x96:
		{ /* '0x96' */
			return 0x6
		}
	case 0x99:
		{ /* '0x99' */
			return 0x5
		}
	case 0x9A:
		{ /* '0x9A' */
			return 0x4
		}
	case 0xA5:
		{ /* '0xA5' */
			return 0x3
		}
	case 0xA6:
		{ /* '0xA6' */
			return 0x2
		}
	case 0xA9:
		{ /* '0xA9' */
			return 0x1
		}
	case 0xAA:
		{ /* '0xAA' */
			return 0x0
		}
	default:
		{
			return 0
		}
	}
}

func LevelInformationNibblePairFirstEnumForFieldNibbleValue(value uint8) (LevelInformationNibblePair, error) {
	for _, sizeValue := range LevelInformationNibblePairValues {
		if sizeValue.NibbleValue() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NibbleValue not found", value)
}
func LevelInformationNibblePairByValue(value uint8) (enum LevelInformationNibblePair, ok bool) {
	switch value {
	case 0x55:
		return LevelInformationNibblePair_Value_F, true
	case 0x56:
		return LevelInformationNibblePair_Value_E, true
	case 0x59:
		return LevelInformationNibblePair_Value_D, true
	case 0x5A:
		return LevelInformationNibblePair_Value_C, true
	case 0x65:
		return LevelInformationNibblePair_Value_B, true
	case 0x66:
		return LevelInformationNibblePair_Value_A, true
	case 0x69:
		return LevelInformationNibblePair_Value_9, true
	case 0x6A:
		return LevelInformationNibblePair_Value_8, true
	case 0x95:
		return LevelInformationNibblePair_Value_7, true
	case 0x96:
		return LevelInformationNibblePair_Value_6, true
	case 0x99:
		return LevelInformationNibblePair_Value_5, true
	case 0x9A:
		return LevelInformationNibblePair_Value_4, true
	case 0xA5:
		return LevelInformationNibblePair_Value_3, true
	case 0xA6:
		return LevelInformationNibblePair_Value_2, true
	case 0xA9:
		return LevelInformationNibblePair_Value_1, true
	case 0xAA:
		return LevelInformationNibblePair_Value_0, true
	}
	return 0, false
}

func LevelInformationNibblePairByName(value string) (enum LevelInformationNibblePair, ok bool) {
	switch value {
	case "Value_F":
		return LevelInformationNibblePair_Value_F, true
	case "Value_E":
		return LevelInformationNibblePair_Value_E, true
	case "Value_D":
		return LevelInformationNibblePair_Value_D, true
	case "Value_C":
		return LevelInformationNibblePair_Value_C, true
	case "Value_B":
		return LevelInformationNibblePair_Value_B, true
	case "Value_A":
		return LevelInformationNibblePair_Value_A, true
	case "Value_9":
		return LevelInformationNibblePair_Value_9, true
	case "Value_8":
		return LevelInformationNibblePair_Value_8, true
	case "Value_7":
		return LevelInformationNibblePair_Value_7, true
	case "Value_6":
		return LevelInformationNibblePair_Value_6, true
	case "Value_5":
		return LevelInformationNibblePair_Value_5, true
	case "Value_4":
		return LevelInformationNibblePair_Value_4, true
	case "Value_3":
		return LevelInformationNibblePair_Value_3, true
	case "Value_2":
		return LevelInformationNibblePair_Value_2, true
	case "Value_1":
		return LevelInformationNibblePair_Value_1, true
	case "Value_0":
		return LevelInformationNibblePair_Value_0, true
	}
	return 0, false
}

func LevelInformationNibblePairKnows(value uint8) bool {
	for _, typeValue := range LevelInformationNibblePairValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLevelInformationNibblePair(structType interface{}) LevelInformationNibblePair {
	castFunc := func(typ interface{}) LevelInformationNibblePair {
		if sLevelInformationNibblePair, ok := typ.(LevelInformationNibblePair); ok {
			return sLevelInformationNibblePair
		}
		return 0
	}
	return castFunc(structType)
}

func (m LevelInformationNibblePair) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m LevelInformationNibblePair) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LevelInformationNibblePairParse(ctx context.Context, theBytes []byte) (LevelInformationNibblePair, error) {
	return LevelInformationNibblePairParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LevelInformationNibblePairParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LevelInformationNibblePair, error) {
	val, err := readBuffer.ReadUint8("LevelInformationNibblePair", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LevelInformationNibblePair")
	}
	if enum, ok := LevelInformationNibblePairByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return LevelInformationNibblePair(val), nil
	} else {
		return enum, nil
	}
}

func (e LevelInformationNibblePair) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LevelInformationNibblePair) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("LevelInformationNibblePair", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LevelInformationNibblePair) PLC4XEnumName() string {
	switch e {
	case LevelInformationNibblePair_Value_F:
		return "Value_F"
	case LevelInformationNibblePair_Value_E:
		return "Value_E"
	case LevelInformationNibblePair_Value_D:
		return "Value_D"
	case LevelInformationNibblePair_Value_C:
		return "Value_C"
	case LevelInformationNibblePair_Value_B:
		return "Value_B"
	case LevelInformationNibblePair_Value_A:
		return "Value_A"
	case LevelInformationNibblePair_Value_9:
		return "Value_9"
	case LevelInformationNibblePair_Value_8:
		return "Value_8"
	case LevelInformationNibblePair_Value_7:
		return "Value_7"
	case LevelInformationNibblePair_Value_6:
		return "Value_6"
	case LevelInformationNibblePair_Value_5:
		return "Value_5"
	case LevelInformationNibblePair_Value_4:
		return "Value_4"
	case LevelInformationNibblePair_Value_3:
		return "Value_3"
	case LevelInformationNibblePair_Value_2:
		return "Value_2"
	case LevelInformationNibblePair_Value_1:
		return "Value_1"
	case LevelInformationNibblePair_Value_0:
		return "Value_0"
	}
	return ""
}

func (e LevelInformationNibblePair) String() string {
	return e.PLC4XEnumName()
}
