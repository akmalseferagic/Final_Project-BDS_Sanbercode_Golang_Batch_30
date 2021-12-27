package controller

import (
	"API-REVIEWHP/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentsInput struct {
	Name            string `json:"your_name"`
	CommentsSection string `json:"comments_section"`
	DevicesID       uint   `json:"devices_id"`
	ReviewsID       uint   `json:"reviews_id"`
}

// GetAllComments godoc
// @Summary Get all Comments.
// @Description Get a list of Comments.
// @Tags Comments
// @Produce json
// @Success 200 {object} []models.Comments
// @Router /comments [get]
func GetAllComments(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var commentss []models.Comments
	db.Find(&commentss)

	c.JSON(http.StatusOK, gin.H{"data": commentss})
}

// CreateComments godoc
// @Summary Create New Comments.
// @Description Creating a new Comments.
// @Tags Comments
// @Param Body body commentsInput true "the body to create a new Comments"
// @Produce json
// @Success 200 {object} models.Comments
// @Router /comments [post]
func CreateComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input commentsInput
	var reviews models.Reviews
	var devices models.Devices
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewsID).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReviewsID not found!"})
		return
	}

	if err := db.Where("id = ?", input.DevicesID).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	// Create Comments
	comments := models.Comments{ReviewsID: input.ReviewsID, DevicesID: input.DevicesID, Name: input.Name, CommentsSection: input.CommentsSection}
	db.Create(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// GetCommentsById godoc
// @Summary Get Comments.
// @Description Get an Comments by id.
// @Tags Comments
// @Produce json
// @Param id path string true "Comments id"
// @Success 200 {object} models.Comments
// @Router /comments/{id} [get]
func GetCommentsById(c *gin.Context) { // Get model if exist
	var comments models.Comments

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// UpdateComments godoc
// @Summary Update Comments.
// @Description Update Comments by id.
// @Tags Comments
// @Produce json
// @Param id path string true "Comments id"
// @Param Body body commentsInput true "the body to update comments"
// @Success 200 {object} models.Comments
// @Router /comments/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateComments(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var comments models.Comments
	var reviews models.Reviews
	var devices models.Devices
	if err := db.Where("id = ?", c.Param("id")).First(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input commentsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewsID).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	if err := db.Where("id = ?", input.DevicesID).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	var updatedInput models.Comments
	updatedInput.Name = input.Name
	updatedInput.CommentsSection = input.CommentsSection
	updatedInput.ReviewsID = input.ReviewsID
	updatedInput.DevicesID = input.DevicesID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&comments).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// DeleteComments godoc
// @Summary Delete one Comments.
// @Description Delete a Comments by id.
// @Tags Comments
// @Produce json
// @Param id path string true "Comments id"
// @Success 200 {object} map[string]boolean
// @Router /comments/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteComments(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var comments models.Comments
	if err := db.Where("id = ?", c.Param("id")).First(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comments)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
