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

package ogwc

import (
	"github.com/c-mueller/ogwc/core"
	"github.com/gin-gonic/gin"
)

func (a *OGWCApplication) updateWithErrorHandling(id string, calc *core.CombatReportCalculation, ctx *gin.Context) {
	err := a.repo.Update(id, *calc)
	if err != nil {
		ctx.JSON(500, errorResponse{
			Code:    500,
			Message: "Failed to Contact Database. This is not your fault!",
		})
	}
}

func (a *OGWCApplication) getParticipantNameQueryParameter(ctx *gin.Context, calc *core.CombatReportCalculation) string {
	name := ctx.Query("name")
	if !calc.Participants.IsPresent(name) {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "Participant name not found!",
		})
		return ""
	}
	return name
}

func (a *OGWCApplication) getCalculationFromContext(ctx *gin.Context) (string, *core.CombatReportCalculation) {
	id := ctx.Param("id")
	calc := a.repo.Get(id)
	if calc == nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: "Calculation not found",
		})
		return id, nil
	}
	return id, calc
}