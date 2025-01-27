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

package npdu

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/pdu"
	"github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
)

type RouterBusyToNetwork struct {
	*_NPDU

	messageType uint8

	rbtnNetworkList []uint16
}

func NewRouterBusyToNetwork(args Args, kwArgs KWArgs, options ...Option) (*RouterBusyToNetwork, error) {
	r := &RouterBusyToNetwork{
		messageType: 0x04,
	}
	ApplyAppliers(options, r)
	options = AddLeafTypeIfAbundant(options, r)
	options = AddNLMIfAbundant(options, model.NewNLMRouterBusyToNetwork(r.rbtnNetworkList, 0))
	npdu, err := NewNPDU(args, kwArgs, options...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating NPDU")
	}
	r._NPDU = npdu.(*_NPDU)
	r.AddDebugContents(r, "rbtnNetworkList")

	r.npduNetMessage = &r.messageType
	return r, nil
}

// TODO: check if this is rather a KWArgs
func WithRouterBusyToNetworkDnet(networkList []uint16) GenericApplier[*RouterBusyToNetwork] {
	return WrapGenericApplier(func(n *RouterBusyToNetwork) { n.rbtnNetworkList = networkList })
}

func (r *RouterBusyToNetwork) GetDebugAttr(attr string) any {
	switch attr {
	case "rbtnNetworkList":
		return r.rbtnNetworkList
	}
	return nil
}

func (r *RouterBusyToNetwork) GetRbtnNetworkList() []uint16 {
	return r.rbtnNetworkList
}

func (r *RouterBusyToNetwork) Encode(npdu Arg) error {
	switch npdu := npdu.(type) {
	case NPCI:
		if err := npdu.GetNPCI().Update(r); err != nil {
			return errors.Wrap(err, "error updating NPDU")
		}
	}
	switch npdu := npdu.(type) {
	case PDUData:
		for _, net := range r.GetRbtnNetworkList() {
			npdu.PutShort(net)
		}
	default:
		return errors.Errorf("invalid NPDU type %T", npdu)
	}
	return nil
}

func (r *RouterBusyToNetwork) Decode(npdu Arg) error {
	if err := r._NPCI.Update(npdu); err != nil {
		return errors.Wrap(err, "error updating NPCI")
	}
	switch npdu := npdu.(type) {
	case NPDU:
		switch rm := npdu.GetRootMessage().(type) {
		case model.NPDU:
			switch nlm := rm.GetNlm().(type) {
			case model.NLMRouterBusyToNetwork:
				r.rbtnNetworkList = nlm.GetDestinationNetworkAddresses()
				r.SetRootMessage(rm)
			}
		}
	}
	switch npdu := npdu.(type) {
	case PDUData:
		r.SetPduData(npdu.GetPduData())
	}
	return nil
}
