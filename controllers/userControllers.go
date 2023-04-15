package controllers

import (
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

type loginInput struct {
	Email    string `json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

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
// @Param Input body loginInput true "login a user"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	loginInput := loginInput{}

	if contentType == appJSON {
		c.ShouldBindJSON(&loginInput)
	} else {
		c.ShouldBind(&loginInput)
	}

	err := db.Debug().Where("email = ?", loginInput.Email).Take(&User).Error

	if err != nil {
		response.Error = "Unauthorized"
		response.Message = "invalid email / password"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(loginInput.Password))

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
