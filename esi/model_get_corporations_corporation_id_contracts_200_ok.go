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

/* A list of GetCorporationsCorporationIdContracts200Ok. */
//easyjson:json
type GetCorporationsCorporationIdContracts200OkList []GetCorporationsCorporationIdContracts200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdContracts200Ok struct {
	AcceptorId          int32     `json:"acceptor_id,omitempty"`           /* Who will accept the contract */
	AssigneeId          int32     `json:"assignee_id,omitempty"`           /* ID to whom the contract is assigned, can be corporation or character ID */
	Availability        string    `json:"availability,omitempty"`          /* To whom the contract is available */
	Buyout              float64   `json:"buyout,omitempty"`                /* Buyout price (for Auctions only) */
	Collateral          float64   `json:"collateral,omitempty"`            /* Collateral price (for Couriers only) */
	ContractId          int32     `json:"contract_id,omitempty"`           /* contract_id integer */
	DateAccepted        time.Time `json:"date_accepted,omitempty"`         /* Date of confirmation of contract */
	DateCompleted       time.Time `json:"date_completed,omitempty"`        /* Date of completed of contract */
	DateExpired         time.Time `json:"date_expired,omitempty"`          /* Expiration date of the contract */
	DateIssued          time.Time `json:"date_issued,omitempty"`           /* Сreation date of the contract */
	DaysToComplete      int32     `json:"days_to_complete,omitempty"`      /* Number of days to perform the contract */
	EndLocationId       int64     `json:"end_location_id,omitempty"`       /* End location ID (for Couriers contract) */
	ForCorporation      bool      `json:"for_corporation,omitempty"`       /* true if the contract was issued on behalf of the issuer's corporation */
	IssuerCorporationId int32     `json:"issuer_corporation_id,omitempty"` /* Character's corporation ID for the issuer */
	IssuerId            int32     `json:"issuer_id,omitempty"`             /* Character ID for the issuer */
	Price               float64   `json:"price,omitempty"`                 /* Price of contract (for ItemsExchange and Auctions) */
	Reward              float64   `json:"reward,omitempty"`                /* Remuneration for contract (for Couriers only) */
	StartLocationId     int64     `json:"start_location_id,omitempty"`     /* Start location ID (for Couriers contract) */
	Status              string    `json:"status,omitempty"`                /* Status of the the contract */
	Title               string    `json:"title,omitempty"`                 /* Title of the contract */
	Type_               string    `json:"type,omitempty"`                  /* Type of the contract */
	Volume              float64   `json:"volume,omitempty"`                /* Volume of items in the contract */
}
