package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"API-REVIEWHP/controller"
	"API-REVIEWHP/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	//DEVICE
	devicesMiddlewareRoute := r.Group("/devices")
	devicesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	devicesMiddlewareRoute.POST("/", controller.CreateDevices)
	devicesMiddlewareRoute.PATCH("/:id", controller.UpdateDevices)
	devicesMiddlewareRoute.DELETE("/:id", controller.DeleteDevices)
	//hanya user biasa
	r.GET("/devices", controller.GetAllDevices)
	r.GET("/devices/:id", controller.GetDevicesById)
	r.GET("/devices/:id/comments", controller.GetCommentsByDevicesId)

	//BRAND
	brandMiddlewareRoute := r.Group("/brand")
	brandMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	brandMiddlewareRoute.POST("/", controller.CreateBrand)
	brandMiddlewareRoute.PATCH("/:id", controller.UpdateBrand)
	brandMiddlewareRoute.DELETE("/:id", controller.DeleteBrand)
	//hanya user biasa
	r.GET("/brand", controller.GetAllBrand)
	r.GET("/brand/:id", controller.GetBrandById)
	r.GET("/brand/:id/devices", controller.GetDevicesByBrandId)

	//SPEK
	specificationMiddlewareRoute := r.Group("/specification")
	specificationMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	specificationMiddlewareRoute.POST("/", controller.CreateSpecification)
	specificationMiddlewareRoute.PATCH("/:id", controller.UpdateSpecification)
	specificationMiddlewareRoute.DELETE("/:id", controller.DeleteSpecification)
	//hanya user biasa
	r.GET("/specification", controller.GetAllSpecification)
	r.GET("/specification/:id", controller.GetSpecificationById)
	r.GET("/specification/:id/devices", controller.GetDevicesBySpecificationId)

	//REVIEW
	reviewsMiddlewareRoute := r.Group("/reviews")
	reviewsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewsMiddlewareRoute.POST("/", controller.CreateReviews)
	reviewsMiddlewareRoute.PATCH("/:id", controller.UpdateReviews)
	reviewsMiddlewareRoute.DELETE("/:id", controller.DeleteReviews)
	//hanya user biasa
	r.GET("/reviews", controller.GetAllReviews)
	r.GET("/reviews/:id", controller.GetReviewsById)
	r.GET("/reviews/:id/opinions", controller.GetOpinionsByReviewsId)
	r.GET("/reviews/:id/comments", controller.GetCommentsByReviewsId)

	//REVIEW
	opinionsMiddlewareRoute := r.Group("/opinions")
	opinionsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	opinionsMiddlewareRoute.PATCH("/:id", controller.UpdateOpinions)
	opinionsMiddlewareRoute.DELETE("/:id", controller.DeleteOpinions)
	//hanya user biasa
	r.POST("/opinions", controller.CreateOpinions)
	r.GET("/opinions", controller.GetAllOpinions)
	r.GET("/opinions/:id", controller.GetOpinionsById)

	//REVIEW
	commentsMiddlewareRoute := r.Group("/comments")
	commentsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	commentsMiddlewareRoute.PATCH("/:id", controller.UpdateComments)
	commentsMiddlewareRoute.DELETE("/:id", controller.DeleteComments)
	//hanya user biasa
	r.POST("/comments", controller.CreateComments)
	r.GET("/comments", controller.GetAllComments)
	r.GET("/comments/:id", controller.GetCommentsById)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
