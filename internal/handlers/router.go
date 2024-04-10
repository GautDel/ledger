package handlers

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/middleware"
)

// New registers the routes and returns the router.
func New(auth *auth.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

    authRouter := router.Group("/auth", middleware.IsAuthenticated)
	{
		authRouter.GET("/", homeHandler)
		authRouter.GET("/login", loginHandler(auth))
		authRouter.GET("/logout", logoutHandler)
		authRouter.GET("/callback", callbackHandler(auth))
	}

	return router
}
