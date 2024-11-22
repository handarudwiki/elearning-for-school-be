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

type StandartController struct {
	standartService services.StandartService
	jwtService      services.JWTService
}

func NewStandartController(app *fiber.App, standartService services.StandartService, jwtService services.JWTService) {
	standart := StandartController{
		standartService: standartService,
		jwtService:      jwtService,
	}

	standarts := app.Group("/api/v1/standarts")
	standarts.Post("/", middlewares.CheckAuth(jwtService), standart.CreateStandart)
	standarts.Get("/:id", standart.GetStandart)
	standarts.Put("/:id", middlewares.CheckAuth(jwtService), standart.Update)
	standarts.Delete("/:id", middlewares.CheckAuth(jwtService), standart.DeleteStandart)
	standarts.Get("/", standart.FindALl)
}

func (c *StandartController) CreateStandart(ctx *fiber.Ctx) error {
	var dto dto.StandartDTO

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

	dto.TeacherID = uint(userId)

	fmt.Println(dto.TeacherID)

	res, err := c.standartService.Create(ctx.Context(), &dto)
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

func (c *StandartController) GetStandart(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.standartService.FindById(ctx.Context(), id)
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

func (c *StandartController) Update(ctx *fiber.Ctx) error {
	var dto dto.StandartDTO

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

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	userID := ctx.Locals("userId").(int)

	dto.TeacherID = uint(userID)

	res, err := c.standartService.Update(ctx.Context(), &dto, id)
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

func (c *StandartController) DeleteStandart(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.standartService.Delete(ctx.Context(), id)
	if err != nil {
		statusCode := helpers.GetHttpStatusCode(err)

		return ctx.Status(statusCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("Success delete standart"),
	)
}

func (c *StandartController) FindALl(ctx *fiber.Ctx) error {
	dto := new(dto.QueryDTO)

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	dto.Page = page
	dto.Size = size

	res, count, err := c.standartService.FindAll(ctx.Context(), dto)
	if err != nil {
		statusCode := helpers.GetHttpStatusCode(err)

		return ctx.Status(statusCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(res, count),
	)
}
