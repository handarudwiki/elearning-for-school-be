package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type AbcentController struct {
	abcentService services.AbcentService
	jwtService    services.JWTService
}

func NewAbcent(app *fiber.App, abcentService services.AbcentService, jwtService services.JWTService) {
	abcentController := AbcentController{
		abcentService: abcentService,
		jwtService:    jwtService,
	}

	abcents := app.Group("/api/v1/abcents")
	abcents.Post("/", middlewares.CheckAuth(jwtService), abcentController.Create)
	abcents.Get("/schedule/:schedule_id", middlewares.CheckAuth(jwtService), abcentController.ScheduleClassroomToday)
}

func (c *AbcentController) Create(ctx *fiber.Ctx) (err error) {
	var dto dto.CreateAbcentDTO

	if err = ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(dto)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	dto.UserID = (ctx.Locals("userId").(int))

	abcentResponse, err := c.abcentService.Create(ctx.Context(), dto)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(abcentResponse),
	)
}

func (c *AbcentController) ScheduleClassroomToday(ctx *fiber.Ctx) (err error) {
	var query dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	query.Page = page
	query.Size = size

	scheduleID, err := strconv.Atoi(ctx.Params("schedule_id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError("Invalid schedule ID"),
		)
	}

	date := ctx.Query("date")

	abcents, paginate, err := c.abcentService.FindByScheduleIDToday(ctx.Context(), scheduleID, date, query)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(abcents, paginate),
	)
}
