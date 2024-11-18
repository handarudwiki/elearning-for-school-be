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

type subjectController struct {
	SubjectService services.SubjectService
	JwtService     services.JWTService
}

func NewSubject(app *fiber.App, subjectService services.SubjectService, jwtService services.JWTService) *subjectController {
	subjecController := subjectController{
		SubjectService: subjectService,
		JwtService:     jwtService,
	}

	subjects := app.Group("/api/v1/subjects")
	subjects.Post("/", middlewares.CheckAuth(subjecController.JwtService), subjecController.Create)
	subjects.Get("/:id", subjecController.FindByID)
	subjects.Get("/", subjecController.FindAll)
	subjects.Put("/:id", middlewares.CheckAuth(subjecController.JwtService), subjecController.Update)
	subjects.Delete("/:id", middlewares.CheckAuth(subjecController.JwtService), subjecController.Delete)
	return &subjecController
}

func (c *subjectController) Create(ctx *fiber.Ctx) (err error) {
	var createSubjectDTO dto.CreateSubjectDTO

	if err = ctx.BodyParser(&createSubjectDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(createSubjectDTO)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	subjectResponse, err := c.SubjectService.Create(&createSubjectDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(subjectResponse),
	)
}

func (c *subjectController) FindByID(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	subjectResponse, err := c.SubjectService.FindByID(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(subjectResponse),
	)
}

func (c *subjectController) FindAll(ctx *fiber.Ctx) (err error) {
	var queryDTO dto.QueryDTO

	page, size := helpers.GetPaginationParams(ctx, commons.DEFAULTPAGE, commons.DEFAULTSIZE)

	queryDTO.Page = page
	queryDTO.Size = size
	search := ctx.Query("search")

	queryDTO.Search = &search

	fmt.Println(queryDTO)

	subjectsResponse, paginate, err := c.SubjectService.FindAll(ctx.Context(), queryDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponsePagination(subjectsResponse, paginate),
	)
}

func (c *subjectController) Update(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	var updateSubjectDTO dto.UpdateSubjectDTO

	if err = ctx.BodyParser(&updateSubjectDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(updateSubjectDTO)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	subjectResponse, err := c.SubjectService.Update(ctx.Context(), id, updateSubjectDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(subjectResponse),
	)
}

func (c *subjectController) Delete(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	err = c.SubjectService.Delete(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess("Subject deleted successfully"),
	)
}
