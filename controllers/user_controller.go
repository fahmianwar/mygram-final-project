package controllers

import (
	"net/http"

	"mygram-final-project/database"
	"mygram-final-project/helpers"
	"mygram-final-project/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// UserRegister godoc
// @Summary Create the user
// @Description Create the user
// @Tags users
// @Accept json
// @Produce json
// @Param models.UserRegister body models.UserRegister true "create user"
// @Success 200 {object} models.User
// @Router /users/register [post]
func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContenType(ctx)
	userRegister := models.UserRegister{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&userRegister)
	} else {
		ctx.ShouldBind(&userRegister)
	}

	user := models.User{}
	user.Username = userRegister.Username
	user.Email = userRegister.Email
	user.Age = userRegister.Age

	err := db.Debug().Create(&user).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"email":    user.Email,
		"username": user.Username,
	})
}

// UserLogin godoc
// @Summary Create the user
// @Description Create the user
// @Tags users
// @Accept json
// @Produce json
// @Param models.UserLogin body models.UserLogin true "login user"
// @Success 200 {object} models.User
// @Router /users/login [post]
func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContenType(ctx)
	userLogin := models.UserLogin{}
	user := models.User{}

	var password string

	if contentType == appJSON {
		ctx.ShouldBindJSON(&userLogin)
	} else {
		ctx.ShouldBind(&userLogin)
	}

	password = userLogin.Password

	err := db.Debug().Where("username = ?", userLogin.Username).Take(&user).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
			"data":    user,
			"pasd":    password,
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
