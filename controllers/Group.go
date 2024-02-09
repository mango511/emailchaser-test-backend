package controller

import (
	"main/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Group struct {
	name string `form:"group`
}

func CreateGroupHandler(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var group Group

		if err := ctx.ShouldBindJSON(&group); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := client.Group.
			Create().
			SetName(group.name).
			Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, u)

	}
}
