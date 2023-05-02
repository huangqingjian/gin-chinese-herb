package middleware

import (
	"gin-first/config"
	"github.com/gin-contrib/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"net/http"
)

// 注册中间件
func RegistMiddleWare(e *gin.Engine) {
	e.Use(Cors(), Session(), Auth(), i18n.Localize(i18n.WithBundle(&i18n.BundleCfg{
		RootPath:         "../conf/locale",
		AcceptLanguage:   []language.Tag{language.Chinese, language.English},
		DefaultLanguage:  language.Chinese,
		FormatBundleFile: "yaml",
		UnmarshalFunc:    yaml.Unmarshal,
	}), i18n.WithGetLngHandle(func(context *gin.Context, defaultLang string) string {
		if context == nil || context.Request == nil {
			return defaultLang
		}
		lang := context.Query("lang")
		if lang != "" {
			return lang
		}
		lang = context.GetHeader("Accept-Language")
		if lang != "" {
			return lang
		}
		return defaultLang
	})))
}

// session
func Session() gin.HandlerFunc {
	store := cookie.NewStore([]byte(config.Global.Session.Secret))
	return sessions.Sessions(config.Global.Session.Id, store)
}

// auth校验（登录态等）
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//userId := sessions.Default(c).Get(constant.USER_ID)
		//if userId == nil {
		//	panic(any("请登录后再访问"))
		//
		//}
		c.Next()
	}
}

// 支持跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
