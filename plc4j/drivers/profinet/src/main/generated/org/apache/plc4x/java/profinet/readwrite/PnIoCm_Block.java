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
package org.apache.plc4x.java.profinet.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class PnIoCm_Block implements Message {

  // Abstract accessors for discriminator values.
  public abstract PnIoCm_BlockType getBlockType();

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;

  public PnIoCm_Block(short blockVersionHigh, short blockVersionLow) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  protected abstract void serializePnIoCm_BlockChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PnIoCm_Block");

    // Discriminator Field (blockType) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "blockType",
        "PnIoCm_BlockType",
        getBlockType(),
        new DataWriterEnumDefault<>(
            PnIoCm_BlockType::getValue, PnIoCm_BlockType::name, writeUnsignedInt(writeBuffer, 16)));

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField("blockLength", blockLength, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch field (Serialize the sub-type)
    serializePnIoCm_BlockChild(writeBuffer);

    writeBuffer.popContext("PnIoCm_Block");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnIoCm_Block _value = this;

    // Discriminator Field (blockType)
    lengthInBits += 16;

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static PnIoCm_Block staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnIoCm_Block staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PnIoCm_Block");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    PnIoCm_BlockType blockType =
        readDiscriminatorField(
            "blockType",
            new DataReaderEnumDefault<>(
                PnIoCm_BlockType::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    PnIoCm_BlockBuilder builder = null;
    if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.AR_BLOCK_REQ)) {
      builder = PnIoCm_Block_ArReq.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.AR_BLOCK_RES)) {
      builder = PnIoCm_Block_ArRes.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IO_CR_BLOCK_REQ)) {
      builder = PnIoCm_Block_IoCrReq.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.IO_CR_BLOCK_RES)) {
      builder = PnIoCm_Block_IoCrRes.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.ALARM_CR_BLOCK_REQ)) {
      builder = PnIoCm_Block_AlarmCrReq.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.ALARM_CR_BLOCK_RES)) {
      builder = PnIoCm_Block_AlarmCrRes.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.EXPECTED_SUBMODULE_BLOCK_REQ)) {
      builder = PnIoCm_Block_ExpectedSubmoduleReq.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.MODULE_DIFF_BLOCK)) {
      builder = PnIoCm_Block_ModuleDiff.staticParseBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, PnIoCm_BlockType.AR_SERVER_BLOCK)) {
      builder = PnIoCm_Block_ArServer.staticParseBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "blockType="
              + blockType
              + "]");
    }

    readBuffer.closeContext("PnIoCm_Block");
    // Create the instance
    PnIoCm_Block _pnIoCm_Block = builder.build(blockVersionHigh, blockVersionLow);
    return _pnIoCm_Block;
  }

  public static interface PnIoCm_BlockBuilder {
    PnIoCm_Block build(short blockVersionHigh, short blockVersionLow);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block)) {
      return false;
    }
    PnIoCm_Block that = (PnIoCm_Block) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getBlockVersionHigh(), getBlockVersionLow());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}