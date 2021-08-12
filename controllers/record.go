// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"encoding/json"

	"github.com/casbin/casdoor/object"
)

// GetRecords
// @Title GetRecords
// @Description get all records
// @Success 200 {array} object.Records The Response object
// @router /get-records [get]
func (c *ApiController) GetRecords() {
	c.Data["json"] = object.GetRecords()
	c.ServeJSON()
}

// GetRecordsByFilter
// @Title GetRecordsByFilter
// @Description get records by filter
// @Param   body    body   object.Records  true  "filter Record message"
// @Success 200 {array} object.Records The Response object
// @router /get-records-filter [post]
func (c *ApiController) GetRecordsByFilter() {
	var record object.Records
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &record)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = object.GetRecordsByField(&record)
	c.ServeJSON()
}