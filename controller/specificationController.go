package controller

import (
	"API-REVIEWHP/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type specificationInput struct {
	Launched         string `json:"launched"`
	Body             string `json:"body"`
	OS               string `json:"os"`
	Storage          string `json:"storage"`
	TouchScreen      string `json:"touch_screen"`
	SizeScreen       string `json:"size_screen"`
	ResolutionScreen string `json:"resolution_screen"`
	Camera           string `json:"camera"`
	Video            string `json:"video"`
	RAM              string `json:"ram"`
	Chipset          string `json:"chipset"`
	CapacityBattery  string `json:"capacity_battery"`
}

// GetAllSpecification godoc
// @Summary Get all Specification.
// @Description Get a list of Specification.
// @Tags Specification
// @Produce json
// @Success 200 {object} []models.Specification
// @Router /specification [get]
func GetAllSpecification(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var specifications []models.Specification
	db.Find(&specifications)

	c.JSON(http.StatusOK, gin.H{"data": specifications})
}

// CreateSpecification godoc
// @Summary Create New Specification.
// @Description Creating a new Specification.
// @Tags Specification
// @Param Body body specificationInput true "the body to create a new Specification"
// @Produce json
// @Success 200 {object} models.Specification
// @Router /specification [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateSpecification(c *gin.Context) {
	// Validate input
	var input specificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Specification
	specification := models.Specification{
		Launched:         input.Launched,
		Body:             input.Body,
		OS:               input.OS,
		Storage:          input.Storage,
		TouchScreen:      input.TouchScreen,
		SizeScreen:       input.SizeScreen,
		ResolutionScreen: input.ResolutionScreen,
		Camera:           input.Camera,
		Video:            input.Video,
		RAM:              input.RAM,
		Chipset:          input.Chipset,
		CapacityBattery:  input.CapacityBattery,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&specification)

	c.JSON(http.StatusOK, gin.H{"data": specification})
}

// GetSpecificationById godoc
// @Summary Get Specification.
// @Description Get an Specification by id.
// @Tags Specification
// @Produce json
// @Param id path string true "Specification id"
// @Success 200 {object} models.Specification
// @Router /specification/{id} [get]
func GetSpecificationById(c *gin.Context) { // Get model if exist
	var specification models.Specification

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&specification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": specification})
}

// GetDevicesBySpecificationId godoc
// @Summary Get Devices.
// @Description Get all Devices by SpecificationId.
// @Tags Specification
// @Produce json
// @Param id path string true "Specification id"
// @Success 200 {object} []models.Devices
// @Router /specification/{id}/devices [get]
func GetDevicesBySpecificationId(c *gin.Context) { // Get model if exist
	var device []models.Devices

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("specification_id = ?", c.Param("id")).Find(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": device})
}

// UpdateSpecification godoc
// @Summary Update Specification.
// @Description Update Specification by id.
// @Tags Specification
// @Produce json
// @Param id path string true "Specification id"
// @Param Body body specificationInput true "the body to update age specification category"
// @Success 200 {object} models.Specification
// @Router /specification/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateSpecification(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var specification models.Specification
	if err := db.Where("id = ?", c.Param("id")).First(&specification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input specificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Specification
	updatedInput.Launched = input.Launched
	updatedInput.Body = input.Body
	updatedInput.OS = input.OS
	updatedInput.Storage = input.Storage
	updatedInput.TouchScreen = input.TouchScreen
	updatedInput.SizeScreen = input.SizeScreen
	updatedInput.ResolutionScreen = input.ResolutionScreen
	updatedInput.Camera = input.Camera
	updatedInput.Video = input.Video
	updatedInput.RAM = input.RAM
	updatedInput.Chipset = input.Chipset
	updatedInput.CapacityBattery = input.CapacityBattery
	updatedInput.UpdatedAt = time.Now()

	db.Model(&specification).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": specification})
}

// DeleteSpecification godoc
// @Summary Delete one Specification.
// @Description Delete a Specification by id.
// @Tags Specification
// @Produce json
// @Param id path string true "Specification id"
// @Success 200 {object} map[string]boolean
// @Router /specification/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteSpecification(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var specification models.Specification
	if err := db.Where("id = ?", c.Param("id")).First(&specification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&specification)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
