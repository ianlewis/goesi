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

/* A list of GetCharactersCharacterIdStatsTravel. */
//easyjson:json
type GetCharactersCharacterIdStatsTravelList []GetCharactersCharacterIdStatsTravel

/* travel object */
//easyjson:json
type GetCharactersCharacterIdStatsTravel struct {
	AccelerationGateActivations int64 `json:"acceleration_gate_activations,omitempty"` /* acceleration_gate_activations integer */
	AlignTo                     int64 `json:"align_to,omitempty"`                      /* align_to integer */
	DistanceWarpedHighSec       int64 `json:"distance_warped_high_sec,omitempty"`      /* distance_warped_high_sec integer */
	DistanceWarpedLowSec        int64 `json:"distance_warped_low_sec,omitempty"`       /* distance_warped_low_sec integer */
	DistanceWarpedNullSec       int64 `json:"distance_warped_null_sec,omitempty"`      /* distance_warped_null_sec integer */
	DistanceWarpedWormhole      int64 `json:"distance_warped_wormhole,omitempty"`      /* distance_warped_wormhole integer */
	DocksHighSec                int64 `json:"docks_high_sec,omitempty"`                /* docks_high_sec integer */
	DocksLowSec                 int64 `json:"docks_low_sec,omitempty"`                 /* docks_low_sec integer */
	DocksNullSec                int64 `json:"docks_null_sec,omitempty"`                /* docks_null_sec integer */
	JumpsStargateHighSec        int64 `json:"jumps_stargate_high_sec,omitempty"`       /* jumps_stargate_high_sec integer */
	JumpsStargateLowSec         int64 `json:"jumps_stargate_low_sec,omitempty"`        /* jumps_stargate_low_sec integer */
	JumpsStargateNullSec        int64 `json:"jumps_stargate_null_sec,omitempty"`       /* jumps_stargate_null_sec integer */
	JumpsWormhole               int64 `json:"jumps_wormhole,omitempty"`                /* jumps_wormhole integer */
	WarpsHighSec                int64 `json:"warps_high_sec,omitempty"`                /* warps_high_sec integer */
	WarpsLowSec                 int64 `json:"warps_low_sec,omitempty"`                 /* warps_low_sec integer */
	WarpsNullSec                int64 `json:"warps_null_sec,omitempty"`                /* warps_null_sec integer */
	WarpsToBookmark             int64 `json:"warps_to_bookmark,omitempty"`             /* warps_to_bookmark integer */
	WarpsToCelestial            int64 `json:"warps_to_celestial,omitempty"`            /* warps_to_celestial integer */
	WarpsToFleetMember          int64 `json:"warps_to_fleet_member,omitempty"`         /* warps_to_fleet_member integer */
	WarpsToScanResult           int64 `json:"warps_to_scan_result,omitempty"`          /* warps_to_scan_result integer */
	WarpsWormhole               int64 `json:"warps_wormhole,omitempty"`                /* warps_wormhole integer */
}
