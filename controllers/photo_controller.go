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

// GetPhotos godoc
// @Summary Get detail
// @Description Get detail all of photos
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetPhotos(ctx *gin.Context) {
	db := database.GetDB()
	photos := []models.Photo{}
	db.Find(&photos)

	for i := 0; i < len(photos); i++ {
		if err := db.Debug().Where("id = ?", photos[i].UserID).First(&photos[i].User).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"data": photos})
}

// GetPhoto godoc
// @Summary Get detail
// @Description Get detail of photo find by id
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [get]
func GetPhoto(ctx *gin.Context) {
	photo := models.Photo{}

	db := database.GetDB()
	if err := db.Where("id = ?", ctx.Param("photoId")).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Debug().Where("id = ?", photo.UserID).First(&photo.User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": photo})
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete of photo find by id
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [delete]
// @Security ApiKeyAuth
func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	if err := db.Where("id = ?", ctx.Param("photoId")).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!!"})
		return
	}

	db.Delete(&photo)

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Photo Success"})
}

// CreatePhoto godoc
// @Summary Create the photo
// @Description Create the photo
// @Tags photos
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "create photo"
// @Success 200 {object} models.Photo
// @Router /photos [post]
// @Security ApiKeyAuth
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)

	photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	photo.UserID = userID

	if err := db.Debug().Where("id = ?", userID).First(&photo.User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	err := db.Debug().Create(&photo).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created Photo Success", "data": photo})
}

// UpdatePhoto godoc
// @Summary Update the photo
// @Description Update the photo
// @Tags photos
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "update photo"
// @Success 200 {object} models.Photo
// @Router /photos/{photoId} [put]
// @Security ApiKeyAuth
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContenType(ctx)
	photo := models.Photo{}

	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	photo.UserID = userID
	photo.ID = uint(photoId)

	err := db.Model(&photo).Where("id = ?", photoId).Updates(models.Photo{Title: photo.Title, PhotoUrl: photo.PhotoUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": photo})
}
