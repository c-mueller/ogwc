// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ogwc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (a *OGWCApplication) addKey(ctx *gin.Context) {
	id := ctx.Param("id")
	key := ctx.Param("key")

	calc := a.repo.Get(id)
	if calc == nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: "Calculation not found",
		})
		return
	}

	if rrRegex.Match([]byte(key)) {
		report, err := a.api.GetHarvestReport(key)
		if err != nil {
			ctx.JSON(404, errorResponse{
				Code:    404,
				Message: fmt.Sprintf("Fetching the API Key %q failed. Error Message: %q", key, err.Error()),
			})
			return
		}

		calc.AddHarvestReport(*report)

	} else if crRegex.Match([]byte(key)) {
		report, err := a.api.GetCombatReport(key)
		if err != nil {
			ctx.JSON(404, errorResponse{
				Code:    404,
				Message: fmt.Sprintf("Fetching the API Key %q failed. Error Message: %q", key, err.Error()),
			})
			return
		}

		calc.AddCombatReport(*report, true)
	} else {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "Invalid API Key!",
		})
		return
	}

	a.repo.Update(id, *calc)
}

func (a *OGWCApplication) getReport(ctx *gin.Context) {
	id := ctx.Param("id")

	calc := a.repo.Get(id)
	if calc == nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: "Calculation not found",
		})
		return
	}

	ctx.JSON(200, calc.GetReport())
}

func (a *OGWCApplication) getCalculation(ctx *gin.Context) {
	id := ctx.Param("id")

	calc := a.repo.Get(id)
	if calc == nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: "Calculation not found",
		})
		return
	}

	ctx.JSON(200, calc)
}

func (a *OGWCApplication) newCalculation(ctx *gin.Context) {
	param := ctx.Param("key")
	if !crRegex.Match([]byte(param)) {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "Invalid API Key",
		})
		return
	}

	cr, err := a.api.GetCombatReport(param)
	if err != nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: fmt.Sprintf("Fetching the API Key %q failed. Error Message: %q", param, err.Error()),
		})
		return
	}

	calculation := cr.ToReportCalculation(true)
	uid := a.repo.Insert(*calculation)

	ctx.JSON(201, calculationCreationResponse{
		Code:          201,
		CalculationID: uid,
	})
}