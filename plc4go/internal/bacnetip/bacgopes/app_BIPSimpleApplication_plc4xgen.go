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

// Code generated by "plc4xGenerator -type=BIPSimpleApplication -prefix=app_"; DO NOT EDIT.

package bacgopes

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *BIPSimpleApplication) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *BIPSimpleApplication) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("BIPSimpleApplication"); err != nil {
		return err
	}
	if err := d.ApplicationIOController.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}
	if err := d.WhoIsIAmServices.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}
	if err := d.ReadWritePropertyServices.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.localAddress)

		if err := writeBuffer.WriteString("localAddress", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if d.asap != nil {
		{
			_value := fmt.Sprintf("%v", d.asap)

			if err := writeBuffer.WriteString("asap", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.smap != nil {
		{
			_value := fmt.Sprintf("%v", d.smap)

			if err := writeBuffer.WriteString("smap", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.nsap != nil {
		{
			_value := fmt.Sprintf("%v", d.nsap)

			if err := writeBuffer.WriteString("nsap", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.nse != nil {
		{
			_value := fmt.Sprintf("%v", d.nse)

			if err := writeBuffer.WriteString("nse", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.bip != nil {
		{
			_value := fmt.Sprintf("%v", d.bip)

			if err := writeBuffer.WriteString("bip", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.annexj != nil {
		{
			_value := fmt.Sprintf("%v", d.annexj)

			if err := writeBuffer.WriteString("annexj", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.mux != nil {
		{
			_value := fmt.Sprintf("%v", d.mux)

			if err := writeBuffer.WriteString("mux", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("BIPSimpleApplication"); err != nil {
		return err
	}
	return nil
}

func (d *BIPSimpleApplication) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
