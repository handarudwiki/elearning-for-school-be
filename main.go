package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/config"
	"github.com/handarudwiki/controllers"
	dbconnection "github.com/handarudwiki/database/migrations/db_connection"
	"github.com/handarudwiki/repositories"
	"github.com/handarudwiki/services"
)

func main() {
	config := config.Get()
	db, err := dbconnection.GetDatabaseConnection(config)

	if err != nil {
		panic(err)
	}

	//repositories
	userRepository := repositories.NewUser(db)
	subjectRepository := repositories.NewSubject(db)
	lectureRepository := repositories.NewLecture(db)

	//services
	jwtService := services.NewJWT(config.JWT)
	userService := services.NewUser(userRepository, jwtService)
	subjectService := services.NewSubject(subjectRepository)
	lectureService := services.NewLecture(lectureRepository)

	app := fiber.New(fiber.Config{
		AppName: "Elearning nih bos",
	})

	controllers.NewUser(app, userService)
	controllers.NewSubject(app, subjectService, jwtService)
	controllers.NewLecture(app, lectureService, jwtService)

	app.Listen(":" + config.Server.Port)
}
