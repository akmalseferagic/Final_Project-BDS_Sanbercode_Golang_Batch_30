package controller

import (
	"API-REVIEWHP/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type opinionsInput struct {
	Name           string `json:"your_name"`
	OpinionSection string `json:"opinion_section"`
	ReviewsID      uint   `json:"reviews_id"`
}

// GetAllOpinions godoc
// @Summary Get all Opinions.
// @Description Get a list of Opinions.
// @Tags Opinions
// @Produce json
// @Success 200 {object} []models.Opinions
// @Router /opinions [get]
func GetAllOpinions(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var opinionss []models.Opinions
	db.Find(&opinionss)

	c.JSON(http.StatusOK, gin.H{"data": opinionss})
}

// CreateOpinions godoc
// @Summary Create New Opinions.
// @Description Creating a new Opinions.
// @Tags Opinions
// @Param Body body opinionsInput true "the body to create a new Opinions"
// @Produce json
// @Success 200 {object} models.Opinions
// @Router /opinions [post]
func CreateOpinions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input opinionsInput
	var reviews models.Reviews
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewsID).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	// Create Opinions
	opinions := models.Opinions{ReviewsID: input.ReviewsID, Name: input.Name, OpinionSection: input.OpinionSection}
	db.Create(&opinions)

	c.JSON(http.StatusOK, gin.H{"data": opinions})
}

// GetOpinionsById godoc
// @Summary Get Opinions.
// @Description Get an Opinions by id.
// @Tags Opinions
// @Produce json
// @Param id path string true "Opinions id"
// @Success 200 {object} models.Opinions
// @Router /opinions/{id} [get]
func GetOpinionsById(c *gin.Context) { // Get model if exist
	var opinions models.Opinions

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&opinions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": opinions})
}

// UpdateOpinions godoc
// @Summary Update Opinions.
// @Description Update Opinions by id.
// @Tags Opinions
// @Produce json
// @Param id path string true "Opinions id"
// @Param Body body opinionsInput true "the body to update opinions"
// @Success 200 {object} models.Opinions
// @Router /opinions/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateOpinions(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var opinions models.Opinions
	var reviews models.Reviews
	if err := db.Where("id = ?", c.Param("id")).First(&opinions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input opinionsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewsID).First(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DeviceID not found!"})
		return
	}

	var updatedInput models.Opinions
	updatedInput.Name = input.Name
	updatedInput.OpinionSection = input.OpinionSection
	updatedInput.ReviewsID = input.ReviewsID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&opinions).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": opinions})
}

// DeleteOpinions godoc
// @Summary Delete one Opinions.
// @Description Delete a Opinions by id.
// @Tags Opinions
// @Produce json
// @Param id path string true "Opinions id"
// @Success 200 {object} map[string]boolean
// @Router /opinions/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteOpinions(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var opinions models.Opinions
	if err := db.Where("id = ?", c.Param("id")).First(&opinions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&opinions)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
