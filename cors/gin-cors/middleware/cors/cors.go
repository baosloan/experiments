package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		//指定允许跨域请求的源。例如: []string{"http://example.com", "https://another.com"}，使用"*"表示允许所有源
		AllowOrigins: []string{"*"},
		//指定允许的HTTP方法
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		//指定允许的HTTP头
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token", "Cookie", "Accept-Language", "Client"},
		//指定可以暴露给浏览器的头。例如：[]string{"X-Custom-Header", "Content-Length"}
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//指定是否允许发送凭证（如 cookies）。默认为 false
		AllowCredentials: true,
		//指定是否允许 WebSocket 请求，默认为 false
		AllowWebSockets: true,
		//指定是否允许使用通配符 * 来匹配域名。例如：如果设置为 true，AllowOrigins 中的 http://*.example.com 将匹配 http://foo.example.com。
		AllowWildcard: true,
		//一个时间段，指定浏览器在发起预检请求（OPTIONS）后可以缓存该请求的时间。
		MaxAge: 12 * time.Hour,
	})
}
