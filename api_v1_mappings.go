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