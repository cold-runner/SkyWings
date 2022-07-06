package controller

import (
	"SkyWings/dao/mysql"
	"SkyWings/logic"
	"SkyWings/models"
	resp "SkyWings/response"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NetPingPong(ctx *gin.Context) {
	resp.SkyWingsNetTest(resp.CodeSuccess, ctx)
}
func NoRoute(ctx *gin.Context) {
	resp.SkyWingsNotExist(resp.CodeNotExist, ctx)
}
func SignUpHandler(ctx *gin.Context) {
	stu := models.Student{}
	// 参数绑定并校验
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		// 参数校验失败
		zap.L().Error("参数校验失败！", zap.Error(err))
		resp.SkyWingsValidateFailed(resp.CodeSignupValidateFail, ctx)
		ctx.Abort()
		return
	}

	// 注册逻辑
	if err := logic.SignUp(&stu); err != nil {

		// 用户已存在
		if errors.Is(err, mysql.ErrorUserExit) {
			resp.SkyWingsSignUpExist(resp.CodeStudentExist, ctx, stu)
			return
		}

		// 数据库错误
		if errors.Is(err, mysql.ErrorDao) {
			resp.SkyWingsSignUpFailed(resp.CodeDaoError, ctx, err)
			return
		}

	}

	// 注册成功响应
	resp.SkyWingsSignUpSuccess(resp.CodeSuccess, ctx, stu)
}
