package controller

import (
	"fmt"
	"github.com/Penetration-Platform-Go/Project-Service/model"
	"github.com/gin-gonic/gin"
)

// CreateProject Method
func CreateProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	var project model.Project
	project.User = username.(string)
	err := ctx.BindJSON(&project)
	if err != nil {
		fmt.Println(err)
		ctx.Status(400)
		return
	}
	result, _ := model.CreateProject(&project)
	if result {
		ctx.Status(200)
	} else {
		ctx.Status(400)
	}
}

// QueryProject Method
func QueryProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	result, err := model.QueryProjectsByUsername(username.(string))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}

// UpdateProject Method
// TODO: Check Project belong user
func UpdateProject(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	var project model.Project
	project.User = username.(string)
	err := ctx.BindJSON(&project)
	if err != nil {
		fmt.Println(err)
		ctx.Status(400)
		return
	}
	flag, result := model.UpdateProject(&project)
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
