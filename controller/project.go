package controller

import (
	"github.com/Penetration-Platform-Go/Project-Service/model"
	"github.com/gin-gonic/gin"
)

// CreateProject Method
func CreateProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	projectValue := ctx.PostForm("value")
	result, _ := model.CreateProject(&model.Project{
		User:  username.(string),
		Value: projectValue,
	})
	if result {
		ctx.Status(200)
	} else {
		ctx.Status(400)
	}
}

// QueryProject Method
func QueryProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	result, err := model.QueryProject(username.(string))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}

// UpdateProject Method
func UpdateProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	flag, result := model.UpdateProject(&model.Project{
		ID:    ctx.PostForm("id"),
		User:  username.(string),
		Value: ctx.PostForm("value"),
	})
	if flag {
		ctx.Status(200)
	} else {
		ctx.String(400, result)
	}
}

// DeleteProject Method
// TODO: Check Project belong user
func DeleteProject(ctx *gin.Context) {
	// username, _ := ctx.Get("username")
	flag, result := model.DeleteProject(ctx.Query("id"))
	if flag {
		ctx.Status(200)
	} else {
		ctx.String(400, result)
	}
}
