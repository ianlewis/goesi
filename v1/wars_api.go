/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.6.dev9
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv1

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type WarsApiService service

/* WarsApiService List wars
Return a list of wars  ---  Alternate route: &#x60;/legacy/wars/&#x60;  Alternate route: &#x60;/latest/wars/&#x60;  Alternate route: &#x60;/dev/wars/&#x60;   ---  This route is cached for up to 3600 seconds


@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "maxWarId" (int32) Only return wars with ID smaller than this.
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []int32*/
func (a *WarsApiService) GetWars(localVarOptionals map[string]interface{}) ([]int32, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []int32
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/wars/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxWarId"], "int32", "maxWarId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxWarId"].(int32); localVarOk {
		localVarQueryParams.Add("max_war_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WarsApiService Get war information
Return details about a war  ---  Alternate route: &#x60;/legacy/wars/{war_id}/&#x60;  Alternate route: &#x60;/latest/wars/{war_id}/&#x60;  Alternate route: &#x60;/dev/wars/{war_id}/&#x60;   ---  This route is cached for up to 3600 seconds


@param warId ID for a war
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return GetWarsWarIdOk*/
func (a *WarsApiService) GetWarsWarId(warId int32, localVarOptionals map[string]interface{}) (GetWarsWarIdOk, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetWarsWarIdOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/wars/{war_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"war_id"+"}", fmt.Sprintf("%v", warId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* WarsApiService List kills for a war
Return a list of kills related to a war  ---  Alternate route: &#x60;/legacy/wars/{war_id}/killmails/&#x60;  Alternate route: &#x60;/latest/wars/{war_id}/killmails/&#x60;  Alternate route: &#x60;/dev/wars/{war_id}/killmails/&#x60;   ---  This route is cached for up to 3600 seconds


@param warId A valid war ID
@param optional (nil or map[string]interface{}) with one or more of:
    @param "datasource" (string) The server name you would like data from
    @param "page" (int32) Which page to query, starting at 1, 2000 killmails per page.
    @param "userAgent" (string) Client identifier, takes precedence over headers
    @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
@return []GetWarsWarIdKillmails200Ok*/
func (a *WarsApiService) GetWarsWarIdKillmails(warId int32, localVarOptionals map[string]interface{}) ([]GetWarsWarIdKillmails200Ok, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []GetWarsWarIdKillmails200Ok
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/wars/{war_id}/killmails/"
	localVarPath = strings.Replace(localVarPath, "{"+"war_id"+"}", fmt.Sprintf("%v", warId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	r, err := a.client.prepareRequest(nil, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
