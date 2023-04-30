package logger

import (
	"gin-first/config"
	e "gin-first/error"
	"gin-first/response"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var (
	Logger *zap.Logger
)

// 初始化logger
func InitLogger() {
	fileWriteSyncer := getFileLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.InfoLevel)
	// zap打印时将整个调用的stack链路会存放到内存中，默认打印调用处的caller信息。所以需要再初始化zap时额外增加AddCallerSkip跳过指定层级的caller
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(Logger)
}

// 获取日志编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder
	time := config.Global.Zap.Time
	if time != "" {
		encoderConfig.TimeKey = time
	}
	 format := config.Global.Zap.Format
	if format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//
func getFileLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Global.Zap.FileName, 				// 日志文件位置
		MaxSize:    config.Global.Zap.MaxSize,                  // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     config.Global.Zap.MaxAge,                   // 保留旧文件的最大天数
		Compress:   config.Global.Zap.Compress,                 // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		Logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// 自定义recovery
func GinRecovery(stack bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := interface{}(recover()); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					Logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					Logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				Logger.Error("Recover Error", zap.Any("error", err))
				switch obj := err.(type) {
				case e.AuthError :
					c.JSON(http.StatusUnauthorized, response.Fail(obj.Code, obj.Message))
				case e.ParamError:
					c.JSON(http.StatusBadRequest, response.Fail(obj.Code, obj.Message))
				case e.ServiceError:
					c.JSON(http.StatusInternalServerError, response.Fail(obj.Code, obj.Message))
				default:
					c.JSON(http.StatusInternalServerError, response.FastFail("服务器内部异常~"))
				}
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}



