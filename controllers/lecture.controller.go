package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/middlewares"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/services"
)

type lectureController struct {
	lectureService services.LectureService
	jwtService     services.JWTService
}

func NewLecture(app *fiber.App, lectureService services.LectureService, jwtService services.JWTService) *lectureController {
	lectureController := lectureController{
		lectureService: lectureService,
		jwtService:     jwtService,
	}

	lectures := app.Group("/api/v1/lectures")
	lectures.Post("/", middlewares.CheckAuth(lectureController.jwtService), lectureController.Create)
	lectures.Get("/:id", lectureController.FindByID)

	return &lectureController
}

func (c *lectureController) Create(ctx *fiber.Ctx) (err error) {
	var createLectureDTO dto.CreateLectureDTO

	if err = ctx.BodyParser(&createLectureDTO); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	validationErrors := helpers.ValidateRequest(createLectureDTO)

	if len(validationErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseErrorWithData("Validation errors", validationErrors),
		)
	}

	createLectureDTO.UserID = uint(ctx.Locals("userId").(int))

	lectureResponse, err := c.lectureService.CreateLecture(ctx.Context(), createLectureDTO)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		helpers.ResponseSuccess(lectureResponse),
	)
}

func (c *lectureController) FindByID(ctx *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	lectureResponse, err := c.lectureService.FindByID(ctx.Context(), id)

	if err != nil {
		httpCode := helpers.GetHttpStatusCode(err)
		return ctx.Status(httpCode).JSON(
			helpers.ResponseError(err.Error()),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		helpers.ResponseSuccess(lectureResponse),
	)
}
