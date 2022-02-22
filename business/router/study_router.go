package router

import (
	"github.com/gin-gonic/gin"
	"study-go/business/controller"
)

func studyRouter(e *gin.Engine) {

	studyController := controller.NewStudy()
	e.GET("study/test", studyController.Sort)
}
