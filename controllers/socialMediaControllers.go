package controllers

import (
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetSocialMedias godoc
// @Summary Get all social media
// @Description Get all social media data
// @Tags social media
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /socialmedia [get]
func GetSocialMedias(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	SocialMedias := []models.Socialmedia{}

	err := db.Find(&SocialMedias).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully get all social media"
	response.Data = SocialMedias
	c.JSON(http.StatusOK, response)
}

// GetSocialMedia godoc
// @Summary Get social media details for the given ID
// @Description Get details of a social media corresponding to the input ID
// @Tags social media
// @Accept json
// @Produce json
// @Param ID path int true "ID of the social media"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /socialmedia/{ID} [get]
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Socialmedia := models.Socialmedia{}
	socialMediaID := c.Param("socialMediaID")

	err := db.First(&Socialmedia, "id = ?", socialMediaID).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully get social media"
	response.Data = Socialmedia
	c.JSON(http.StatusOK, response)
}

// CreateSocialMedia godoc
// @Summary Post a new social media
// @Description Post details of a new social media based on current user
// @Tags social media
// @Accept json
// @Produce json
// @Param models.Socialmedia body models.Socialmedia true "create a social media"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /socialmedia [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	response := models.Response{}

	Socialmedia := models.Socialmedia{}

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = uint(userData["id"].(float64))

	err := db.Create(&Socialmedia).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully created social media"
	response.Data = Socialmedia
	c.JSON(http.StatusCreated, response)
}

// UpdateSocialMedia godoc
// @Summary Update social media for the given ID
// @Description Update details of a social media corresponding to the input ID
// @Tags social media
// @Accept json
// @Produce json
// @Param ID path int true "ID of the social media"
// @Param models.Socialmedia body models.Socialmedia true "update social media"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /socialmedia/{ID} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Socialmedia := models.Socialmedia{}
	socialMediaID := c.Param("socialMediaID")

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	err := db.Model(&Socialmedia).Where("id = ?", socialMediaID).Updates(models.Socialmedia{Name: Socialmedia.Name, SocialMediaURL: Socialmedia.SocialMediaURL}).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully updated social media"
	c.JSON(http.StatusOK, response)
}

// DeleteSocialMedia godoc
// @Summary Delete social media details for a given ID
// @Description Delete details of a social media corresponding to the input ID
// @Tags social media
// @Accept json
// @Produce json
// @Param ID path int true "ID of the social media"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /socialmedia/{ID} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Socialmedia := models.Socialmedia{}
	socialMediaID := c.Param("socialMediaID")

	err := db.Where("id = ?", socialMediaID).Delete(&Socialmedia).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully deleted social media"
	c.JSON(http.StatusOK, response)
}
