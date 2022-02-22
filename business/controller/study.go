package controller

import (
	"github.com/gin-gonic/gin"
)

type study struct{}

func NewStudy() *study {
	return &study{}
}

func (this *study) Sort(ctx *gin.Context ) {
	ctx.JSON(200, gin.H{"msg": "hello study controller"})
}
