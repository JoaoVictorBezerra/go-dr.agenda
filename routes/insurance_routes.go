package routes

import (
	"dr.agenda/handler"
	"github.com/gin-gonic/gin"
)

func RegisterInsuranceRoutes(r *gin.Engine, handler *handler.InsuranceHandler) {
	insurance := r.Group("/api/insurance")
	{
		insurance.GET("", handler.GetInsurances)
		insurance.GET("/:id", handler.GetInsuranceById)
		insurance.POST("", handler.CreateInsurance)
		insurance.PUT("/:id", handler.UpdateInsurance)
		insurance.DELETE("/:id", handler.SuspendInsurance)
	}
}
