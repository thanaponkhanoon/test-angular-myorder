package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanaponkhanoon/test-angular-myorder/controller"
	"github.com/thanaponkhanoon/test-angular-myorder/entity"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	router := r.Group("/")
	{
		// User Routes
		router.GET("/employees", controller.ListEmployees)
		router.GET("/employee/:id", controller.GetEmployee)
		router.PATCH("/employees", controller.UpdateEmployee)
		router.DELETE("/employees/:id", controller.DeleteEmployee)
		router.PATCH("/employee", controller.CreateEmployee)
		// Category Routes
		router.GET("/catagories", controller.ListCategory)
		router.GET("/catagory/:id", controller.GetCategory)
		router.POST("/catagories", controller.CreateCategory)
		router.PATCH("/catagories", controller.UpdateCategory)
		router.DELETE("/catagories/:id", controller.DeleteCategory)
		// Unit Routes
		router.GET("/units", controller.ListUnits)
		router.GET("/unit/:id", controller.GetUnit)
		router.POST("/units", controller.CreateUnit)
		router.PATCH("/units", controller.UpdateUnit)
		router.DELETE("units/:id", controller.DeleteUnit)
		// Equipment Routes
		router.GET("/equipments", controller.ListEquipments)
		router.GET("/equipment/:id", controller.GetEquipment)
		router.POST("/equipment", controller.CreateEquipment)
		router.PATCH("/equipments", controller.UpdateEquipment)
		router.DELETE("/equipments/:id", controller.DeleteEquipment)
	}
	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
