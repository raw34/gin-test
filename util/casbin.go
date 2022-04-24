package util

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CheckPermission(ctx *gin.Context, sub, obj, act string) {
	logger := NewLogger()
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		logger.Fatal("load file failed, %v", err.Error())
	}
	logger.Info("sub = %s obj = %s act = %s", sub, obj, act)
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		logger.Error("enforce failed %s", err.Error())
		ctx.String(http.StatusInternalServerError, "内部服务器错误")
		return
	}
	if !ok {
		logger.Info("权限验证不通过")
		ctx.String(http.StatusOK, "权限验证不通过")
		return
	}
	logger.Info("权限验证通过")
	ctx.String(http.StatusOK, "权限验证通过")
}
