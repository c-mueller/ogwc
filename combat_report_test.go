package ogwc

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

	calc := cr.ToReportCalculation(true)

	h2 := parseSampleHR_2(t)

	calc.AddHarvestReport(h2)

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

	calc.GetReport()

	data, _ := json.Marshal(calc)

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
