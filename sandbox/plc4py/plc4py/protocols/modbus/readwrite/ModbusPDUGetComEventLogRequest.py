#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from ctypes import c_bool
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDUBuilder
from plc4py.spi.generation.WriteBuffer import WriteBuffer
import math


@dataclass
class ModbusPDUGetComEventLogRequest(PlcMessage, ModbusPDU):
    # Accessors for discriminator values.
    error_flag: c_bool = c_bool(false)
    function_flag: c_uint8 = 0x0C
    response: c_bool = c_bool(false)

    def __post_init__(self):
        super().__init__()

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        start_pos: int = write_buffer.get_pos()
        write_buffer.push_context("ModbusPDUGetComEventLogRequest")

        write_buffer.pop_context("ModbusPDUGetComEventLogRequest")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = super().get_length_in_bits()
        _value: ModbusPDUGetComEventLogRequest = self

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: c_bool):
        read_buffer.pull_context("ModbusPDUGetComEventLogRequest")
        start_pos: int = read_buffer.get_pos()
        cur_pos: int = 0

        read_buffer.close_context("ModbusPDUGetComEventLogRequest")
        # Create the instance
        return ModbusPDUGetComEventLogRequestBuilder()

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUGetComEventLogRequest):
            return False

        that: ModbusPDUGetComEventLogRequest = ModbusPDUGetComEventLogRequest(o)
        return super().equals(that) and True

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            write_buffer_box_based.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class ModbusPDUGetComEventLogRequestBuilder(ModbusPDUBuilder):
    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUGetComEventLogRequest:
        modbus_pdu_get_com_event_log_request: ModbusPDUGetComEventLogRequest = (
            ModbusPDUGetComEventLogRequest()
        )
        return modbus_pdu_get_com_event_log_request
