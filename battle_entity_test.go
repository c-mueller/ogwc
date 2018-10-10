package ogwc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BattleEntity_Config_Loading(t *testing.T) {
	cnt := 0
	for _, v := range Entities {
		fmt.Println(v)
		cnt++
	}

	assert.True(t, cnt > 10)
}
