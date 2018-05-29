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

/* A list of GetCorporationsCorporationIdCustomsOffices200Ok. */
//easyjson:json
type GetCorporationsCorporationIdCustomsOffices200OkList []GetCorporationsCorporationIdCustomsOffices200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdCustomsOffices200Ok struct {
	AllianceTaxRate          float32 `json:"alliance_tax_rate,omitempty"`           /* Only present if alliance access is allowed */
	AllowAccessWithStandings bool    `json:"allow_access_with_standings,omitempty"` /* standing_level and any standing related tax rate only present when this is true */
	AllowAllianceAccess      bool    `json:"allow_alliance_access,omitempty"`       /* allow_alliance_access boolean */
	BadStandingTaxRate       float32 `json:"bad_standing_tax_rate,omitempty"`       /* bad_standing_tax_rate number */
	CorporationTaxRate       float32 `json:"corporation_tax_rate,omitempty"`        /* corporation_tax_rate number */
	ExcellentStandingTaxRate float32 `json:"excellent_standing_tax_rate,omitempty"` /* Tax rate for entities with excellent level of standing, only present if this level is allowed, same for all other standing related tax rates */
	GoodStandingTaxRate      float32 `json:"good_standing_tax_rate,omitempty"`      /* good_standing_tax_rate number */
	NeutralStandingTaxRate   float32 `json:"neutral_standing_tax_rate,omitempty"`   /* neutral_standing_tax_rate number */
	OfficeId                 int64   `json:"office_id,omitempty"`                   /* unique ID of this customs office */
	ReinforceExitEnd         int32   `json:"reinforce_exit_end,omitempty"`          /* reinforce_exit_end integer */
	ReinforceExitStart       int32   `json:"reinforce_exit_start,omitempty"`        /* Together with reinforce_exit_end, marks a 2-hour period where this customs office could exit reinforcement mode during the day after initial attack */
	StandingLevel            string  `json:"standing_level,omitempty"`              /* Access is allowed only for entities with this level of standing or better */
	SystemId                 int32   `json:"system_id,omitempty"`                   /* ID of the solar system this customs office is located in */
	TerribleStandingTaxRate  float32 `json:"terrible_standing_tax_rate,omitempty"`  /* terrible_standing_tax_rate number */
}
