package main

import (
	"database/sql"
	albums "example/web-service-gin/v1/Albums"
	employees "example/web-service-gin/v1/Employees"
	docs "example/web-service-gin/v1/docs"
	"log"
	"net/http"

	//albums "/Albums/albums"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {

	// Define the SQL Server connection string
	connectionString := "server=localhost;user id=arumteguh;password=123456;port=1434;database=Dev1"

	// Open a connection to the SQL Server database
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	albumRepository := albums.NewDefaultAlbumRepository("database")
	albumService := albums.NewDefaultAlbumService(albumRepository)
	albumController := albums.NewAlbumController(albumService)

	employeeRepository := employees.NewDefaultEmployeeRepository(db)
	employeeService := employees.NewDefaultEmployeeService(employeeRepository)
	employeeController := employees.NewEmployeeController(employeeService)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	v1 := router.Group("/")
	{
		eg := v1.Group("/employee")
		{
			eg.GET("/employee", employeeController.GetEmployees)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/album", albumController.GetAlbums)

	router.GET("/employee", employeeController.GetEmployees)
	router.GET("/employee/:id", employeeController.GetEmployee)
	router.POST("/employee", employeeController.CreateEmployee)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
