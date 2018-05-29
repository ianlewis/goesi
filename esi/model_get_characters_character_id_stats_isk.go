/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.3
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

/* A list of GetCharactersCharacterIdStatsIsk. */
//easyjson:json
type GetCharactersCharacterIdStatsIskList []GetCharactersCharacterIdStatsIsk

/* isk object */
//easyjson:json
type GetCharactersCharacterIdStatsIsk struct {
	In  int64 `json:"in,omitempty"`  /* in integer */
	Out int64 `json:"out,omitempty"` /* out integer */
}
