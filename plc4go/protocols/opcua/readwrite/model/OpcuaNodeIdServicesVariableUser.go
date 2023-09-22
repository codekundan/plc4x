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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableUser is an enum
type OpcuaNodeIdServicesVariableUser int32

type IOpcuaNodeIdServicesVariableUser interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableUser_UserTokenType_EnumStrings                        OpcuaNodeIdServicesVariableUser = 7596
	OpcuaNodeIdServicesVariableUser_UserManagementType_Users                         OpcuaNodeIdServicesVariableUser = 24265
	OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordLength                OpcuaNodeIdServicesVariableUser = 24266
	OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordOptions               OpcuaNodeIdServicesVariableUser = 24267
	OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordRestrictions          OpcuaNodeIdServicesVariableUser = 24268
	OpcuaNodeIdServicesVariableUser_UserManagementType_AddUser_InputArguments        OpcuaNodeIdServicesVariableUser = 24270
	OpcuaNodeIdServicesVariableUser_UserManagementType_ModifyUser_InputArguments     OpcuaNodeIdServicesVariableUser = 24272
	OpcuaNodeIdServicesVariableUser_UserManagementType_RemoveUser_InputArguments     OpcuaNodeIdServicesVariableUser = 24274
	OpcuaNodeIdServicesVariableUser_UserManagementType_ChangePassword_InputArguments OpcuaNodeIdServicesVariableUser = 24276
	OpcuaNodeIdServicesVariableUser_UserConfigurationMask_OptionSetValues            OpcuaNodeIdServicesVariableUser = 24280
	OpcuaNodeIdServicesVariableUser_UserManagement_PasswordRestrictions              OpcuaNodeIdServicesVariableUser = 24291
	OpcuaNodeIdServicesVariableUser_UserManagement_Users                             OpcuaNodeIdServicesVariableUser = 24301
	OpcuaNodeIdServicesVariableUser_UserManagement_PasswordLength                    OpcuaNodeIdServicesVariableUser = 24302
	OpcuaNodeIdServicesVariableUser_UserManagement_PasswordOptions                   OpcuaNodeIdServicesVariableUser = 24303
	OpcuaNodeIdServicesVariableUser_UserManagement_AddUser_InputArguments            OpcuaNodeIdServicesVariableUser = 24305
	OpcuaNodeIdServicesVariableUser_UserManagement_ModifyUser_InputArguments         OpcuaNodeIdServicesVariableUser = 24307
	OpcuaNodeIdServicesVariableUser_UserManagement_RemoveUser_InputArguments         OpcuaNodeIdServicesVariableUser = 24309
	OpcuaNodeIdServicesVariableUser_UserManagement_ChangePassword_InputArguments     OpcuaNodeIdServicesVariableUser = 24311
)

var OpcuaNodeIdServicesVariableUserValues []OpcuaNodeIdServicesVariableUser

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableUserValues = []OpcuaNodeIdServicesVariableUser{
		OpcuaNodeIdServicesVariableUser_UserTokenType_EnumStrings,
		OpcuaNodeIdServicesVariableUser_UserManagementType_Users,
		OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordLength,
		OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordOptions,
		OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordRestrictions,
		OpcuaNodeIdServicesVariableUser_UserManagementType_AddUser_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserManagementType_ModifyUser_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserManagementType_RemoveUser_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserManagementType_ChangePassword_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserConfigurationMask_OptionSetValues,
		OpcuaNodeIdServicesVariableUser_UserManagement_PasswordRestrictions,
		OpcuaNodeIdServicesVariableUser_UserManagement_Users,
		OpcuaNodeIdServicesVariableUser_UserManagement_PasswordLength,
		OpcuaNodeIdServicesVariableUser_UserManagement_PasswordOptions,
		OpcuaNodeIdServicesVariableUser_UserManagement_AddUser_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserManagement_ModifyUser_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserManagement_RemoveUser_InputArguments,
		OpcuaNodeIdServicesVariableUser_UserManagement_ChangePassword_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableUserByValue(value int32) (enum OpcuaNodeIdServicesVariableUser, ok bool) {
	switch value {
	case 24265:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_Users, true
	case 24266:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordLength, true
	case 24267:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordOptions, true
	case 24268:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordRestrictions, true
	case 24270:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_AddUser_InputArguments, true
	case 24272:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_ModifyUser_InputArguments, true
	case 24274:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_RemoveUser_InputArguments, true
	case 24276:
		return OpcuaNodeIdServicesVariableUser_UserManagementType_ChangePassword_InputArguments, true
	case 24280:
		return OpcuaNodeIdServicesVariableUser_UserConfigurationMask_OptionSetValues, true
	case 24291:
		return OpcuaNodeIdServicesVariableUser_UserManagement_PasswordRestrictions, true
	case 24301:
		return OpcuaNodeIdServicesVariableUser_UserManagement_Users, true
	case 24302:
		return OpcuaNodeIdServicesVariableUser_UserManagement_PasswordLength, true
	case 24303:
		return OpcuaNodeIdServicesVariableUser_UserManagement_PasswordOptions, true
	case 24305:
		return OpcuaNodeIdServicesVariableUser_UserManagement_AddUser_InputArguments, true
	case 24307:
		return OpcuaNodeIdServicesVariableUser_UserManagement_ModifyUser_InputArguments, true
	case 24309:
		return OpcuaNodeIdServicesVariableUser_UserManagement_RemoveUser_InputArguments, true
	case 24311:
		return OpcuaNodeIdServicesVariableUser_UserManagement_ChangePassword_InputArguments, true
	case 7596:
		return OpcuaNodeIdServicesVariableUser_UserTokenType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUserByName(value string) (enum OpcuaNodeIdServicesVariableUser, ok bool) {
	switch value {
	case "UserManagementType_Users":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_Users, true
	case "UserManagementType_PasswordLength":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordLength, true
	case "UserManagementType_PasswordOptions":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordOptions, true
	case "UserManagementType_PasswordRestrictions":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordRestrictions, true
	case "UserManagementType_AddUser_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_AddUser_InputArguments, true
	case "UserManagementType_ModifyUser_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_ModifyUser_InputArguments, true
	case "UserManagementType_RemoveUser_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_RemoveUser_InputArguments, true
	case "UserManagementType_ChangePassword_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagementType_ChangePassword_InputArguments, true
	case "UserConfigurationMask_OptionSetValues":
		return OpcuaNodeIdServicesVariableUser_UserConfigurationMask_OptionSetValues, true
	case "UserManagement_PasswordRestrictions":
		return OpcuaNodeIdServicesVariableUser_UserManagement_PasswordRestrictions, true
	case "UserManagement_Users":
		return OpcuaNodeIdServicesVariableUser_UserManagement_Users, true
	case "UserManagement_PasswordLength":
		return OpcuaNodeIdServicesVariableUser_UserManagement_PasswordLength, true
	case "UserManagement_PasswordOptions":
		return OpcuaNodeIdServicesVariableUser_UserManagement_PasswordOptions, true
	case "UserManagement_AddUser_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagement_AddUser_InputArguments, true
	case "UserManagement_ModifyUser_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagement_ModifyUser_InputArguments, true
	case "UserManagement_RemoveUser_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagement_RemoveUser_InputArguments, true
	case "UserManagement_ChangePassword_InputArguments":
		return OpcuaNodeIdServicesVariableUser_UserManagement_ChangePassword_InputArguments, true
	case "UserTokenType_EnumStrings":
		return OpcuaNodeIdServicesVariableUser_UserTokenType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableUserKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableUserValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableUser(structType any) OpcuaNodeIdServicesVariableUser {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableUser {
		if sOpcuaNodeIdServicesVariableUser, ok := typ.(OpcuaNodeIdServicesVariableUser); ok {
			return sOpcuaNodeIdServicesVariableUser
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableUser) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableUser) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableUserParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableUser, error) {
	return OpcuaNodeIdServicesVariableUserParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableUserParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableUser, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableUser", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableUser")
	}
	if enum, ok := OpcuaNodeIdServicesVariableUserByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableUser")
		return OpcuaNodeIdServicesVariableUser(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableUser) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableUser) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableUser", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableUser) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableUser_UserManagementType_Users:
		return "UserManagementType_Users"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordLength:
		return "UserManagementType_PasswordLength"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordOptions:
		return "UserManagementType_PasswordOptions"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_PasswordRestrictions:
		return "UserManagementType_PasswordRestrictions"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_AddUser_InputArguments:
		return "UserManagementType_AddUser_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_ModifyUser_InputArguments:
		return "UserManagementType_ModifyUser_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_RemoveUser_InputArguments:
		return "UserManagementType_RemoveUser_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserManagementType_ChangePassword_InputArguments:
		return "UserManagementType_ChangePassword_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserConfigurationMask_OptionSetValues:
		return "UserConfigurationMask_OptionSetValues"
	case OpcuaNodeIdServicesVariableUser_UserManagement_PasswordRestrictions:
		return "UserManagement_PasswordRestrictions"
	case OpcuaNodeIdServicesVariableUser_UserManagement_Users:
		return "UserManagement_Users"
	case OpcuaNodeIdServicesVariableUser_UserManagement_PasswordLength:
		return "UserManagement_PasswordLength"
	case OpcuaNodeIdServicesVariableUser_UserManagement_PasswordOptions:
		return "UserManagement_PasswordOptions"
	case OpcuaNodeIdServicesVariableUser_UserManagement_AddUser_InputArguments:
		return "UserManagement_AddUser_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserManagement_ModifyUser_InputArguments:
		return "UserManagement_ModifyUser_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserManagement_RemoveUser_InputArguments:
		return "UserManagement_RemoveUser_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserManagement_ChangePassword_InputArguments:
		return "UserManagement_ChangePassword_InputArguments"
	case OpcuaNodeIdServicesVariableUser_UserTokenType_EnumStrings:
		return "UserTokenType_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableUser) String() string {
	return e.PLC4XEnumName()
}