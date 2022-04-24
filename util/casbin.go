package util

import (
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func NewCasbinEnforcer() *casbin.Enforcer {
	// 从CSV文件adapter加载策略规则
	// 使用自己的 adapter 替换。
	a := fileadapter.NewAdapter("config/policy.csv")

	// 创建enforcer
	e, err := casbin.NewEnforcer("config/model.conf", a)
	if err != nil {
		NewLogger().Fatal("load file failed, %v", err.Error())
	}

	return e
}

func CheckPermission(ctx *gin.Context, sub, obj, act string) {
	logger := NewLogger()
	logger.Info("sub = %s obj = %s act = %s", sub, obj, act)
	e := NewCasbinEnforcer()
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

func GetRoles(ctx *gin.Context, username string) {
	e := NewCasbinEnforcer()
	roles, err := e.GetRolesForUser(username)
	if err != nil {
		NewLogger().Fatal("load file failed, %v", err.Error())
	}

	ctx.JSON(http.StatusOK, roles)
}
