package main

import (
	"gorm.io/driver/postgres"
	"log"
	"student/dbaccess"
	"student/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	
	err = db.AutoMigrate(&dbaccess.Student{}, &dbaccess.Subject{})
	if err != nil {
		log.Fatal("Failed to auto-migrate the database:", err)
	}

	r := gin.Default()

	service.Setup(db)
	
	r.GET("/students", service.GetAllStudentsHandler)
	r.GET("/students/:id", service.GetStudentByIDHandler)
	r.POST("/students", service.CreateStudentHandler)
	r.PUT("/students/:id", service.UpdateStudentHandler)
	r.DELETE("/students/:id", service.DeleteStudentHandler)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}

