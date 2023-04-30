package session

import (
	"gin-first/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 从session中取值
func Get(c *gin.Context, key interface{}) interface{} {
	s := sessions.Default(c)
	return s.Get(key)
}

// 设置session值
func Set(c *gin.Context, key interface{}, val interface{}) {
	s := sessions.Default(c)
	s.Options(sessions.Options{MaxAge: config.Global.Session.MaxAge})
	s.Set(key, val)
	s.Save()
}

// 删除session
func Del(c *gin.Context) {
	s := sessions.Default(c)
	s.Options(sessions.Options{MaxAge: -1})
	s.Save()
}
