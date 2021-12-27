package controller

import (
	"API-REVIEWHP/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type devicesInput struct {
	BrandID         uint   `json:"brand_id"`
	SpecificationID uint   `json:"specification_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           string `json:"price"`
}

// GetAllDevices godoc
// @Summary Get all Devices.
// @Description Get a list of Devices.
// @Tags Devices
// @Produce json
// @Success 200 {object} []models.Devices
// @Router /devices [get]
func GetAllDevices(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var devicess []models.Devices
	db.Find(&devicess)

	c.JSON(http.StatusOK, gin.H{"data": devicess})
}

// GetCommentsByDevicesId godoc
// @Summary Get Comments By id Devices.
// @Description Get all Opinion by DevicesId.
// @Tags Devices
// @Produce json
// @Param id path string true "Devices id"
// @Success 200 {object} []models.Comments
// @Router /devices/{id}/comments [get]
func GetCommentsByDevicesId(c *gin.Context) { // Get model if exist
	var comments []models.Comments

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("devices_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateDevices godoc
// @Summary Create New Devices.
// @Description Creating a new Devices.
// @Tags Devices
// @Param Body body devicesInput true "the body to create a new Devices"
// @Produce json
// @Success 200 {object} models.Devices
// @Router /devices [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateDevices(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input devicesInput
	var brand models.Brand
	var specification models.Specification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.BrandID).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brandID not found!"})
		return
	}

	if err := db.Where("id = ?", input.SpecificationID).First(&specification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "specificationID not found!"})
		return
	}

	// Create Devices
	devices := models.Devices{BrandID: input.BrandID, SpecificationID: input.SpecificationID, Description: input.Description, Name: input.Name, Price: input.Price}
	db.Create(&devices)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// GetDevicesById godoc
// @Summary Get Devices.
// @Description Get an Devices by id.
// @Tags Devices
// @Produce json
// @Param id path string true "Devices id"
// @Success 200 {object} models.Devices
// @Router /devices/{id} [get]
func GetDevicesById(c *gin.Context) { // Get model if exist
	var devices models.Devices

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// UpdateDevices godoc
// @Summary Update Devices.
// @Description Update Devices by id.
// @Tags Devices
// @Produce json
// @Param id path string true "Devices id"
// @Param Body body devicesInput true "the body to update devices"
// @Success 200 {object} models.Devices
// @Router /devices/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateDevices(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var devices models.Devices
	var brand models.Brand
	var specification models.Specification
	if err := db.Where("id = ?", c.Param("id")).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input devicesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.BrandID).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brandID not found!"})
		return
	}

	if err := db.Where("id = ?", input.SpecificationID).First(&specification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "specificationID not found!"})
		return
	}

	var updatedInput models.Devices
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.BrandID = input.BrandID
	updatedInput.SpecificationID = input.SpecificationID
	updatedInput.Price = input.Price
	updatedInput.UpdatedAt = time.Now()

	db.Model(&devices).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// DeleteDevices godoc
// @Summary Delete one Devices.
// @Description Delete a Devices by id.
// @Tags Devices
// @Produce json
// @Param id path string true "Devices id"
// @Success 200 {object} map[string]boolean
// @Router /devices/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteDevices(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var devices models.Devices
	if err := db.Where("id = ?", c.Param("id")).First(&devices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&devices)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
