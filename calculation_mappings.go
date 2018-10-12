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

package ogwc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type calculationCreationResponse struct {
	Code           int    `json:"code"`
	CalculationID  string `json:"calculation_id"`
	CalculationURL string `json:"calculation_url"`
}

func (a *OGWCApplication) getReport(ctx *gin.Context) {
	_, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	ctx.JSON(200, calc.GetReport())
}

func (a *OGWCApplication) getCalculation(ctx *gin.Context) {
	_, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
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
		Code:           201,
		CalculationID:  uid,
		// TODO Implement proper scheme detection
		CalculationURL: fmt.Sprintf("http://%s/api/v1/calculation/%s",  ctx.Request.Host, uid),
	})
}
