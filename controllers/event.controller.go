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

type eventController struct {
	eventService services.EventService
	jwtService   services.JWTService
}

func NewEvent(app *fiber.App, eventService services.EventService, jwtService services.JWTService) {
	event := eventController{
		eventService: eventService,
		jwtService:   jwtService,
	}

	events := app.Group("/api/v1/events")

	events.Post("/", middlewares.CheckAuth(jwtService), event.Create)
	events.Get("/:id", event.FindById)
	events.Put("/:id", middlewares.CheckAuth(jwtService), event.Update)
	events.Delete("/:id", middlewares.CheckAuth(jwtService), event.Delete)
	events.Get("/", event.GetAll)
}

func (c *eventController) Create(ctx *fiber.Ctx) error {
	var dto *dto.EventDto

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	errors := helpers.ValidateRequest(dto)

	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Bad request", errors),
		)
	}

	userId := ctx.Locals("userId").(int)

	dto.UserID = uint(userId)

	res, err := c.eventService.Create(ctx.Context(), dto)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *eventController) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.eventService.FindById(ctx.Context(), id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *eventController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	var dto *dto.EventDto

	if err := ctx.BodyParser(&dto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	errors := helpers.ValidateRequest(dto)

	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Bad request", errors),
		)
	}

	userId := ctx.Locals("userId").(int)

	dto.UserID = uint(userId)

	res, err := c.eventService.Update(ctx.Context(), dto, id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *eventController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.eventService.Delete(ctx.Context(), id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("Event deleted"),
	)
}

func (c *eventController) GetAll(ctx *fiber.Ctx) error {
	var dto dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	dto.Page = page
	dto.Size = size

	search := ctx.Query("search")

	dto.Search = &search

	res, paginate, err := c.eventService.FindAll(ctx.Context(), dto)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(res, paginate),
	)
}
