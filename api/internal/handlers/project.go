package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
	"ledgerbolt.systems/internal/validator"
)

func getProjects(ctx *gin.Context) {
	conn := db.GetPool()

	projects, err := models.GetProjects(conn, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Failed to get projects from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, projects)
}

func getProject(ctx *gin.Context) {
	conn := db.GetPool()
	projectID := ctx.Param("id")

	project, err := models.GetProject(conn, ctx, auth.GetUser(ctx), projectID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Failed to get project from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func getProjectByClient(ctx *gin.Context) {
	conn := db.GetPool()
	clientID := ctx.Param("id")

	project, err := models.GetProjectByClient(conn, ctx, auth.GetUser(ctx), clientID)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Failed to get projects from database",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func createProject(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.Project

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to read request body",
				"error":   err.Error(),
			},
		)
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to create project",
				"error":   err.Error(),
			},
		)
		return
	}

	err = models.CreateProject(conn, reqBody, ctx, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, 
            gin.H{
                "message": "Failed to create project", 
                "error": err.Error(),
            },
        )
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project successfully created!"})

}

func updateProject(ctx *gin.Context) {
	conn := db.GetPool()
	var reqBody models.Project
	pID := ctx.Param("id")

    err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to read request body",
				"error":   err.Error(),
			},
		)
		return
	}

	err = validator.Validate(&reqBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to create project",
				"error":   err.Error(),
			},
		)
		return
	}

	err = models.UpdateProject(conn, reqBody, ctx, pID, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Failed to create project",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project successfully updated!"})
}

func destroyProject(ctx *gin.Context) {
	conn := db.GetPool()
	pID := ctx.Param("id")

    err := models.DestroyProject(conn, ctx, pID, auth.GetUser(ctx))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Failed to remove project",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project successfully removed!"})
}
