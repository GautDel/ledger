package handlers

import (
	"encoding/gob"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"ledgerbolt.systems/internal/middleware"
)

func New() *gin.Engine {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://192.168.1.15:4200"}, // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin","Cache-Control", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins here or check against a whitelist
			return true
		},
		MaxAge: 12 * time.Hour,
	}
	router.Use(cors.New(config))

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Use(middleware.EnsureValidToken())
	router.GET("/", homeHandler)

	clientRouter := router.Group("/clients")
	{
        clientRouter.GET("/sort/:sort", getClientsHandler)
		clientRouter.POST("/search", searchClientsHandler)
		clientRouter.POST("/create", newClientHandler)
		clientRouter.PUT("/update/:id", updateClientHandler)
		clientRouter.PUT("/star/:id", starClientHandler)
		clientRouter.DELETE("/remove/:id", destroyClientHandler)
		clientRouter.GET("/:id", getClientHandler)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/", GetUserHandler)
		userRouter.PUT("/update", UpdateUserHandler)
	}

	payStatRouter := router.Group("/payment-status")
	{
		payStatRouter.GET("/", getPaymentStatus)
		payStatRouter.POST("/create", createPaymentStatus)
		payStatRouter.PUT("/update/:id", updatePaymentStatus)
		payStatRouter.DELETE("/remove/:id", destroyPaymentStatus)
		payStatRouter.GET("/:id", getSinglePaymentStatus)
	}

	projectRouter := router.Group("/projects")
	{
        projectRouter.GET("/sort/:sort", getProjects)
		projectRouter.POST("/create", createProject)
		projectRouter.PUT("/update/:id", updateProject)
		projectRouter.DELETE("/remove/:id", destroyProject)
		projectRouter.GET("/:id", getProject)

		projectRouter.GET("/client/:id", getProjectByClient)
	}

	invoiceRouter := router.Group("/invoices")
	{
		invoiceRouter.GET("/", getInvoices)
		invoiceRouter.POST("/create", createInvoice)
		invoiceRouter.PUT("/update/:id", updateInvoice)
		invoiceRouter.DELETE("/remove/:id", destroyInvoice)
		invoiceRouter.GET("/:id", getInvoice)
	}

	bankRouter := router.Group("/bank")
	{
		bankRouter.GET("/", getBanks)
		bankRouter.POST("/create", createBank)
		bankRouter.PUT("/update/:id", updateBank)
		bankRouter.DELETE("/remove/:id", destroyBank)
		bankRouter.GET("/:id", getBank)
	}

	serviceRouter := router.Group("/services")
	{
		serviceRouter.GET("/", getServices)
		serviceRouter.POST("/create", createService)
		serviceRouter.PUT("/update/:id", updateService)
		serviceRouter.DELETE("/remove/:id", destroyService)
		serviceRouter.GET("/:id", getService)
	}

	return router
}
