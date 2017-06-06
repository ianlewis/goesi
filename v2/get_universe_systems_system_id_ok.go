/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.9.dev1
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

package goesiv2

/* 200 ok object */
type GetUniverseSystemsSystemIdOk struct {
	ConstellationId int32                              `json:"constellation_id,omitempty"` /* The constellation this solar system is in */
	Name            string                             `json:"name,omitempty"`             /* name string */
	Planets         []GetUniverseSystemsSystemIdPlanet `json:"planets,omitempty"`          /* planets array */
	Position        GetUniverseSystemsSystemIdPosition `json:"position,omitempty"`
	SecurityClass   string                             `json:"security_class,omitempty"`  /* security_class string */
	SecurityStatus  float32                            `json:"security_status,omitempty"` /* security_status number */
	Stargates       []int32                            `json:"stargates,omitempty"`       /* stargates array */
	SystemId        int32                              `json:"system_id,omitempty"`       /* system_id integer */
}
