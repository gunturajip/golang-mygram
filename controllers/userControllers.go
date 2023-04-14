package controllers

import (
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user using email, username, and password
// @Tags user
// @Accept json
// @Produce json
// @Param models.User body models.User true "register a user"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Data = User
	c.JSON(http.StatusCreated, response)
}

// UserLogin godoc
// @Summary Login an existing user
// @Description Register an existing user using email, and password
// @Tags user
// @Accept json
// @Produce json
// @Param models.User body models.User true "login a user"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		response.Error = "Unauthorized"
		response.Message = "invalid email / password"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		response.Error = "Unauthorized"
		response.Message = "invalid email / password"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	response.Data = token
	c.JSON(http.StatusOK, response)
}
