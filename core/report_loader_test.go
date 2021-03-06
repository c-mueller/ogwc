// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2020 Christian Müller <dev@c-mueller.xyz>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOGAPIRestAPI_GetReport(t *testing.T) {

	rs := `{"metal":25000000,"crystal":25000000,"deuterium":0,"id":"meddl"}`

	var res IdentifiableResources

	json.Unmarshal([]byte(rs), &res)

	ires := res.ToIdentifiableResources()

	s, _ := json.Marshal(ires)

	fmt.Println(string(s))
}
