package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type ScheduleController struct {
	scheduleService services.ScheduleService
	jwtService      services.JWTService
}

func NewScheduleController(app *fiber.App, scheduleService services.ScheduleService, jwtService services.JWTService) {
	schedule := ScheduleController{
		scheduleService: scheduleService,
		jwtService:      jwtService,
	}

	schedules := app.Group("/api/v1/schedules")
	schedules.Post("/", middlewares.CheckAuth(jwtService), schedule.Create)
	schedules.Get("/:id", schedule.GetByID)
}

func (c *ScheduleController) Create(ctx *fiber.Ctx) error {
	var dto dto.ScheduleDTO

	if err := ctx.BodyParser(&dto); err != nil {
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

	res, err := c.scheduleService.CreateSchedule(ctx.Context(), &dto)
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

func (c *ScheduleController) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.scheduleService.GetScheduleByID(ctx.Context(), id)
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