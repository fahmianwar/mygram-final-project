package controllers

import (
	"net/http"
	"strconv"

	"mygram-final-project/database"
	"mygram-final-project/helpers"
	"mygram-final-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetComments godoc
// @Summary Get details
// @Description Get details of all comment
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments [get]
func GetComments(ctx *gin.Context) {
	db := database.GetDB()
	comments := []models.Comment{}
	db.Find(&comments)

	for i := 0; i < len(comments); i++ {
		if err := db.Debug().Where("id = ?", comments[i].UserID).First(&comments[i].User).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comments})
}

// GetComment godoc
// @Summary Get detail
// @Description Get detail of comment find by id
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [get]
func GetComment(ctx *gin.Context) {
	comment := models.Comment{}

	db := database.GetDB()
	if err := db.Where("id = ?", ctx.Param("commentId")).First(&comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Debug().Where("id = ?", comment.UserID).First(&comment.User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteComment godoc
// @Summary Delete detail
// @Description Delete of comment find by id
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Security ApiKeyAuth
// @Router /comments/{commentId} [delete]
func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	if err := db.Where("id = ?", ctx.Param("commentId")).First(&comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!!"})
		return
	}

	db.Delete(&comment)

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Comment Success"})
}

// CreateComment godoc
// @Summary Create the comment
// @Description Create the comment
// @Tags comments
// @Accept json
// @Produce json
// @Param models.CommentRequest body models.CommentRequest true "create comment"
// @Success 200 {object} models.Comment
// @Security ApiKeyAuth
// @Router /comments [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)

	comment := models.Comment{}
	commentRequest := models.CommentRequest{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&commentRequest)
	} else {
		ctx.ShouldBind(&commentRequest)
	}

	comment.UserID = userID
	comment.Message = commentRequest.Message
	comment.PhotoID = commentRequest.PhotoID

	if err := db.Debug().Where("id = ?", userID).First(&comment.User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	err := db.Debug().Create(&comment).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created Comment Success", "data": comment})
}

// UpdateComment godoc
// @Summary Update the comment
// @Description Update the comment
// @Tags comments
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "update comment"
// @Success 200 {object} models.Comment
// @Security ApiKeyAuth
// @Router /comments/{commentId} [put]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)
	comment := models.Comment{}

	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	comment.UserID = userID
	comment.ID = uint(commentId)

	err := db.Model(&comment).Where("id = ?", commentId).Updates(models.Comment{Message: comment.Message}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": comment})
}
