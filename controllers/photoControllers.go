package controllers

import (
	"golang-mygram/database"
	"golang-mygram/helpers"
	"golang-mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetPhotos godoc
// @Summary Get all photos
// @Description Get all photos data
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /photos [get]
func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Photos := []models.Photo{}

	err := db.Find(&Photos).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully get all photos"
	response.Data = Photos
	c.JSON(http.StatusOK, response)
}

// GetPhoto godoc
// @Summary Get photo details for the given photo ID
// @Description Get details of a photo corresponding to the input ID
// @Tags photo
// @Accept json
// @Produce json
// @Param ID path int true "ID of the photo"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /photos/{ID} [get]
func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Photo := models.Photo{}
	photoID := c.Param("photoID")

	err := db.First(&Photo, "id = ?", photoID).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully get photo"
	response.Data = Photo
	c.JSON(http.StatusOK, response)
}

// CreatePhoto godoc
// @Summary Post a new photo
// @Description Post details of a new photo based on current user
// @Tags photo
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "create a photo"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	response := models.Response{}

	Photo := models.Photo{}

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint(userData["id"].(float64))

	err := db.Create(&Photo).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully created photo"
	response.Data = Photo
	c.JSON(http.StatusCreated, response)
}

// UpdatePhoto godoc
// @Summary Update photo for the given ID
// @Description Update details of a photo corresponding to the input ID
// @Tags photo
// @Accept json
// @Produce json
// @Param ID path int true "ID of the photo"
// @Param models.Photo body models.Photo true "update photo"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /photos/{ID} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Photo := models.Photo{}
	photoID := c.Param("photoID")

	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Photo).Where("id = ?", photoID).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully updated photo"
	c.JSON(http.StatusOK, response)
}

// DeletePhoto godoc
// @Summary Delete photo details for a given ID
// @Description Delete details of a photo corresponding to the input ID
// @Tags photo
// @Accept json
// @Produce json
// @Param ID path int true "ID of the photo"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /photos/{ID} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	response := models.Response{}

	Photo := models.Photo{}
	photoID := c.Param("photoID")

	err := db.Where("id = ?", photoID).Delete(&Photo).Error
	if err != nil {
		response.Error = "Bad Request"
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = "successfully deleted photo"
	c.JSON(http.StatusOK, response)
}
