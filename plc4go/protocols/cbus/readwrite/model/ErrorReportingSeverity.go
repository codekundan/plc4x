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

// ErrorReportingSeverity is an enum
type ErrorReportingSeverity uint8

type IErrorReportingSeverity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ErrorReportingSeverity_ALL_OK          ErrorReportingSeverity = 0x0
	ErrorReportingSeverity_OK              ErrorReportingSeverity = 0x1
	ErrorReportingSeverity_MINOR_FAILURE   ErrorReportingSeverity = 0x2
	ErrorReportingSeverity_GENERAL_FAILURE ErrorReportingSeverity = 0x3
	ErrorReportingSeverity_EXTREME_FAILURE ErrorReportingSeverity = 0x4
	ErrorReportingSeverity_RESERVED_1      ErrorReportingSeverity = 0x5
	ErrorReportingSeverity_RESERVED_2      ErrorReportingSeverity = 0x6
	ErrorReportingSeverity_RESERVED_3      ErrorReportingSeverity = 0x7
)

var ErrorReportingSeverityValues []ErrorReportingSeverity

func init() {
	_ = errors.New
	ErrorReportingSeverityValues = []ErrorReportingSeverity{
		ErrorReportingSeverity_ALL_OK,
		ErrorReportingSeverity_OK,
		ErrorReportingSeverity_MINOR_FAILURE,
		ErrorReportingSeverity_GENERAL_FAILURE,
		ErrorReportingSeverity_EXTREME_FAILURE,
		ErrorReportingSeverity_RESERVED_1,
		ErrorReportingSeverity_RESERVED_2,
		ErrorReportingSeverity_RESERVED_3,
	}
}

func ErrorReportingSeverityByValue(value uint8) (enum ErrorReportingSeverity, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSeverity_ALL_OK, true
	case 0x1:
		return ErrorReportingSeverity_OK, true
	case 0x2:
		return ErrorReportingSeverity_MINOR_FAILURE, true
	case 0x3:
		return ErrorReportingSeverity_GENERAL_FAILURE, true
	case 0x4:
		return ErrorReportingSeverity_EXTREME_FAILURE, true
	case 0x5:
		return ErrorReportingSeverity_RESERVED_1, true
	case 0x6:
		return ErrorReportingSeverity_RESERVED_2, true
	case 0x7:
		return ErrorReportingSeverity_RESERVED_3, true
	}
	return 0, false
}

func ErrorReportingSeverityByName(value string) (enum ErrorReportingSeverity, ok bool) {
	switch value {
	case "ALL_OK":
		return ErrorReportingSeverity_ALL_OK, true
	case "OK":
		return ErrorReportingSeverity_OK, true
	case "MINOR_FAILURE":
		return ErrorReportingSeverity_MINOR_FAILURE, true
	case "GENERAL_FAILURE":
		return ErrorReportingSeverity_GENERAL_FAILURE, true
	case "EXTREME_FAILURE":
		return ErrorReportingSeverity_EXTREME_FAILURE, true
	case "RESERVED_1":
		return ErrorReportingSeverity_RESERVED_1, true
	case "RESERVED_2":
		return ErrorReportingSeverity_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSeverity_RESERVED_3, true
	}
	return 0, false
}

func ErrorReportingSeverityKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSeverityValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSeverity(structType interface{}) ErrorReportingSeverity {
	castFunc := func(typ interface{}) ErrorReportingSeverity {
		if sErrorReportingSeverity, ok := typ.(ErrorReportingSeverity); ok {
			return sErrorReportingSeverity
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSeverity) GetLengthInBits(ctx context.Context) uint16 {
	return 3
}

func (m ErrorReportingSeverity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSeverityParse(ctx context.Context, theBytes []byte) (ErrorReportingSeverity, error) {
	return ErrorReportingSeverityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSeverityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSeverity, error) {
	val, err := readBuffer.ReadUint8("ErrorReportingSeverity", 3)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSeverity")
	}
	if enum, ok := ErrorReportingSeverityByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingSeverity(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSeverity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSeverity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ErrorReportingSeverity", 3, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSeverity) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSeverity_ALL_OK:
		return "ALL_OK"
	case ErrorReportingSeverity_OK:
		return "OK"
	case ErrorReportingSeverity_MINOR_FAILURE:
		return "MINOR_FAILURE"
	case ErrorReportingSeverity_GENERAL_FAILURE:
		return "GENERAL_FAILURE"
	case ErrorReportingSeverity_EXTREME_FAILURE:
		return "EXTREME_FAILURE"
	case ErrorReportingSeverity_RESERVED_1:
		return "RESERVED_1"
	case ErrorReportingSeverity_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSeverity_RESERVED_3:
		return "RESERVED_3"
	}
	return ""
}

func (e ErrorReportingSeverity) String() string {
	return e.PLC4XEnumName()
}
