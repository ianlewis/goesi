/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.0
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

package esi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"fmt"

	"github.com/antihax/goesi/optional"
)

// Linger please
var (
	_ context.Context
)

type CalendarApiService service

/*
CalendarApiService List calendar event summaries
Get 50 event summaries from the calendar. If no from_event ID is given, the resource will return the next 50 chronological event summaries from now. If a from_event ID is specified, it will return the next 50 chronological event summaries from after that event.  ---  This route is cached for up to 5 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param optional nil or *GetCharactersCharacterIdCalendarOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "FromEvent" (optional.Int32) -  The event ID to retrieve events from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []GetCharactersCharacterIdCalendar200Ok
*/

type GetCharactersCharacterIdCalendarOpts struct {
	Datasource optional.String
	FromEvent  optional.Int32
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *CalendarApiService) GetCharactersCharacterIdCalendar(ctx context.Context, characterId int32, localVarOptionals *GetCharactersCharacterIdCalendarOpts) ([]GetCharactersCharacterIdCalendar200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdCalendar200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/calendar/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromEvent.IsSet() {
		localVarQueryParams.Add("from_event", parameterToString(localVarOptionals.FromEvent.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []GetCharactersCharacterIdCalendar200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 502 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v BadGateway
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CalendarApiService Get an event
Get all the information for a specific event  ---  This route is cached for up to 5 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param eventId The id of the event requested
 * @param optional nil or *GetCharactersCharacterIdCalendarEventIdOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return GetCharactersCharacterIdCalendarEventIdOk
*/

type GetCharactersCharacterIdCalendarEventIdOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *CalendarApiService) GetCharactersCharacterIdCalendarEventId(ctx context.Context, characterId int32, eventId int32, localVarOptionals *GetCharactersCharacterIdCalendarEventIdOpts) (GetCharactersCharacterIdCalendarEventIdOk, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdCalendarEventIdOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v3/characters/{character_id}/calendar/{event_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", fmt.Sprintf("%v", eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v GetCharactersCharacterIdCalendarEventIdOk
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 502 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v BadGateway
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CalendarApiService Get attendees
Get all invited attendees for a given event  ---  This route is cached for up to 600 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param eventId The id of the event requested
 * @param optional nil or *GetCharactersCharacterIdCalendarEventIdAttendeesOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent

@return []GetCharactersCharacterIdCalendarEventIdAttendees200Ok
*/

type GetCharactersCharacterIdCalendarEventIdAttendeesOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *CalendarApiService) GetCharactersCharacterIdCalendarEventIdAttendees(ctx context.Context, characterId int32, eventId int32, localVarOptionals *GetCharactersCharacterIdCalendarEventIdAttendeesOpts) ([]GetCharactersCharacterIdCalendarEventIdAttendees200Ok, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetCharactersCharacterIdCalendarEventIdAttendees200OkList
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/characters/{character_id}/calendar/{event_id}/attendees/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", fmt.Sprintf("%v", eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return localVarReturnValue, nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v []GetCharactersCharacterIdCalendarEventIdAttendees200Ok
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 502 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v BadGateway
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarReturnValue, localVarHttpResponse, err
			}

			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CalendarApiService Respond to an event
Set your response status to an event  ---
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param characterId An EVE character ID
 * @param eventId The ID of the event requested
 * @param response The response value to set, overriding current value.
 * @param optional nil or *PutCharactersCharacterIdCalendarEventIdOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "Token" (optional.String) -  Access token to use if unable to set a header
     * @param "UserAgent" (optional.String) -  Client identifier, takes precedence over headers
     * @param "XUserAgent" (optional.String) -  Client identifier, takes precedence over User-Agent


*/

type PutCharactersCharacterIdCalendarEventIdOpts struct {
	Datasource optional.String
	Token      optional.String
	UserAgent  optional.String
	XUserAgent optional.String
}

func (a *CalendarApiService) PutCharactersCharacterIdCalendarEventId(ctx context.Context, characterId int32, eventId int32, response PutCharactersCharacterIdCalendarEventIdResponse, localVarOptionals *PutCharactersCharacterIdCalendarEventIdOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v3/characters/{character_id}/calendar/{event_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"event_id"+"}", fmt.Sprintf("%v", eventId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if characterId < 1 {
		return nil, reportError("characterId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserAgent.IsSet() {
		localVarQueryParams.Add("user_agent", parameterToString(localVarOptionals.UserAgent.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XUserAgent.IsSet() {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarOptionals.XUserAgent.Value(), "")
	}
	// body params
	localVarPostBody = &response
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 400 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarHttpResponse, err
			}

			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarHttpResponse, err
			}

			var v Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarHttpResponse, err
			}

			var v Forbidden
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarHttpResponse, err
			}

			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 502 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarHttpResponse, err
			}

			var v BadGateway
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
			if err != nil {
				return localVarHttpResponse, err
			}

			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}