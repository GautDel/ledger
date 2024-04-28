package handlers

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/middleware"
)

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

	clientRouter := router.Group("/clients", middleware.IsAuthenticated)
	{
		clientRouter.GET("/", getClientsHandler)
		clientRouter.GET("/search", searchClientsHandler)
		clientRouter.POST("/create", newClientHandler)
		clientRouter.PUT("/update/:id", updateClientHandler)
		clientRouter.DELETE("/remove/:id", destroyClientHandler)
		clientRouter.GET("/:id", getClientHandler)
	}

	userRouter := router.Group("/user", middleware.IsAuthenticated)
	{
		userRouter.GET("/", GetUserHandler)
		userRouter.PUT("/update", UpdateUserHandler)
	}

	payStatRouter := router.Group("/payment-status", middleware.IsAuthenticated)
	{
		payStatRouter.GET("/", getPaymentStatus)
        payStatRouter.POST("/create", createPaymentStatus)
        payStatRouter.PUT("/update/:id", updatePaymentStatus)
        payStatRouter.DELETE("/remove/:id", destroyPaymentStatus)
        payStatRouter.GET("/:id", getSinglePaymentStatus)
	}

	projectRouter := router.Group("/projects", middleware.IsAuthenticated)
	{
		projectRouter.GET("/", getProjects)
        projectRouter.POST("/create", createProject)
        projectRouter.PUT("/update/:id", updateProject)
        projectRouter.DELETE("/remove/:id", destroyProject)
        projectRouter.GET("/:id", getProject)

        projectRouter.GET("/client/:id", getProjectByClient)
	}

	invoiceRouter := router.Group("/invoices", middleware.IsAuthenticated)
	{
		invoiceRouter.GET("/", getInvoices)
        invoiceRouter.POST("/create", createInvoice)
        invoiceRouter.PUT("/update/:id", updateInvoice)
        invoiceRouter.DELETE("/remove/:id", destroyInvoice)
        invoiceRouter.GET("/:id", getInvoice)
	}

	bankRouter := router.Group("/bank", middleware.IsAuthenticated)
	{
        bankRouter.GET("/", getBanks)
        bankRouter.POST("/create", createBank)
        bankRouter.PUT("/update/:id", updateBank)
        bankRouter.DELETE("/remove/:id", destroyBank)
        bankRouter.GET("/:id", getBank)
	}

	return router
}
