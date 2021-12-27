package controller

import (
	"API-REVIEWHP/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewsInput struct {
	Name      string `json:"name_article"`
	Article   string `json:"article"`
	DevicesID uint   `json:"devices_id"`
}

// GetAllReviews godoc
// @Summary Get all Reviews.
// @Description Get a list of Reviews.
// @Tags Reviews
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /reviews [get]
func GetAllReviews(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var reviewss []models.Reviews
	db.Find(&reviewss)

	c.JSON(http.StatusOK, gin.H{"data": reviewss})
}

// GetOpinionsByReviewsId godoc
// @Summary Get Opinions By id Reviews.
// @Description Get all Opinion by ReviewsId.
// @Tags Reviews
// @Produce json
// @Param id path string true "Reviews id"
// @Success 200 {object} []models.Opinions
// @Router /reviews/{id}/opinions [get]
func GetOpinionsByReviewsId(c *gin.Context) { // Get model if exist
	var opinions []models.Opinions

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("reviews_id = ?", c.Param("id")).Find(&opinions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": opinions})
}

// GetCommentsByReviewsId godoc
// @Summary Get Comments By id Reviews.
// @Description Get all Opinion by ReviewsId.
// @Tags Reviews
// @Produce json
// @Param id path string true "Reviews id"
// @Success 200 {object} []models.Comments
// @Router /reviews/{id}/comments [get]
func GetCommentsByReviewsId(c *gin.Context) { // Get model if exist
	var comments []models.Comments

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("reviews_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateReviews godoc
// @Summary Create New Reviews.
// @Description Creating a new Reviews.
// @Tags Reviews
// @Param Body body reviewsInput true "the body to create a new Reviews"
// @Produce json
// @Success 200 {object} models.Reviews
// @Router /reviews [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateReviews(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input reviewsInput
	var devices models.Devices
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.DevicesID).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	// Create Reviews
	reviews := models.Reviews{DevicesID: input.DevicesID, Article: input.Article, Name: input.Name}
	db.Create(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// GetReviewsById godoc
// @Summary Get Reviews.
// @Description Get an Reviews by id.
// @Tags Reviews
// @Produce json
// @Param id path string true "Reviews id"
// @Success 200 {object} models.Reviews
// @Router /reviews/{id} [get]
func GetReviewsById(c *gin.Context) { // Get model if exist
	var reviews models.Reviews

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// UpdateReviews godoc
// @Summary Update Reviews.
// @Description Update Reviews by id.
// @Tags Reviews
// @Produce json
// @Param id path string true "Reviews id"
// @Param Body body reviewsInput true "the body to update reviews"
// @Success 200 {object} models.Reviews
// @Router /reviews/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateReviews(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var reviews models.Reviews
	var devices models.Devices
	if err := db.Where("id = ?", c.Param("id")).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input reviewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.DevicesID).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	var updatedInput models.Reviews
	updatedInput.Name = input.Name
	updatedInput.Article = input.Article
	updatedInput.DevicesID = input.DevicesID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&reviews).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// DeleteReviews godoc
// @Summary Delete one Reviews.
// @Description Delete a Reviews by id.
// @Tags Reviews
// @Produce json
// @Param id path string true "Reviews id"
// @Success 200 {object} map[string]boolean
// @Router /reviews/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteReviews(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var reviews models.Reviews
	if err := db.Where("id = ?", c.Param("id")).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
