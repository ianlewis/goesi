/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.2.9
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
	"time"
)

/* A list of GetCharactersCharacterIdOnlineOk. */
//easyjson:json
type GetCharactersCharacterIdOnlineOkList []GetCharactersCharacterIdOnlineOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdOnlineOk struct {
	LastLogin  time.Time `json:"last_login,omitempty"`  /* Timestamp of the last login */
	LastLogout time.Time `json:"last_logout,omitempty"` /* Timestamp of the last logout */
	Logins     int32     `json:"logins,omitempty"`      /* Total number of times the character has logged in */
	Online     bool      `json:"online,omitempty"`      /* If the character is online */
}
