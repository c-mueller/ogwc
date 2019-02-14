// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2019 Christian Müller <dev@c-mueller.xyz>.
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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

type OGameAPI interface {
	GetCombatReport(key string) (*CombatReport, error)
	GetHarvestReport(key string) (*HarvestReport, error)
	GetMissileReport(key string) (*MissileReport, error)
}

type OGAPIRestAPI struct {
	QueryURL string
}

func (a OGAPIRestAPI) GetCombatReport(key string) (*CombatReport, error) {
	url := fmt.Sprintf(a.QueryURL, key)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cr CombatReport
	err = json.Unmarshal(data, &cr)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

func (a OGAPIRestAPI) GetMissileReport(key string) (*MissileReport, error) {
	url := fmt.Sprintf(a.QueryURL, key)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var mr MissileReport
	err = json.Unmarshal(data, &mr)
	if err != nil {
		return nil, err
	}
	return &mr, nil
}

func (OGAPIRestAPI) GetHarvestReport(key string) (*HarvestReport, error) {
	url := fmt.Sprintf("https://ogapi.rest/v1/report/%s/0", key)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cr HarvestReport
	err = json.Unmarshal(data, &cr)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}
