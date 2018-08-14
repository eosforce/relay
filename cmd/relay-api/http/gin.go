package http

import "github.com/gin-gonic/gin"

// Router gin root router
var Router *gin.Engine

// InitRouter init router
func InitRouter(isDebug bool) {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	Router = gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	Router.Use(logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	Router.Use(gin.Recovery())

	return
}
