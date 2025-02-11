package handlers

import (
	"github.com/gin-gonic/gin"
)

// LGPayload expected payload from requesr
type LGPayload struct {
	Src     string `form:"src" json:"src"`
	Proto   int    `form:"proto" json:"proto"`
	Dst     string `form:"dst" json:"dst"`
	Captcha string `form:"captcha" json:"captcha"`
}

// addRoute helps to declare all applcaition related routes
func (r routes) addRoute(router *gin.RouterGroup) {
	lg := router.Group("/lg")
	lg.GET("/src", getSrcHost)
	lg.GET("/protocol", getProtocol)
	// lg.POST("/submit", getResult)

	ip := router.Group("/ip")
	ip.POST("", addIPInfo)
	ip.GET("", getIPInfo)
	ip.GET("/:ip", getIPInfo)
	ip.POST("/:ip", updateIPInfo)
	ip.DELETE("/:ip", deleteIPInfo)
	ip.GET("/vendor", getVendor)

	access := router.Group("/audit")
	access.GET("", getAccessData)
	access.POST("/:t", setAccessData)
}

// addRoute helps to declare all applcaition related routes
func (r routes) addWSRoute(router *gin.RouterGroup) {
	router.GET("/query", wsHandler)
}
