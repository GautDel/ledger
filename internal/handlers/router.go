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

	router.GET("/", middleware.IsAuthenticated, homeHandler)

	authRouter := router.Group("/auth")
	{
		authRouter.GET("/login", loginHandler(auth))
		authRouter.GET("/logout", logoutHandler)
		authRouter.GET("/callback", callbackHandler(auth))

	}

    clientRouter := router.Group("/clients")
    {
        clientRouter.GET("/", getClientsHandler)
        clientRouter.GET("/search", searchClientsHandler)
        clientRouter.POST("/create", newClientHandler)
        clientRouter.PUT("/update/:id", updateClientHandler)
        clientRouter.DELETE("/remove/:id", destroyClientHandler)
        clientRouter.GET("/:id", getClientHandler)
    }

	return router
}
