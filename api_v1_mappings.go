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

package ogwc

import "github.com/gin-gonic/gin"

func (a *OGWCApplication) registerV1ApiMappings() {
	v1Api := a.engine.Group("/api/v1")
	a.registerV1APIRootMappings(v1Api)

	v1CalcApi := v1Api.Group("/calculation/:id")
	a.registerV1APIGenericCalculationMappings(v1CalcApi)

	v1ParticipantApi := v1CalcApi.Group("/participant")
	a.registerV1APIParticipantMappings(v1ParticipantApi)

}

func (a *OGWCApplication) registerV1APIParticipantMappings(v1ParticipantApi *gin.RouterGroup) {
	v1ParticipantApi.POST("/fleet-loss", a.updateAdditionalFleetLoss)
	v1ParticipantApi.POST("/resource-loss", a.updateAdditionalResourceLoss)
	v1ParticipantApi.POST("/add", a.addParticipant)
	v1ParticipantApi.POST("/delete", a.deleteParticipant)
	v1ParticipantApi.POST("/win/percentage", a.updateWinPercentageOfParticipant)
	v1ParticipantApi.POST("/win/fixed", a.updateFixedWinOfParticipant)
	v1ParticipantApi.POST("/win/none", a.updateDisableWinOfParticipant)
}

func (a *OGWCApplication) registerV1APIGenericCalculationMappings(v1CalcApi *gin.RouterGroup) {
	v1CalcApi.GET("/", a.getCalculation)
	v1CalcApi.GET("/report", a.getReport)
	v1CalcApi.GET("/report/transfers", a.getTransfers)
	v1CalcApi.POST("/add/:key", a.addKey)
	v1CalcApi.POST("/rebalance-win", a.rebalancePercentage)
}

func (a *OGWCApplication) registerV1APIRootMappings(v1Api *gin.RouterGroup) {
	v1Api.GET("/version", a.getVersionInfo)
	v1Api.POST("/submit/:key", a.newCalculation)
}

func (a *OGWCApplication) getVersionInfo(ctx *gin.Context) {
	ctx.JSON(200, a.versionInfo)
}