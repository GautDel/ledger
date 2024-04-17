package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ledgerbolt.systems/internal/auth"
	"ledgerbolt.systems/internal/db"
	"ledgerbolt.systems/internal/models"
	"ledgerbolt.systems/internal/validator"
)

type ProjectRB struct {
	Name        string `json:"Name" validate:"required,min=3,max=50"`
	Description string `json:"Description" validate:"required,min=3,max=1000"`
	ClientID    int    `json:"ClientID" validate:"required"`
	Notes       string `json:"Notes"`
}

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
	idStr := ctx.Param("id")
	clientID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to read param",
				"error":   err.Error(),
			},
		)
		return
	}

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
	var reqBody ProjectRB

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

	project := models.Project{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		ClientID:    reqBody.ClientID,
		Notes:       reqBody.Notes,
	}

	err = models.CreateProject(conn, project, ctx, auth.GetUser(ctx))
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
	var reqBody ProjectRB
	idStr := ctx.Param("id")
	pID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to read param",
				"error":   err.Error(),
			},
		)
		return
	}

	err = ctx.ShouldBindJSON(&reqBody)
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

	project := models.Project{
		ID:          pID,
		Name:        reqBody.Name,
		Description: reqBody.Description,
		ClientID:    reqBody.ClientID,
		Notes:       reqBody.Notes,
	}

	err = models.UpdateProject(conn, project, ctx, auth.GetUser(ctx))
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
	idStr := ctx.Param("id")
	pID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Failed to read param",
				"error":   err.Error(),
			},
		)

		return
	}

	err = models.DestroyProject(conn, ctx, pID, auth.GetUser(ctx))
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
