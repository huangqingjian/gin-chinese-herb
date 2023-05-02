package test

import (
	"encoding/json"
	"gin-first/config"
	"gin-first/db"
	"gin-first/logger"
	"gin-first/middleware"
	"gin-first/redis"
	"gin-first/response"
	"gin-first/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

// 初始化
func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	config.InitConfigV2(apppath + string(filepath.Separator) + "conf" + string(filepath.Separator) + "app.ini")
	logger.InitLogger()
	redis.InitRedis()
	db.InitDB()
}

// 初始化router
func initRouter() *gin.Engine {
	e := gin.New()
	e.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册中间件
	middleware.RegistMiddleWare(e)
	// 添加路由
	router.RegistRouter(e)
	return e
}

// 查询banner测试
func TestGetBannerList(t *testing.T) {
	router := initRouter()
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/api/banner/list", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	var response = response.Response{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 0, response.Code)
}
