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
	classroomRepository := repositories.NewClassroom(db)
	taskRepository := repositories.NewTask(db)
	classrooomTaskRepository := repositories.NewClassroomTask(db)
	classroomSubjectRepository := repositories.NewClassroomSubject(db)

	//services
	jwtService := services.NewJWT(config.JWT)
	userService := services.NewUser(userRepository, jwtService)
	subjectService := services.NewSubject(subjectRepository)
	lectureService := services.NewLecture(lectureRepository)
	classroomService := services.NewClassroom(classroomRepository, userRepository)
	taskService := services.NewTask(taskRepository, userService)
	classroomTaskService := services.NewClassroomTask(classrooomTaskRepository, taskService, classroomService)
	classroomSubjectService := services.NewClassroomSubject(classroomSubjectRepository, classroomService, subjectService)

	app := fiber.New(fiber.Config{
		AppName: "Elearning nih bos",
	})

	//controllers
	controllers.NewUser(app, userService)
	controllers.NewSubject(app, subjectService, jwtService)
	controllers.NewLecture(app, lectureService, jwtService)
	controllers.NewClassroom(app, classroomService, jwtService)
	controllers.NewTask(app, taskService, jwtService)
	controllers.NewClassroomTaskController(app, classroomTaskService, jwtService)
	controllers.NewClassroomSubject(app, classroomSubjectService, jwtService)

	app.Listen(":" + config.Server.Port)
}
