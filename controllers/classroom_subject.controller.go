package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type classroomSubjectController struct {
	classroomSubjectService services.ClassroomSubjectService
	jwtService              services.JWTService
}

func NewClassroomSubject(app *fiber.App, classroomSubjectService services.ClassroomSubjectService, jwtService services.JWTService) {
	classroomSubjectController := classroomSubjectController{
		classroomSubjectService: classroomSubjectService,
		jwtService:              jwtService,
	}

	classroomSubjects := app.Group("/api/v1/classroom-subjects")

	classroomSubjects.Get("/", middlewares.CheckAuth(jwtService), classroomSubjectController.FindByTeacherID)
	classroomSubjects.Post("/", middlewares.CheckAuth(jwtService), classroomSubjectController.Create)
}

func (c *classroomSubjectController) FindByTeacherID(ctx *fiber.Ctx) error {
	teacherId := ctx.Locals("userId").(int)

	res, err := c.classroomSubjectService.FindByTeacherID(ctx.Context(), teacherId)

	if err != nil {
		statusCode := helpers.GetHttpStatusCode(err)

		return ctx.Status(statusCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *classroomSubjectController) Create(ctx *fiber.Ctx) error {
	dto := new(dto.CreateClassrooomSubject)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := helpers.ValidateRequest(dto)

	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Invalid input", errors),
		)
	}

	userId := ctx.Locals("userId").(int)

	dto.TeacherID = uint(userId)

	res, err := c.classroomSubjectService.Create(ctx.Context(), *dto)

	if err != nil {
		statusCode := helpers.GetHttpStatusCode(err)

		return ctx.Status(statusCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(res),
	)
}
