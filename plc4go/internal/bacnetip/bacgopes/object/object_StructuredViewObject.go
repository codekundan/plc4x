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

package object

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type StructuredViewObject struct {
	Object
}

func NewStructuredViewObject(options ...Option) (*StructuredViewObject, error) {
	s := new(StructuredViewObject)
	objectType := "structuredView"
	properties := []Property{
		NewReadableProperty("nodeType", V2P(NewNodeType)),
		NewOptionalProperty("nodeSubtype", V2P(NewCharacterString)),
		NewReadableProperty("subordinateList", ArrayOfP(NewDeviceObjectReference, 0, 0)),
		NewOptionalProperty("subordinateAnnotations", ArrayOfP(NewCharacterString, 0, 0)),
		NewOptionalProperty("subordinateTags", ArrayOfP(NewNameValueCollection, 0, 0)),
		NewOptionalProperty("subordinateNodeTypes", ArrayOfP(NewNodeType, 0, 0)),
		NewOptionalProperty("subordinateRelationships", ArrayOfP(NewRelationship, 0, 0)),
		NewOptionalProperty("defaultSubordinateRelationship", V2P(NewRelationship)),
		NewOptionalProperty("represents", V2P(NewDeviceObjectReference)),
	}
	var err error
	s.Object, err = NewObject(Combine(options, WithObjectType(objectType), WithObjectExtraProperties(properties))...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating object")
	}
	if _, err := RegisterObjectType(NKW(KWCls, s)); err != nil {
		return nil, errors.Wrap(err, "error registering object type")
	}
	return s, nil
}
