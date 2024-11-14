package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type taskController struct {
	taskService services.TaskService
	jwtService  services.JWTService
}

func NewTask(app *fiber.App, taskService services.TaskService, jwtService services.JWTService) {
	controller := taskController{
		taskService: taskService,
		jwtService:  jwtService,
	}

	tasks := app.Group("/api/v1/tasks")
	tasks.Post("/", middlewares.CheckAuth(jwtService), controller.CreateTask)
	tasks.Get("/:id", controller.GetSingle)
}
func (c *taskController) CreateTask(ctx *fiber.Ctx) error {
	var createTaskDto dto.CreateTaskDTO

	// Parse body request ke struct DTO
	if err := ctx.BodyParser(&createTaskDto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	// Validasi request
	errors := helpers.ValidateRequest(createTaskDto)
	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Bad request", errors),
		)
	}

	userId, ok := ctx.Locals("userId").(int)
	fmt.Println(userId)

	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError("User ID not found or invalid type"),
		)
	}
	createTaskDto.UserID = uint(userId)
	res, err := c.taskService.Create(ctx.Context(), createTaskDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}
	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *taskController) GetSingle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}
	res, err := c.taskService.FindByID(ctx.Context(), taskId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}
	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}
