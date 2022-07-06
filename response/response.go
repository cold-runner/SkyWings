package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SkyWingsNetTest(code RespCode, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  code.msg(),
	})
}
func SkyWingsNotExist(code RespCode, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  code.msg(),
	})
}

func SkyWingsSignUpSuccess(code RespCode, ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  code.msg(),
		//"data": data,
	})
}
func SkyWingsValidateFailed(code RespCode, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": CodeSignupValidateFail,
		"Msg":  code.msg(),
	})
}
func SkyWingsSignUpFailed(code RespCode, ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": code,
		"Msg":  code.msg(),
	})
}
func SkyWingsSignUpExist(code RespCode, ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"Code": CodeStudentExist,
		"Msg":  code.msg(),
		//"data": data,
	})
}
