package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type ClassroomTaskController struct {
	classroomTaskService services.ClassroomTaskService
	jwtService           services.JWTService
}

func NewClassroomTaskController(app *fiber.App, classroomTaskService services.ClassroomTaskService, jwtService services.JWTService) {
	classroomTask := ClassroomTaskController{
		classroomTaskService: classroomTaskService,
		jwtService:           jwtService,
	}

	api := app.Group("/api/v1/")
	api.Post("/classroom-tasks", middlewares.CheckAuth(jwtService), classroomTask.AsignTask)
	api.Get("/classroom-tasks/:id", classroomTask.GetSingle)
	api.Get("/classroom-tasks/classroom/:classroomID", classroomTask.GetByClassroomID)
}

func (c *ClassroomTaskController) AsignTask(ctx *fiber.Ctx) error {
	var asignTaskDto dto.AsignTaskDTO

	if err := ctx.BodyParser(&asignTaskDto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	errors := helpers.ValidateRequest(asignTaskDto)
	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Bad request", errors),
		)
	}

	userId, ok := ctx.Locals("userId").(int)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError("User ID not found or invalid type"),
		)
	}

	asignTaskDto.TeacherID = uint(userId)
	res, err := c.classroomTaskService.AsignTask(ctx.Context(), asignTaskDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(res),
	)

}

func (c *ClassroomTaskController) GetSingle(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.classroomTaskService.FindByID(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *ClassroomTaskController) GetByClassroomID(ctx *fiber.Ctx) error {
	classroomID, err := strconv.Atoi(ctx.Params("classroomID"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.classroomTaskService.FindByClassroomID(ctx.Context(), classroomID)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}
