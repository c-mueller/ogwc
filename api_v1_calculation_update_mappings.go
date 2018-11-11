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
	"github.com/c-mueller/ogwc/core"
	"github.com/gin-gonic/gin"
	"strconv"
)

type additionalFleetLossRequest struct {
	Name      string     `json:"name"`
	LostFleet core.Fleet `json:"lost_fleet"`
}

type additionalResourceLossRequest struct {
	Name          string         `json:"name"`
	LostResources core.Resources `json:"lost_resources"`
}

func (a *OGWCApplication) rebalancePercentage(ctx *gin.Context) {
	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	calc.RebalanceDistributionPercentage()

	a.updateWithErrorHandling(id, calc, ctx)
}

func (a *OGWCApplication) updateWinPercentageOfParticipant(ctx *gin.Context) {
	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	name := a.getParticipantNameQueryParameter(ctx, calc)
	if len(name) == 0 {
		return
	}

	percentage := ctx.Query("percentage")

	percFloat, err := strconv.ParseFloat(percentage, 64)
	if err != nil {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "Invalid Percentage",
		})
		return
	}

	idx, p := calc.Participants.Find(name)

	p.DistribuitonMode = core.PERCENTAGE
	p.WinPercentage = percFloat
	p.FixedResourceAmount = nil

	calc.Participants[idx] = *p

	a.updateWithErrorHandling(id, calc, ctx)
}

func (a *OGWCApplication) updateFixedWinOfParticipant(ctx *gin.Context) {
	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	name := a.getParticipantNameQueryParameter(ctx, calc)
	if len(name) == 0 {
		return
	}

	var res core.Resources

	err := ctx.BindJSON(&res)
	if err != nil {
		return
	}

	idx, p := calc.Participants.Find(name)

	p.DistribuitonMode = core.FIXED_AMOUNT
	p.WinPercentage = 0
	p.FixedResourceAmount = &res

	calc.Participants[idx] = *p

	a.updateWithErrorHandling(id, calc, ctx)

}

func (a *OGWCApplication) updateDisableWinOfParticipant(ctx *gin.Context) {
	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	name := a.getParticipantNameQueryParameter(ctx, calc)
	if len(name) == 0 {
		return
	}

	idx, p := calc.Participants.Find(name)

	p.DistribuitonMode = core.NONE
	p.WinPercentage = 0
	p.FixedResourceAmount = nil

	calc.Participants[idx] = *p

	a.updateWithErrorHandling(id, calc, ctx)

}

func (a *OGWCApplication) deleteParticipant(ctx *gin.Context) {
	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	name := a.getParticipantNameQueryParameter(ctx, calc)
	if len(name) == 0 {
		return
	}

	if !calc.InitialFleet[name].IsZero() {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "This Participant cannot be removed.",
		})
		return
	}

	idx, _ := calc.Participants.Find(name)

	calc.Participants = append(calc.Participants[:idx], calc.Participants[(idx + 1):]...)

	a.updateWithErrorHandling(id, calc, ctx)
}

func (a *OGWCApplication) addParticipant(ctx *gin.Context) {
	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	name := ctx.Query("name")
	if len(name) == 0 {
		return
	}

	if calc.Participants.IsPresent(name) {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "The Participant name has to be unique",
		})
		return
	}

	calc.Participants = append(calc.Participants, core.Participant{
		Name:             name,
		DistribuitonMode: core.NONE,
	})

	a.updateWithErrorHandling(id, calc, ctx)
}

func (a *OGWCApplication) updateAdditionalFleetLoss(ctx *gin.Context) {
	var requestData additionalFleetLossRequest

	if err := ctx.BindJSON(&requestData); err != nil {
		log.Error(err.Error())
		return
	}

	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	idx, participant := calc.Participants.Find(requestData.Name)

	if participant == nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: fmt.Sprintf("Participant with name %q not found", requestData.Name),
		})
		return
	}

	participant.SetFleetLoss(requestData.LostFleet)

	calc.Participants[idx] = *participant

	a.updateWithErrorHandling(id, calc, ctx)
}

func (a *OGWCApplication) updateAdditionalResourceLoss(ctx *gin.Context) {
	var requestData additionalResourceLossRequest

	if err := ctx.BindJSON(&requestData); err != nil {
		log.Error(err.Error())
		return
	}

	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
		return
	}

	idx, participant := calc.Participants.Find(requestData.Name)

	if participant == nil {
		ctx.JSON(404, errorResponse{
			Code:    404,
			Message: fmt.Sprintf("Participant with name %q not found", requestData.Name),
		})
		return
	}

	participant.SetResourceLoss(requestData.LostResources)

	calc.Participants[idx] = *participant

	a.updateWithErrorHandling(id, calc, ctx)
}

func (a *OGWCApplication) addKey(ctx *gin.Context) {
	key := ctx.Param("key")

	id, calc := a.getCalculationFromContext(ctx)
	if calc == nil {
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
	} else if mrRegex.Match([]byte(key)) {
		report, err := a.api.GetMissileReport(key)
		if err != nil {
			ctx.JSON(404, errorResponse{
				Code:    404,
				Message: fmt.Sprintf("Fetching the API Key %q failed. Error Message: %q", key, err.Error()),
			})
			return
		}

		calc.AddMissileReport(*report)
	} else {
		ctx.JSON(400, errorResponse{
			Code:    400,
			Message: "Invalid API Key!",
		})
		return
	}

	a.updateWithErrorHandling(id, calc, ctx)
}
