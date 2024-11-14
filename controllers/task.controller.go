package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/commons"
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
	tasks.Get("/", controller.GetTasks)
	tasks.Put("/:id", middlewares.CheckAuth(jwtService), controller.UpdateTask)
	tasks.Delete("/:id", middlewares.CheckAuth(jwtService), controller.DeleteTask)
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

func (c *taskController) UpdateTask(ctx *fiber.Ctx) error {
	var updateTaskDto dto.UpdateTaskDTO
	if err := ctx.BodyParser(&updateTaskDto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	errors := helpers.ValidateRequest(updateTaskDto)
	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Bad request", errors),
		)
	}

	id := ctx.Params("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.taskService.Update(ctx.Context(), taskId, updateTaskDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *taskController) DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.taskService.Delete(ctx.Context(), taskId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("Task deleted successfully"),
	)
}

func (c *taskController) GetTasks(ctx *fiber.Ctx) error {

	var dto dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	search := ctx.Query("search")

	dto.Page = page
	dto.Size = size
	dto.Search = &search

	isActive := ctx.Query("is_active")

	if isActive != "" {
		isActiveBool, err := strconv.ParseBool(isActive)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				helpers.ResponseError(err.Error()),
			)
		}
		dto.Is_active = &isActiveBool
	}

	tasks, paginate, err := c.taskService.GetAll(ctx.Context(), dto)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)

	}
	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(tasks, paginate),
	)
}
