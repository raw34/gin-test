package util

import (
	"net/http"
	"sync"

	"github.com/casbin/casbin/v2"

	entadapter "github.com/casbin/ent-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var syncedEnforcer *casbin.SyncedEnforcer
var once sync.Once

func NewCasbinEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		logger := NewLogger()
		a, err := entadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
		if err != nil {
			logger.Fatal("load mysql failed, %v", err.Error())
		}

		// 创建enforcer
		syncedEnforcer, err = casbin.NewSyncedEnforcer("config/model.conf", a)
		if err != nil {
			logger.Fatal("load file failed, %v", err.Error())
		}
		syncedEnforcer.EnableLog(true)
		syncedEnforcer.EnableAutoSave(true)

		// 加载策略
		err = syncedEnforcer.LoadPolicy()
		if err != nil {
			logger.Fatal("load policy failed, %v", err.Error())
		}
	})

	return syncedEnforcer
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
	logger := NewLogger()
	logger.Info("username = %s", username)
	e := NewCasbinEnforcer()
	roles, err := e.GetRolesForUser(username)
	if err != nil {
		logger.Fatal("load file failed, %v", err.Error())
	}

	ctx.JSON(http.StatusOK, roles)
}

func AddRoles(ctx *gin.Context, username string, roles []string) {
	logger := NewLogger()
	logger.Info("username = %s, roles = %s", username, roles)
	e := NewCasbinEnforcer()
	res, err := e.AddRolesForUser(username, roles)
	if err != nil {
		logger.Fatal("load file failed, %v", err.Error())
	}

	ctx.JSON(http.StatusOK, res)
}
