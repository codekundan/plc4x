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

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.api;

import com.fasterxml.jackson.core.type.TypeReference;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.ApiClient;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.ApiException;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Configuration;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Pair;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.DLARString;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.SubscriptionData;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.SubscriptionProperties;
import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model.SubscriptionSettings;

import java.util.*;

@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class SubscriptionInformationAndSettingsApi {
  private ApiClient apiClient;

  public SubscriptionInformationAndSettingsApi() {
    this(Configuration.getDefaultApiClient());
  }

  public SubscriptionInformationAndSettingsApi(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  public ApiClient getApiClient() {
    return apiClient;
  }

  public void setApiClient(ApiClient apiClient) {
    this.apiClient = apiClient;
  }

  /**
   * Get all information about a subscription
   * Get all information about a subscription. It contains subscribed nodes and subscription properties.
   * @param clientName The ClientName is a string which contains the identification of the client. (required)
   * @param subscriptionName The SubscriptionName is a string which contains the identification of the subscription. (required)
   * @return SubscriptionData
   * @throws ApiException if fails to make API call
   */
  public SubscriptionData getClientSubscription(String clientName, String subscriptionName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'clientName' is set
    if (clientName == null) {
      throw new ApiException(400, "Missing the required parameter 'clientName' when calling getClientSubscription");
    }
    
    // verify the required parameter 'subscriptionName' is set
    if (subscriptionName == null) {
      throw new ApiException(400, "Missing the required parameter 'subscriptionName' when calling getClientSubscription");
    }
    
    // create path and map variables
    String localVarPath = "/nodes/datalayer/subscriptions/clients/{ClientName}/subscriptions/{SubscriptionName}"
      .replaceAll("\\{" + "ClientName" + "\\}", apiClient.escapeString(clientName.toString()))
      .replaceAll("\\{" + "SubscriptionName" + "\\}", apiClient.escapeString(subscriptionName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<SubscriptionData> localVarReturnType = new TypeReference<SubscriptionData>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
  /**
   * Get list of subscribed nodes for this client and subscription
   * Get list of subscribed nodes for this client and subscription.
   * @param clientName The ClientName is a string which contains the identification of the client. (required)
   * @param subscriptionName The SubscriptionName is a string which contains the identification of the subscription. (required)
   * @return DLARString
   * @throws ApiException if fails to make API call
   */
  public DLARString getClientSubscriptionNodes(String clientName, String subscriptionName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'clientName' is set
    if (clientName == null) {
      throw new ApiException(400, "Missing the required parameter 'clientName' when calling getClientSubscriptionNodes");
    }
    
    // verify the required parameter 'subscriptionName' is set
    if (subscriptionName == null) {
      throw new ApiException(400, "Missing the required parameter 'subscriptionName' when calling getClientSubscriptionNodes");
    }
    
    // create path and map variables
    String localVarPath = "/nodes/datalayer/subscriptions/clients/{ClientName}/subscriptions/{SubscriptionName}/nodes"
      .replaceAll("\\{" + "ClientName" + "\\}", apiClient.escapeString(clientName.toString()))
      .replaceAll("\\{" + "SubscriptionName" + "\\}", apiClient.escapeString(subscriptionName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<DLARString> localVarReturnType = new TypeReference<DLARString>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
  /**
   * Get properties for this client and subscription.
   * Get properties for this client and subscription.
   * @param clientName The ClientName is a string which contains the identification of the client. (required)
   * @param subscriptionName The SubscriptionName is a string which contains the identification of the subscription. (required)
   * @return SubscriptionProperties
   * @throws ApiException if fails to make API call
   */
  public SubscriptionProperties getClientSubscriptionProperties(String clientName, String subscriptionName) throws ApiException {
    Object localVarPostBody = null;
    
    // verify the required parameter 'clientName' is set
    if (clientName == null) {
      throw new ApiException(400, "Missing the required parameter 'clientName' when calling getClientSubscriptionProperties");
    }
    
    // verify the required parameter 'subscriptionName' is set
    if (subscriptionName == null) {
      throw new ApiException(400, "Missing the required parameter 'subscriptionName' when calling getClientSubscriptionProperties");
    }
    
    // create path and map variables
    String localVarPath = "/nodes/datalayer/subscriptions/clients/{ClientName}/subscriptions/{SubscriptionName}/properties"
      .replaceAll("\\{" + "ClientName" + "\\}", apiClient.escapeString(clientName.toString()))
      .replaceAll("\\{" + "SubscriptionName" + "\\}", apiClient.escapeString(subscriptionName.toString()));

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<SubscriptionProperties> localVarReturnType = new TypeReference<SubscriptionProperties>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
  /**
   * Data Layer Subscriptions - Global Settings
   * Global settings for Data Layer subscriptions.
   * @return SubscriptionSettings
   * @throws ApiException if fails to make API call
   */
  public SubscriptionSettings getSubscriptionSettings() throws ApiException {
    Object localVarPostBody = null;
    
    // create path and map variables
    String localVarPath = "/nodes/datalayer/subscriptions/settings";

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<SubscriptionSettings> localVarReturnType = new TypeReference<SubscriptionSettings>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "GET",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
  /**
   * Sets Data Layer Subscriptions - Global Settings
   * Sets global settings for Data Layer subscriptions. Global settings change will not affect existing subscriptions.
   * @param subscriptionSettings New global subscription settings. (optional)
   * @return SubscriptionSettings
   * @throws ApiException if fails to make API call
   */
  public SubscriptionSettings setSubscriptionSettings(SubscriptionSettings subscriptionSettings) throws ApiException {
    Object localVarPostBody = subscriptionSettings;
    
    // create path and map variables
    String localVarPath = "/nodes/datalayer/subscriptions/settings";

    StringJoiner localVarQueryStringJoiner = new StringJoiner("&");
    String localVarQueryParameterBaseName;
    List<Pair> localVarQueryParams = new ArrayList<Pair>();
    List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
    Map<String, String> localVarHeaderParams = new HashMap<String, String>();
    Map<String, String> localVarCookieParams = new HashMap<String, String>();
    Map<String, Object> localVarFormParams = new HashMap<String, Object>();

    
    
    
    final String[] localVarAccepts = {
      "application/json"
    };
    final String localVarAccept = apiClient.selectHeaderAccept(localVarAccepts);

    final String[] localVarContentTypes = {
      "application/json"
    };
    final String localVarContentType = apiClient.selectHeaderContentType(localVarContentTypes);

    String[] localVarAuthNames = new String[] { "Bearer", "UsernamePassword" };

    TypeReference<SubscriptionSettings> localVarReturnType = new TypeReference<SubscriptionSettings>() {};
    return apiClient.invokeAPI(
        localVarPath,
        "PUT",
        localVarQueryParams,
        localVarCollectionQueryParams,
        localVarQueryStringJoiner.toString(),
        localVarPostBody,
        localVarHeaderParams,
        localVarCookieParams,
        localVarFormParams,
        localVarAccept,
        localVarContentType,
        localVarAuthNames,
        localVarReturnType
    );
  }
}