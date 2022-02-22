package controller

import (
	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestStudy_Sort(t *testing.T) {
	NewStudy().Sort(&gin.Context{})
}
