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
package org.apache.plc4x.test;

import org.apache.commons.lang3.SystemUtils;
import org.junit.jupiter.api.extension.ConditionEvaluationResult;
import org.junit.jupiter.api.extension.ExecutionCondition;
import org.junit.jupiter.api.extension.ExtensionContext;

import java.io.*;

public class DisableInDockerFlagCondition implements ExecutionCondition {

    private static final boolean isDocker;
    static {
        boolean localIsDocker = false;
        if (SystemUtils.IS_OS_LINUX) {
            File dockerEnvFile = new File("/.dockerenv");
            if(dockerEnvFile.exists()) {
                localIsDocker = true;
            }
        }
        isDocker = localIsDocker;
    }

    @Override
    public ConditionEvaluationResult evaluateExecutionCondition(ExtensionContext extensionContext) {
        if(isDocker) {
            return ConditionEvaluationResult.disabled("Docker detected");
        }
        return ConditionEvaluationResult.enabled("Docker not detected");
    }

}
