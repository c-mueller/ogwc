// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
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
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_ToReportCalculation(t *testing.T) {
	cr := parseSampleCR(t)
	cr2 := parseSampleCR_2(t)
	cr3 := parseSampleCR_3(t)

	calc := cr.ToReportCalculation(true)

	h1 := parseSampleHR_1(t)
	h2 := parseSampleHR_2(t)
	h3 := parseSampleHR_3(t)

	calc.AddHarvestReport(h1)
	calc.AddHarvestReport(h2)
	calc.AddHarvestReport(h3)

	calc.AddCombatReport(cr2, true)
	calc.AddCombatReport(cr3, true)

	calc.AddAdditionalLossForParticipant("Chris", Fleet{
		Deathstar: 52,
	})
	calc.AddAdditionalLossForParticipant("Who", Fleet{
		Deathstar: 32,
	})

	calc.AddParticipant(Participant{
		Name:             "Scrop",
		DistribuitonMode: FIXED_AMOUNT,
		FixedResourceAmount: &Resources{
			Metal:   500000000,
			Crystal: 320000000,
		},
	})

	calc.AddParticipant(Participant{
		Name:             "Wuppi",
		DistribuitonMode: FIXED_AMOUNT,
		FixedResourceAmount: &Resources{
			Metal:   500000000,
			Crystal: 320000000,
		},
	})

	data, _ := json.MarshalIndent(calc.GetReport(), "", "  ")

	fmt.Println(string(data))
}

func Test_Parse_CombatReport(t *testing.T) {
	cr := parseSampleCR(t)

	assert.Equal(t, "6:264:10", cr.Generic.CombatCoordinates)
	assert.Equal(t, 3, len(cr.Attackers))
	assert.Equal(t, 1, len(cr.Defenders))
	assert.Equal(t, 4, len(cr.Rounds))
}

func parseSampleCR(t *testing.T) CombatReport {
	f, err := os.Open("testdata/combat_report.json")
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var cr CombatReport

	err = json.Unmarshal(data, &cr)
	assert.NoError(t, err)

	return cr
}

func parseSampleCR_2(t *testing.T) CombatReport {
	f, err := os.Open("testdata/combat_report_2.json")
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var cr CombatReport

	err = json.Unmarshal(data, &cr)
	assert.NoError(t, err)

	return cr
}
func parseSampleCR_3(t *testing.T) CombatReport {
	f, err := os.Open("testdata/combat_report_3.json")
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var cr CombatReport

	err = json.Unmarshal(data, &cr)
	assert.NoError(t, err)

	return cr
}

func parseSampleHR_1(t *testing.T) HarvestReport {
	f, err := os.Open("testdata/harvest_report_1.json")
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var cr HarvestReport

	err = json.Unmarshal(data, &cr)
	assert.NoError(t, err)

	return cr
}

func parseSampleHR_2(t *testing.T) HarvestReport {
	f, err := os.Open("testdata/harvest_report_2.json")
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var cr HarvestReport

	err = json.Unmarshal(data, &cr)
	assert.NoError(t, err)

	return cr
}
func parseSampleHR_3(t *testing.T) HarvestReport {
	f, err := os.Open("testdata/harvest_report_3.json")
	assert.NoError(t, err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	var cr HarvestReport

	err = json.Unmarshal(data, &cr)
	assert.NoError(t, err)

	return cr
}
