package controllers

import (
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetComments godoc
// @Summary Get all comments
// @Description Get all comments data
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /comments [get]
func GetComments(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Comments := []models.Comment{}

	err := db.Find(&Comments).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully get all comments"
	response.Data = Comments
	c.JSON(http.StatusOK, response)
}

// GetComment godoc
// @Summary Get comment details for the given ID
// @Description Get details of a comment corresponding to the input ID
// @Tags comment
// @Accept json
// @Produce json
// @Param ID path int true "ID of the comment"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /comments/{ID} [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Comment := models.Comment{}
	commentID := c.Param("commentID")

	err := db.First(&Comment, "id = ?", commentID).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully get comment"
	response.Data = Comment
	c.JSON(http.StatusOK, response)
}

// CreateComment godoc
// @Summary Post a new comment
// @Description Post details of a new comment based on current user and certain photo
// @Tags comment
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "create a comment"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	response := models.Response{}

	Comment := models.Comment{}

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Photo := models.Photo{}
	errPhoto := db.First(&Photo, "id = ?", Comment.PhotoID).Error
	if errPhoto != nil {
		response.Error = "Bad Request"
		response.Message = errPhoto.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	Comment.UserID = uint(userData["id"].(float64))

	errComment := db.Create(&Comment).Error
	if errComment != nil {
		response.Error = "Bad Request"
		response.Message = errComment.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully created comment"
	response.Data = Comment
	c.JSON(http.StatusCreated, response)
}

// UpdateComment godoc
// @Summary Update comment for the given ID
// @Description Update details of a comment corresponding to the input ID
// @Tags comment
// @Accept json
// @Produce json
// @Param ID path int true "ID of the comment"
// @Param models.Comment body models.Comment true "update comment"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /comments/{ID} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Comment := models.Comment{}
	commentID := c.Param("commentID")

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Photo := models.Photo{}
	errPhoto := db.First(&Photo, "id = ?", Comment.PhotoID).Error
	if errPhoto != nil {
		response.Error = "Bad Request"
		response.Message = errPhoto.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errComment := db.Model(&Comment).Where("id = ?", commentID).Updates(models.Comment{Message: Comment.Message}).Error
	if errComment != nil {
		response.Error = "Bad Request"
		response.Message = errComment.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully updated comment"
	c.JSON(http.StatusOK, response)
}

// DeleteComment godoc
// @Summary Delete comment details for a given ID
// @Description Delete details of a comment corresponding to the input ID
// @Tags comment
// @Accept json
// @Produce json
// @Param ID path int true "ID of the comment"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /comments/{ID} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Comment := models.Comment{}
	commentID := c.Param("commentID")

	err := db.Where("id = ?", commentID).Delete(&Comment).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully deleted comment"
	c.JSON(http.StatusOK, response)
}
