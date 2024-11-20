package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type infoController struct {
	jwtService  services.JWTService
	infoService services.InfoService
}

func NewInfo(app *fiber.App, jwtService services.JWTService, infoService services.InfoService) {
	infoContorller := infoController{
		jwtService:  jwtService,
		infoService: infoService,
	}

	infos := app.Group("/api/v1/infos")
	infos.Post("/", middlewares.CheckAuth(jwtService), infoContorller.Create)
	infos.Get("/:id", infoContorller.FindById)
}

func (c *infoController) Create(ctx *fiber.Ctx) error {
	var dto *dto.InfoDto

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

	res, err := c.infoService.Create(ctx.Context(), dto)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(res),
	)
}

func (c *infoController) FindById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	res, err := c.infoService.FindById(ctx.Context(), id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(res),
	)
}
