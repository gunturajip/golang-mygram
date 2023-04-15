package controllers

import (
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

type registerInput struct {
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required, numeric~Age must be numeric"`
}

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
// @Param Input body registerInput true "register a user"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users/register [post]
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	registerInput := registerInput{}

	if contentType == appJSON {
		c.ShouldBindJSON(&registerInput)
	} else {
		c.ShouldBind(&registerInput)
	}

	User.Email = registerInput.Email
	User.Password = registerInput.Password
	User.Username = registerInput.Username
	User.Age = registerInput.Age

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
