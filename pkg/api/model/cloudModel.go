// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>

package model

//GetUserToken GetUserToken
//swagger:parameters getToken
type GetUserToken struct {
	// in: body
	Body struct {
		//eid
		//in: body
		//required: true
		EID string `json:"eid" validate:"eid|between:30,33"`
		//可控范围
		//in: body
		//required: false
		Range string `json:"range" validate:"range|in:source,node,all"`
		//有效期
		//in: body
		//required: false
		ValidityPeriod int `json:"validity_period" validate:"validity_period"` //20191120
		//数据中心标识
		//in: body
		//required: true
		RegionTag  string `json:"region_tag" validate:"region_tag"`
		BeforeTime int    `json:"before_time"`
	}
}

//TokenInfo TokenInfo
type TokenInfo struct {
	EID   string `json:"eid"`
	Token string `json:"token"`
	CA    string `json:"ca"`
	Key   string `json:"key"`
}