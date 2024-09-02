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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// RsaEncryptedSecret is the corresponding interface of RsaEncryptedSecret
type RsaEncryptedSecret interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// RsaEncryptedSecretExactly can be used when we want exactly this type and not a type which fulfills RsaEncryptedSecret.
// This is useful for switch cases.
type RsaEncryptedSecretExactly interface {
	RsaEncryptedSecret
	isRsaEncryptedSecret() bool
}

// _RsaEncryptedSecret is the data-structure of this message
type _RsaEncryptedSecret struct {
}

var _ RsaEncryptedSecret = (*_RsaEncryptedSecret)(nil)

// NewRsaEncryptedSecret factory function for _RsaEncryptedSecret
func NewRsaEncryptedSecret() *_RsaEncryptedSecret {
	return &_RsaEncryptedSecret{}
}

// Deprecated: use the interface for direct cast
func CastRsaEncryptedSecret(structType any) RsaEncryptedSecret {
	if casted, ok := structType.(RsaEncryptedSecret); ok {
		return casted
	}
	if casted, ok := structType.(*RsaEncryptedSecret); ok {
		return *casted
	}
	return nil
}

func (m *_RsaEncryptedSecret) GetTypeName() string {
	return "RsaEncryptedSecret"
}

func (m *_RsaEncryptedSecret) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_RsaEncryptedSecret) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RsaEncryptedSecretParse(ctx context.Context, theBytes []byte) (RsaEncryptedSecret, error) {
	return RsaEncryptedSecretParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RsaEncryptedSecretParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (RsaEncryptedSecret, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (RsaEncryptedSecret, error) {
		return RsaEncryptedSecretParseWithBuffer(ctx, readBuffer)
	}
}

func RsaEncryptedSecretParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RsaEncryptedSecret, error) {
	v, err := (&_RsaEncryptedSecret{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_RsaEncryptedSecret) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__rsaEncryptedSecret RsaEncryptedSecret, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RsaEncryptedSecret"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RsaEncryptedSecret")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("RsaEncryptedSecret"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RsaEncryptedSecret")
	}

	return m, nil
}

func (m *_RsaEncryptedSecret) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RsaEncryptedSecret) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("RsaEncryptedSecret"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RsaEncryptedSecret")
	}

	if popErr := writeBuffer.PopContext("RsaEncryptedSecret"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RsaEncryptedSecret")
	}
	return nil
}

func (m *_RsaEncryptedSecret) isRsaEncryptedSecret() bool {
	return true
}

func (m *_RsaEncryptedSecret) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
