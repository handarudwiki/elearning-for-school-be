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
	infos.Put("/:id", middlewares.CheckAuth(jwtService), infoContorller.Update)
	infos.Delete("/:id", middlewares.CheckAuth(jwtService), infoContorller.Delete)
	infos.Get("/", infoContorller.FindAll)
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

func (c *infoController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	var infoDto dto.InfoDto

	if err = ctx.BodyParser(&infoDto); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}
	fmt.Printf("ID: %v\n", id)

	validationErrors := helpers.ValidateRequest(infoDto)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	infoDto.UserID = uint(ctx.Locals("userId").(int))

	classroomResponse, err := c.infoService.Update(ctx.Context(), &infoDto, id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(classroomResponse),
	)
}

func (c *infoController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.infoService.Delete(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("Success delete classroom"),
	)
}

func (c *infoController) FindAll(ctx *fiber.Ctx) error {
	var queryDTO dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	queryDTO.Page = page
	queryDTO.Size = size

	search := ctx.Query("search")

	queryDTO.Search = &search

	infos, totalPages, err := c.infoService.FindAll(ctx.Context(), queryDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(infos, totalPages),
	)
}
