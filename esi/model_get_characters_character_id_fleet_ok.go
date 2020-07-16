/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.7.2
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

/* A list of GetCharactersCharacterIdFleetOk. */
//easyjson:json
type GetCharactersCharacterIdFleetOkList []GetCharactersCharacterIdFleetOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdFleetOk struct {
	FleetId int64  `json:"fleet_id,omitempty"` /* The character's current fleet ID */
	Role    string `json:"role,omitempty"`     /* Member’s role in fleet */
	SquadId int64  `json:"squad_id,omitempty"` /* ID of the squad the member is in. If not applicable, will be set to -1 */
	WingId  int64  `json:"wing_id,omitempty"`  /* ID of the wing the member is in. If not applicable, will be set to -1 */
}
